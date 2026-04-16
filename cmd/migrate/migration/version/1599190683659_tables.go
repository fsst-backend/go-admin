package version

import (
	"runtime"
	"strconv"

	"github.com/go-admin-team/go-admin-core/sdk/config"

	adminmodels "go-admin/app/admin/models"
	jobsmodels "go-admin/app/jobs/models"
	toolsmodels "go-admin/app/other/models/tools"
	"go-admin/cmd/migrate/migration"
	commonmodels "go-admin/common/models"

	"gorm.io/gorm"
)

func init() {
	_, fileName, _, _ := runtime.Caller(0)
	migration.Migrate.SetVersion(migration.GetFilename(fileName), _1599190683659Tables)
}

func _1599190683659Tables(db *gorm.DB, version string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if config.DatabaseConfig.Driver == "mysql" || config.DatabaseConfig.Driver == "tidb" {
			tx = tx.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4")
		}
		err := tx.Migrator().AutoMigrate(
			new(adminmodels.SysDept),
			new(adminmodels.SysRoleDept),
			new(adminmodels.SysConfig),
			new(toolsmodels.SysTables),
			new(toolsmodels.SysColumns),
			new(adminmodels.SysMenu),
			new(adminmodels.SysRoleMenu),
			new(adminmodels.SysLoginLog),
			new(adminmodels.SysOperaLog),
			new(adminmodels.SysUserRole),
			new(adminmodels.SysRolePermission),
			new(adminmodels.SysUser),
			new(adminmodels.SysRole),
			new(adminmodels.SysPost),
			new(adminmodels.SysDictData),
			new(adminmodels.SysDictType),
			new(jobsmodels.SysJob),
			new(adminmodels.SysApi),
			new(adminmodels.SysPermission),
			new(adminmodels.SysPermissionApi),
			new(adminmodels.CasbinRule),
			new(commonmodels.TbDemo),
		)
		if err != nil {
			return err
		}
		if err := migration.InitDb(tx); err != nil {
			return err
		}
		if err := fixSysMenuPaths(tx); err != nil {
			return err
		}
		// Session 重置 Statement，防止前面操作（InitDb / fixSysMenuPaths）
		// 在 tx 上残留的 schema 缓存导致 Create 时 reflect panic。
		return tx.Session(&gorm.Session{}).Create(&commonmodels.Migration{
			Version: version,
		}).Error
	})
}

// fixSysMenuPaths 原 1653638869132_migrate：补全 sys_menu.menu_path
// 每次查询/更新都用 tx.Session 隔离 Statement，避免污染调用方的 tx。
func fixSysMenuPaths(tx *gorm.DB) error {
	var list []adminmodels.SysMenu
	if err := tx.Session(&gorm.Session{}).Model(&adminmodels.SysMenu{}).Order("parent_id,menu_id").Find(&list).Error; err != nil {
		return err
	}
	for _, v := range list {
		var path string
		if v.ParentId == 0 {
			path = "/0/" + strconv.Itoa(v.MenuId)
		} else {
			var parent adminmodels.SysMenu
			err := tx.Session(&gorm.Session{}).Model(&adminmodels.SysMenu{}).Where("menu_id = ?", v.ParentId).First(&parent).Error
			if err != nil {
				if err == gorm.ErrRecordNotFound {
					continue
				}
				return err
			}
			path = parent.MenuPath + "/" + strconv.Itoa(v.MenuId)
		}
		if err := tx.Session(&gorm.Session{}).Model(&adminmodels.SysMenu{}).Where("menu_id = ?", v.MenuId).Update("menu_path", path).Error; err != nil {
			return err
		}
	}
	return nil
}
