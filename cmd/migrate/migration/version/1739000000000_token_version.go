package version

import (
	"runtime"

	"go-admin/cmd/migrate/migration"
	"go-admin/cmd/migrate/migration/models"
	common "go-admin/common/models"

	"gorm.io/gorm"
)

func init() {
	_, fileName, _, _ := runtime.Caller(0)
	migration.Migrate.SetVersion(migration.GetFilename(fileName), _1739000000000TokenVersion)
}

// _1739000000000TokenVersion 添加 token_version 字段，用于单设备登录（新登录踢出旧会话）
func _1739000000000TokenVersion(db *gorm.DB, version string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Migrator().AutoMigrate(&models.SysUser{}); err != nil {
			return err
		}
		if err := tx.Migrator().AutoMigrate(&models.SysLoginLog{}); err != nil {
			return err
		}
		return tx.Create(&common.Migration{Version: version}).Error
	})
}
