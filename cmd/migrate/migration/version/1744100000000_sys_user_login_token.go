package version

import (
	"runtime"

	adminmodels "go-admin/app/admin/models"
	"go-admin/cmd/migrate/migration"
	commonmodels "go-admin/common/models"

	"gorm.io/gorm"
)

func init() {
	_, fileName, _, _ := runtime.Caller(0)
	migration.Migrate.SetVersion(migration.GetFilename(fileName), _1744100000000SysUserLoginToken)
}

// _1744100000000SysUserLoginToken 单设备登录版本表；若 sys_user 上仍有旧列则回填后删除（合并原 173900/174400 职责）
func _1744100000000SysUserLoginToken(db *gorm.DB, version string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Migrator().AutoMigrate(&commonmodels.SysUserLoginToken{}); err != nil {
			return err
		}
		if tx.Migrator().HasColumn(&adminmodels.SysUser{}, "token_version") {
			if err := tx.Exec(`
INSERT INTO sys_user_login_token (user_id, login_channel, token_version)
SELECT user_id, ?, COALESCE(token_version, 0) FROM sys_user
`, "admin").Error; err != nil {
				return err
			}
		}
		if tx.Migrator().HasColumn(&adminmodels.SysUser{}, "token_version_cs") {
			if err := tx.Exec(`
INSERT INTO sys_user_login_token (user_id, login_channel, token_version)
SELECT user_id, ?, COALESCE(token_version_cs, 0) FROM sys_user
`, "cs_app").Error; err != nil {
				return err
			}
		}
		if tx.Migrator().HasColumn(&adminmodels.SysUser{}, "token_version_cs") {
			if err := tx.Migrator().DropColumn(&adminmodels.SysUser{}, "token_version_cs"); err != nil {
				return err
			}
		}
		if tx.Migrator().HasColumn(&adminmodels.SysUser{}, "token_version") {
			if err := tx.Migrator().DropColumn(&adminmodels.SysUser{}, "token_version"); err != nil {
				return err
			}
		}
		// Session 重置 Statement，防止前面 Migrator 操作残留的 schema 缓存导致 Create 时 reflect panic。
		return tx.Session(&gorm.Session{}).Create(&commonmodels.Migration{Version: version}).Error
	})
}
