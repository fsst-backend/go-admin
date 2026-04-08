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

func ExecSql(db *gorm.DB, filePath string) error {
	sql, err := readSQLFile(filePath)
	if err != nil {
		fmt.Println("数据库基础数据初始化脚本读取失败！原因:", err.Error(), "path:", filePath)
		return err
	}
	sqlList := strings.Split(sql, ";")
	for i := 0; i < len(sqlList)-1; i++ {
		if strings.Contains(sqlList[i], "--") {
			fmt.Println(sqlList[i])
			continue
		}
		stmt := strings.Replace(sqlList[i]+";", "\n", "", -1)
		stmt = strings.TrimSpace(stmt)
		if err = db.Exec(stmt).Error; err != nil {
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
