package models

import (
	"fmt"
	"strconv"

	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"gorm.io/gorm"

	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk/config"
)

type DataPermission struct {
	DataScope string
	UserId    int
	DeptId    int
	RoleId    int
}

func (e *DataPermission) GetDataScope(tableName string, db *gorm.DB) (*gorm.DB, error) {

	if !config.ApplicationConfig.EnableDP {
		usageStr := `数据权限已经为您` + pkg.Green(`关闭`) + `，如需开启请参考配置文件字段说明`
		log.Debug("%s\n", usageStr)
		return db, nil
	}
	// 查询用户
	user := new(SysUser)
	if err := db.First(user, e.UserId).Error; err != nil {
		return nil, fmt.Errorf("获取用户数据出错: %w", err)
	}

	// 查询用户所有角色
	var userRoles []SysUserRole
	if err := db.Where("user_id = ?", e.UserId).Find(&userRoles).Error; err != nil {
		return nil, fmt.Errorf("获取用户角色出错: %w", err)
	}

	if len(userRoles) == 0 {
		// 没有角色 → 默认只能看自己
		return db.Where(tableName+".create_by = ?", user.UserId), nil
	}

	// 收集所有角色
	roleIds := make([]int64, 0, len(userRoles))
	for _, ur := range userRoles {
		roleIds = append(roleIds, int64(ur.RoleId))
	}

	var roles []SysRole
	if err := db.Where("role_id IN ?", roleIds).Find(&roles).Error; err != nil {
		return nil, fmt.Errorf("获取角色信息出错: %w", err)
	}

	// 构建 OR 条件
	db = db.Where(func(tx *gorm.DB) *gorm.DB {
		for i, role := range roles {
			switch role.DataScope {
			case "2": // 自定义部门
				var roleDepts []SysRoleDept
				if err := db.Where("role_id = ?", role.RoleId).Find(&roleDepts).Error; err != nil {
					continue
				}
				deptIds := make([]int64, 0, len(roleDepts))
				for _, rd := range roleDepts {
					deptIds = append(deptIds, int64(rd.DeptId))
				}
				if i == 0 {
					tx = tx.Where(tableName+".create_by IN (SELECT user_id FROM sys_user WHERE dept_id IN ?)", deptIds)
				} else {
					tx = tx.Or(tableName+".create_by IN (SELECT user_id FROM sys_user WHERE dept_id IN ?)", deptIds)
				}

			case "3": // 本部门
				if i == 0 {
					tx = tx.Where(tableName+".create_by IN (SELECT user_id FROM sys_user WHERE dept_id = ?)", user.DeptId)
				} else {
					tx = tx.Or(tableName+".create_by IN (SELECT user_id FROM sys_user WHERE dept_id = ?)", user.DeptId)
				}

			case "4": // 本部门及子部门
				pathLike := "%/" + strconv.FormatInt(int64(user.DeptId), 10) + "/%"
				if i == 0 {
					tx = tx.Where(tableName+".create_by IN (SELECT user_id FROM sys_user WHERE dept_id IN (SELECT dept_id FROM sys_dept WHERE dept_path LIKE ?))", pathLike)
				} else {
					tx = tx.Or(tableName+".create_by IN (SELECT user_id FROM sys_user WHERE dept_id IN (SELECT dept_id FROM sys_dept WHERE dept_path LIKE ?))", pathLike)
				}

			case "5", "": // 仅本人
				if i == 0 {
					tx = tx.Where(tableName+".create_by = ?", user.UserId)
				} else {
					tx = tx.Or(tableName+".create_by = ?", user.UserId)
				}

			default:
				// 默认全权限，可不加限制
			}
		}
		return tx
	})

	return db, nil
}

//func DataScopes(tableName string, userId int) func(db *gorm.DB) *gorm.DB {
//	return func(db *gorm.DB) *gorm.DB {
//		user := new(SysUser)
//		role := new(SysRole)
//		user.UserId = userId
//		err := db.Find(user, userId).Error
//		if err != nil {
//			db.Error = errors.New("获取用户数据出错 msg:" + err.Error())
//			return db
//		}
//		err = db.Find(role, user.RoleId).Error
//		if err != nil {
//			db.Error = errors.New("获取用户数据出错 msg:" + err.Error())
//			return db
//		}
//		if role.DataScope == "2" {
//			return db.Where(tableName+".create_by in (select sys_user.user_id from sys_role_dept left join sys_user on sys_user.dept_id=sys_role_dept.dept_id where sys_role_dept.role_id = ?)", user.RoleId)
//		}
//		if role.DataScope == "3" {
//			return db.Where(tableName+".create_by in (SELECT user_id from sys_user where dept_id = ? )", user.DeptId)
//		}
//		if role.DataScope == "4" {
//			return db.Where(tableName+".create_by in (SELECT user_id from sys_user where sys_user.dept_id in(select dept_id from sys_dept where dept_path like ? ))", "%"+pkg.IntToString(user.DeptId)+"%")
//		}
//		if role.DataScope == "5" || role.DataScope == "" {
//			return db.Where(tableName+".create_by = ?", userId)
//		}
//		return db
//	}
//}
