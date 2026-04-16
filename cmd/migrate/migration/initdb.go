package migration

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"go-admin/common/global"

	"gorm.io/gorm"
)

// isDuplicateKeyError 种子脚本重复执行时主键/唯一键冲突，可跳过（避免 sys_migration 缺失时整段 InitDb 失败）
func isDuplicateKeyError(err error) bool {
	if err == nil {
		return false
	}
	msg := err.Error()
	return strings.Contains(msg, "1062") || // MySQL / TiDB: Duplicate entry
		strings.Contains(msg, "Duplicate entry") ||
		strings.Contains(msg, "duplicate key value violates unique constraint") || // PostgreSQL
		strings.Contains(msg, "UNIQUE constraint failed") // SQLite
}

// sqlBaseDir 返回存放 db.sql 等脚本的目录：环境变量 > /app（容器常见）> 仓库内 config/
func sqlBaseDir() string {
	if d := strings.TrimSpace(os.Getenv("GO_ADMIN_SQL_DIR")); d != "" {
		return d
	}
	if st, err := os.Stat("/app/db.sql"); err == nil && !st.IsDir() {
		return "/app"
	}
	return "config"
}

// InitDb 执行基础数据 SQL（与仓库中 config/*.sql 对应；生产可把 SQL 放在 GO_ADMIN_SQL_DIR 或挂载到 /app）
func InitDb(db *gorm.DB) (err error) {
	base := sqlBaseDir()
	join := func(name string) string { return filepath.Join(base, name) }

	switch global.Driver {
	case "postgres":
		if err = ExecSql(db, join("db.sql")); err != nil {
			return err
		}
		return ExecSql(db, join("pg.sql"))
	case "mysql", "tidb":
		if err = ExecSql(db, join("db-begin-mysql.sql")); err != nil {
			return err
		}
		if err = ExecSql(db, join("db.sql")); err != nil {
			return err
		}
		return ExecSql(db, join("db-end-mysql.sql"))
	case "sqlserver":
		return ExecSql(db, join("db-sqlserver.sql"))
	default:
		// sqlite3 等：仅执行主脚本
		return ExecSql(db, join("db.sql"))
	}
}

// stripSQLLineComments 去掉仅由「行首 --」构成的注释行，保留 INSERT 等语句。
// 历史实现用 strings.Contains(chunk, "--") 跳过整段：db.sql 里「-- 表说明」与 INSERT 在同一段（按 ; 切分）时会把 INSERT INTO sys_user 等一并跳过。
func stripSQLLineComments(chunk string) string {
	lines := strings.Split(chunk, "\n")
	out := make([]string, 0, len(lines))
	for _, line := range lines {
		t := strings.TrimSpace(line)
		if t == "" {
			continue
		}
		if strings.HasPrefix(t, "--") {
			continue
		}
		out = append(out, line)
	}
	return strings.TrimSpace(strings.Join(out, "\n"))
}

func ExecSql(db *gorm.DB, filePath string) error {
	sql, err := readSQLFile(filePath)
	if err != nil {
		fmt.Println("数据库基础数据初始化脚本读取失败！原因:", err.Error(), "path:", filePath)
		return err
	}
	sqlList := strings.Split(sql, ";")
	for i := 0; i < len(sqlList)-1; i++ {
		stmt := stripSQLLineComments(sqlList[i])
		if stmt == "" {
			continue
		}
		stmt = strings.Replace(stmt+";", "\n", "", -1)
		stmt = strings.TrimSpace(stmt)
		// 使用新 Session 执行每条语句，避免前一条的错误污染 GORM 事务状态。
		// 否则 duplicate key 等可跳过的错误会导致后续所有 tx.Exec 失败。
		execDB := db.Session(&gorm.Session{NewDB: true})
		if err = execDB.Exec(stmt).Error; err != nil {
			if isDuplicateKeyError(err) {
				log.Printf("initdb 跳过重复数据（已存在）: %v", err)
				continue
			}
			log.Printf("error sql: %s", stmt)
			if !strings.Contains(err.Error(), "Query was empty") {
				return err
			}
		}
	}
	return nil
}

func readSQLFile(filePath string) (string, error) {
	contents, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	// 与历史实现一致：去掉首行换行，避免首条语句异常
	result := strings.Replace(string(contents), "\n", "", 1)
	return result, nil
}
