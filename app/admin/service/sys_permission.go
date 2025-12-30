package service

import (
	"errors"

	"github.com/casbin/casbin/v2"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	cDto "go-admin/common/dto"
	cModels "go-admin/common/models"
	"go-admin/common/mycasbin"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"
)

type SysPermission struct {
	service.Service
}

// GetPage 获取 SysPermission 列表
func (e *SysPermission) GetPage(c *dto.SysPermissionGetPageReq, list *[]models.SysPermission, count *int64) error {
	err := e.Orm.
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.PaginateOffsetLimit(c.GetLimit(), c.GetOffset()),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("Service GetSysPermissionPage error:%s", err)
		return err
	}
	return nil
}

// Get 获取单个 SysPermission 对象
func (e *SysPermission) Get(d *dto.SysPermissionGetReq, model *models.SysPermission) error {
	err := e.Orm.
		FirstOrInit(model, d.GetId()).
		Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		_ = e.AddError(err)
		return err
	}
	if model.Id == 0 {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetSysPermission error: %s", err)
		_ = e.AddError(err)
		return err
	}

	// 追加加载权限关联的接口列表
	var permApis []models.SysPermissionApi
	if err := e.Orm.Where("permission_id = ?", model.Id).Find(&permApis).Error; err != nil {
		return err
	}
	if len(permApis) == 0 {
		return nil
	}

	apiIdSet := make(map[int]struct{}, len(permApis))
	for _, pa := range permApis {
		apiIdSet[pa.ApiId] = struct{}{}
	}
	apiIds := make([]int, 0, len(apiIdSet))
	for id := range apiIdSet {
		apiIds = append(apiIds, id)
	}

	var apis []models.SysApi
	if err := e.Orm.Where("id in ?", apiIds).Find(&apis).Error; err != nil {
		return err
	}
	model.Apis = apis

	return nil
}

// Insert 创建 SysPermission 对象
func (e *SysPermission) Insert(c *dto.SysPermissionInsertReq) error {
	var data models.SysPermission
	c.Generate(&data)

	// 使用事务创建权限和 API 关联
	return e.Orm.Transaction(func(tx *gorm.DB) error {
		// 1. 创建权限记录
		if err := tx.Create(&data).Error; err != nil {
			e.Log.Errorf("Service InsertSysPermission error:%s", err)
			return err
		}

		// 2. 关联 API，如果有提供 API ID 列表
		if len(c.ApiIds) > 0 {
			permissionApis := make([]models.SysPermissionApi, 0, len(c.ApiIds))
			for _, apiId := range c.ApiIds {
				permissionApis = append(permissionApis, models.SysPermissionApi{
					PermissionId: data.Id,
					ApiId:        apiId,
					ControlBy: cModels.ControlBy{
						CreateBy: c.CreateBy,
						UpdateBy: c.UpdateBy,
					},
				})
			}
			if err := tx.Create(&permissionApis).Error; err != nil {
				e.Log.Errorf("Service InsertSysPermissionApi error:%s", err)
				return err
			}
		}

		return nil
	})
}

// Update 修改 SysPermission 对象
func (e *SysPermission) Update(c *dto.SysPermissionUpdateReq, cb *casbin.SyncedEnforcer) error {
	var model models.SysPermission
	if err := e.Orm.First(&model, c.GetId()).Error; err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	c.Generate(&model)

	// 使用事务更新权限和 API 关联
	return e.Orm.Transaction(func(tx *gorm.DB) error {
		// 1. 更新权限基本信息
		updateData := map[string]interface{}{
			"code":      model.Code,
			"name":      model.Name,
			"type":      model.Type,
			"parent_id": model.ParentId,
			"sort":      model.Sort,
			"status":    model.Status,
			"remark":    model.Remark,
		}
		db := tx.Model(&model).Where("id = ?", c.GetId()).Updates(updateData)
		if err := db.Error; err != nil {
			e.Log.Errorf("Service UpdateSysPermission error:%s", err)
			return err
		}
		if db.RowsAffected == 0 {
			return errors.New("无权更新该数据")
		}

		// 2. 删除旧的 API 关联
		if err := tx.Where("permission_id = ?", c.Id).Delete(&models.SysPermissionApi{}).Error; err != nil {
			e.Log.Errorf("Service DeleteSysPermissionApi error:%s", err)
			return err
		}

		// 3. 创建新的 API 关联
		if len(c.ApiIds) > 0 {
			permissionApis := make([]models.SysPermissionApi, 0, len(c.ApiIds))
			for _, apiId := range c.ApiIds {
				permissionApis = append(permissionApis, models.SysPermissionApi{
					PermissionId: c.Id,
					ApiId:        apiId,
					ControlBy: cModels.ControlBy{
						CreateBy: c.CreateBy,
						UpdateBy: c.UpdateBy,
					},
				})
			}
			if err := tx.Create(&permissionApis).Error; err != nil {
				e.Log.Errorf("Service InsertSysPermissionApi error:%s", err)
				return err
			}
		}

		// 4. 查询使用了该权限的所有角色
		var rolePerms []models.SysRolePermission
		if err := tx.Where("permission_id = ?", c.Id).Find(&rolePerms).Error; err != nil {
			e.Log.Errorf("Query role permissions error:%s", err)
			return err
		}

		// 5. 如果有角色使用了该权限，需要重建这些角色的 Casbin 策略
		if len(rolePerms) > 0 && cb != nil {
			// 获取所有相关角色 ID
			roleIdSet := make(map[int]struct{}, len(rolePerms))
			for _, rp := range rolePerms {
				roleIdSet[rp.RoleId] = struct{}{}
			}
			roleIds := make([]int, 0, len(roleIdSet))
			for id := range roleIdSet {
				roleIds = append(roleIds, id)
			}

			// 查询角色信息
			var roles []models.SysRole
			if err := tx.Where("role_id in ?", roleIds).Find(&roles).Error; err != nil {
				e.Log.Errorf("Query roles error:%s", err)
				return err
			}

			// 为每个角色重建 Casbin 策略
			for _, role := range roles {
				if role.RoleKey == mycasbin.SuperAdmin {
					continue // 跳过 SuperAdmin 角色
				}
				if err := e.rebuildRoleCasbinPolicy(tx, cb, role.RoleId, role.RoleKey); err != nil {
					e.Log.Errorf("Rebuild role casbin policy error:%s", err)
					return err
				}
			}
		}

		return nil
	})
}

// Remove 删除 SysPermission
func (e *SysPermission) Remove(d *dto.SysPermissionDeleteReq, cb *casbin.SyncedEnforcer) error {
	// 使用事务处理删除操作
	return e.Orm.Transaction(func(tx *gorm.DB) error {
		// 1. 先查询要删除的权限 ID 列表
		permIds := d.GetId().([]int)
		if len(permIds) == 0 {
			return errors.New("没有提供要删除的权限 ID")
		}

		// 2. 查询使用了这些权限的所有角色
		var rolePerms []models.SysRolePermission
		if err := tx.Where("permission_id in ?", permIds).Find(&rolePerms).Error; err != nil {
			e.Log.Errorf("Query role permissions error:%s", err)
			return err
		}

		// 3. 获取受影响的角色 ID
		roleIdSet := make(map[int]struct{})
		for _, rp := range rolePerms {
			roleIdSet[rp.RoleId] = struct{}{}
		}
		roleIds := make([]int, 0, len(roleIdSet))
		for id := range roleIdSet {
			roleIds = append(roleIds, id)
		}

		// 4. 删除权限与 API 的关联
		if err := tx.Where("permission_id in ?", permIds).Delete(&models.SysPermissionApi{}).Error; err != nil {
			e.Log.Errorf("Delete permission api relations error:%s", err)
			return err
		}

		// 5. 删除角色与权限的关联
		if err := tx.Where("permission_id in ?", permIds).Delete(&models.SysRolePermission{}).Error; err != nil {
			e.Log.Errorf("Delete role permission relations error:%s", err)
			return err
		}

		// 6. 删除权限记录
		var data models.SysPermission
		db := tx.Delete(&data, permIds)
		if err := db.Error; err != nil {
			e.Log.Errorf("Service RemoveSysPermission error:%s", err)
			return err
		}
		if db.RowsAffected == 0 {
			return errors.New("无权删除该数据")
		}

		// 7. 如果有角色使用了这些权限，需要重建这些角色的 Casbin 策略
		if len(roleIds) > 0 && cb != nil {
			// 查询角色信息
			var roles []models.SysRole
			if err := tx.Where("role_id in ?", roleIds).Find(&roles).Error; err != nil {
				e.Log.Errorf("Query roles error:%s", err)
				return err
			}

			// 为每个角色重建 Casbin 策略
			for _, role := range roles {
				if err := e.rebuildRoleCasbinPolicy(tx, cb, role.RoleId, role.RoleKey); err != nil {
					e.Log.Errorf("Rebuild role casbin policy error:%s", err)
					return err
				}
			}
		}

		return nil
	})
}

// rebuildRoleCasbinPolicy 重建指定角色的 Casbin 策略
func (e *SysPermission) rebuildRoleCasbinPolicy(tx *gorm.DB, cb *casbin.SyncedEnforcer, roleId int, roleKey string) error {
	// 1. 查询该角色的所有权限
	var rolePerms []models.SysRolePermission
	if err := tx.Where("role_id = ?", roleId).Find(&rolePerms).Error; err != nil {
		return err
	}

	// 2. 获取权限 ID 列表
	permIds := make([]int, 0, len(rolePerms))
	for _, rp := range rolePerms {
		permIds = append(permIds, rp.PermissionId)
	}

	if len(permIds) == 0 {
		// 没有权限,删除所有旧策略后返回
		_, err := cb.RemoveFilteredPolicy(0, roleKey)
		return err
	}

	// 3. 从权限 API 关系表获取 API 权限
	var permApis []models.SysPermissionApi
	if err := tx.Where("permission_id in ?", permIds).Find(&permApis).Error; err != nil {
		return err
	}

	// 4. 获取所有相关的 API ID
	apiIds := make([]int, 0, len(permApis))
	for _, permApi := range permApis {
		apiIds = append(apiIds, permApi.ApiId)
	}

	if len(apiIds) == 0 {
		// 没有 API 关联,删除所有旧策略后返回
		_, err := cb.RemoveFilteredPolicy(0, roleKey)
		return err
	}

	// 5. 一次性获取所有 API 信息
	var apis []models.SysApi
	if err := tx.Where("id in ?", apiIds).Find(&apis).Error; err != nil {
		return err
	}

	// 6. 将 API 信息放入 map 中便于快速查找
	apiMap := make(map[int]models.SysApi, len(apis))
	for _, api := range apis {
		apiMap[api.Id] = api
	}

	// 7. 构建新的 Casbin 策略
	policies := [][]string{}
	for _, permApi := range permApis {
		api, exists := apiMap[permApi.ApiId]
		if !exists {
			continue // 如果 API 不存在,跳过
		}
		policy := []string{roleKey, api.Path, api.Action}
		policies = append(policies, policy)
	}

	// 8. 删除旧的 Casbin 策略
	_, err := cb.RemoveFilteredPolicy(0, roleKey)
	if err != nil {
		return err
	}

	// 9. 写入新的 Casbin 策略
	if len(policies) > 0 {
		_, err = cb.AddNamedPolicies("p", policies)
		if err != nil {
			return err
		}
	}

	return nil
}
