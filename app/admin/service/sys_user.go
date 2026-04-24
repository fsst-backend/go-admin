package service

import (
	"errors"
	"fmt"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"

	"github.com/casbin/casbin/v2"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type SysUser struct {
	service.Service
}

// GetPage 获取SysUser列表
func (e *SysUser) GetPage(c *dto.SysUserGetPageReq, p *actions.DataPermission, list *[]models.SysUser, count *int64) error {
	var err error
	var data models.SysUser

	db := e.Orm.Debug().
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.PaginateOffsetLimit(c.GetLimit(), c.GetOffset()), // 使用 offset/limit 分页
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count)
	if err = db.Error; err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}

	// 手动加载部门信息，避免使用外键/关联查询
	if len(*list) == 0 {
		return nil
	}

	// 收集部门ID
	deptIdSet := make(map[int]struct{})
	for i := range *list {
		u := (*list)[i]
		if u.DeptId != 0 {
			deptIdSet[u.DeptId] = struct{}{}
		}
	}

	// 收集用户ID用于查询角色
	userIds := make([]int, 0, len(*list))
	for i := range *list {
		u := (*list)[i]
		userIds = append(userIds, u.UserId)
	}

	// 收集岗位ID
	postIdSet := make(map[int]struct{})
	for i := range *list {
		u := (*list)[i]
		if u.PostId != 0 {
			postIdSet[u.PostId] = struct{}{}
		}
	}

	// 加载部门信息
	deptMap := make(map[int]*models.SysDept)
	if len(deptIdSet) > 0 {
		deptIds := make([]int, 0, len(deptIdSet))
		for id := range deptIdSet {
			deptIds = append(deptIds, id)
		}

		var depts []models.SysDept
		if err = e.Orm.Where("dept_id in ?", deptIds).Find(&depts).Error; err != nil {
			e.Log.Errorf("db error: %s", err)
			return err
		}

		for i := range depts {
			dept := &depts[i]
			deptMap[dept.DeptId] = dept
		}
	}

	// 加载角色信息
	roleMap := make(map[int]*models.SysRole)
	userRoleMap := make(map[int][]*models.SysRole)
	if len(userIds) > 0 {
		// 查询用户角色关联关系
		var userRoles []models.SysUserRole
		if err = e.Orm.Where("user_id in ?", userIds).Find(&userRoles).Error; err != nil {
			e.Log.Errorf("db error: %s", err)
			return err
		}

		// 获取所有相关的角色ID
		roleIds := make([]int, 0)
		for _, ur := range userRoles {
			roleIds = append(roleIds, ur.RoleId)
		}

		// 查询角色信息
		var roles []models.SysRole
		if len(roleIds) > 0 {
			if err = e.Orm.Where("role_id in ?", roleIds).Find(&roles).Error; err != nil {
				e.Log.Errorf("db error: %s", err)
				return err
			}

			for i := range roles {
				role := &roles[i]
				roleMap[role.RoleId] = role
			}
		}

		// 建立用户与角色的映射关系
		for _, ur := range userRoles {
			if role, ok := roleMap[ur.RoleId]; ok {
				if _, exists := userRoleMap[ur.UserId]; !exists {
					userRoleMap[ur.UserId] = make([]*models.SysRole, 0)
				}
				userRoleMap[ur.UserId] = append(userRoleMap[ur.UserId], role)
			}
		}
	}

	// 加载岗位信息
	postMap := make(map[int]*models.SysPost)
	if len(postIdSet) > 0 {
		postIds := make([]int, 0, len(postIdSet))
		for id := range postIdSet {
			postIds = append(postIds, id)
		}

		var posts []models.SysPost
		if err = e.Orm.Where("post_id in ?", postIds).Find(&posts).Error; err != nil {
			e.Log.Errorf("db error: %s", err)
			return err
		}

		for i := range posts {
			post := &posts[i]
			postMap[post.PostId] = post
		}
	}

	// 填充用户信息
	for i := range *list {
		u := &(*list)[i]
		if dept, ok := deptMap[u.DeptId]; ok {
			u.Dept = dept
		}
		if roles, ok := userRoleMap[u.UserId]; ok {
			u.RoleIds = make([]int, len(roles))
			for j, role := range roles {
				u.RoleIds[j] = role.RoleId
			}
		}
		if post, ok := postMap[u.PostId]; ok {
			u.Post = post
		}
	}

	return nil
}

// Get 获取SysUser对象
func (e *SysUser) Get(d *dto.SysUserById, p *actions.DataPermission, model *models.SysUser) error {
	var data models.SysUser

	err := e.Orm.Model(&data).Debug().
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}

	// 手动加载部门信息
	if model.DeptId != 0 {
		var dept models.SysDept
		err = e.Orm.First(&dept, model.DeptId).Error
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}
		} else {
			model.Dept = &dept
		}
	}

	// 手动加载角色信息
	var userRoles []models.SysUserRole
	err = e.Orm.Where("user_id = ?", model.UserId).Find(&userRoles).Error
	if err != nil {
		return err
	}

	// 获取角色ID列表
	roleIds := make([]int, 0, len(userRoles))
	for _, ur := range userRoles {
		roleIds = append(roleIds, ur.RoleId)
	}

	// 查询角色信息
	if len(roleIds) > 0 {
		var roles []models.SysRole
		err = e.Orm.Where("role_id in ?", roleIds).Find(&roles).Error
		if err != nil {
			return err
		}

		model.RoleIds = make([]int, len(roles))
		for i, role := range roles {
			model.RoleIds[i] = role.RoleId
		}
	}

	// 手动加载岗位信息
	if model.PostId != 0 {
		var post models.SysPost
		err = e.Orm.First(&post, model.PostId).Error
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}
		} else {
			model.Post = &post
		}
	}

	return nil
}

// Insert 创建SysUser对象
func (e *SysUser) Insert(c *dto.SysUserInsertReq, cb *casbin.SyncedEnforcer) error {
	var err error
	var data models.SysUser
	var i int64
	err = e.Orm.Model(&data).Where("username = ?", c.Username).Count(&i).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if i > 0 {
		err := errors.New("用户名已存在！")
		e.Log.Errorf("db error: %s", err)
		return err
	}
	c.Generate(&data)

	// 使用事务确保数据一致性
	tx := e.Orm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// 创建用户
	err = tx.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}

	// 如果有角色ID，创建用户角色关联
	if len(c.RoleIds) > 0 {
		var roles []models.SysRole
		// 验证角色是否存在
		err = tx.Where("role_id in ?", c.RoleIds).Find(&roles).Error
		if err != nil {
			e.Log.Errorf("Query roles error: %s", err)
			return err
		}

		if len(roles) != len(c.RoleIds) {
			err = errors.New("部分角色不存在")
			e.Log.Errorf("Some roles not found")
			return err
		}

		// 创建用户角色关系
		userRoles := make([]models.SysUserRole, 0, len(c.RoleIds))
		for _, roleId := range c.RoleIds {
			ur := models.SysUserRole{
				UserId: data.UserId,
				RoleId: roleId,
			}
			ur.SetCreateBy(c.CreateBy)
			ur.SetUpdateBy(c.UpdateBy)
			userRoles = append(userRoles, ur)
		}

		err = tx.Create(&userRoles).Error
		if err != nil {
			e.Log.Errorf("Create user-role relations error: %s", err)
			return err
		}

		// 数据库操作成功后，同步到Casbin
		if cb != nil {
			userSubject := fmt.Sprintf("user_%d", data.UserId)

			// 添加Casbin用户角色关联
			if len(roles) > 0 {
				for _, role := range roles {
					_, err = cb.AddGroupingPolicy(userSubject, role.RoleKey)
					if err != nil {
						e.Log.Errorf("Add casbin user-role relation error: %s", err)
						return err
					}
				}
			}
		}
	}

	return nil
}

// Update 修改SysUser对象
func (e *SysUser) Update(c *dto.SysUserUpdateReq, p *actions.DataPermission) error {
	var err error
	var model models.SysUser
	db := e.Orm.Scopes(
		actions.Permission(model.TableName(), p),
	).First(&model, c.GetId())
	if err = db.Error; err != nil {
		e.Log.Errorf("Service UpdateSysUser error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")

	}
	c.Generate(&model)
	update := e.Orm.Model(&model).Where("user_id = ?", &model.UserId).Omit("password", "salt").Updates(&model)
	if err = update.Error; err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if update.RowsAffected == 0 {
		err = errors.New("update userinfo error")
		log.Warnf("db update error")
		return err
	}
	return nil
}

// UpdateAvatar 更新用户头像
func (e *SysUser) UpdateAvatar(userId int, c *dto.UpdateSysUserAvatarReq, p *actions.DataPermission) error {
	var err error
	var model models.SysUser
	db := e.Orm.Scopes(
		actions.Permission(model.TableName(), p),
	).Where("user_id = ?", userId).First(&model)
	if err = db.Error; err != nil {
		e.Log.Errorf("Service UpdateSysUser error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	// 只更新头像字段，使用 map 指定特定字段
	updateData := map[string]interface{}{
		"avatar":    c.Avatar,
		"update_by": c.UpdateBy,
	}
	err = e.Orm.Table(model.TableName()).Where("user_id = ?", userId).Updates(updateData).Error
	if err != nil {
		e.Log.Errorf("Service UpdateSysUser error: %s", err)
		return err
	}
	return nil
}

// UpdateStatus 更新用户状态
func (e *SysUser) UpdateStatus(c *dto.UpdateSysUserStatusReq, p *actions.DataPermission) error {
	var err error
	var model models.SysUser
	db := e.Orm.Scopes(
		actions.Permission(model.TableName(), p),
	).First(&model, c.GetId())
	if err = db.Error; err != nil {
		e.Log.Errorf("Service UpdateSysUser error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	updateData := map[string]interface{}{
		"status": c.Status,
	}
	if c.UpdateBy > 0 {
		updateData["update_by"] = c.UpdateBy
	}
	err = e.Orm.Table(model.TableName()).Where("user_id =? ", c.UserId).Updates(updateData).Error
	if err != nil {
		e.Log.Errorf("Service UpdateSysUser error: %s", err)
		return err
	}
	return nil
}

// ResetPwd 重置用户密码
func (e *SysUser) ResetPwd(c *dto.ResetSysUserPwdReq, p *actions.DataPermission) error {
	var err error
	var model models.SysUser
	db := e.Orm.Scopes(
		actions.Permission(model.TableName(), p),
	).First(&model, c.GetId())
	if err = db.Error; err != nil {
		e.Log.Errorf("At Service ResetSysUserPwd error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	c.Generate(&model)
	err = e.Orm.Omit("username", "nick_name", "phone", "role_id", "avatar", "sex").Save(&model).Error
	if err != nil {
		e.Log.Errorf("At Service ResetSysUserPwd error: %s", err)
		return err
	}
	return nil
}

// Remove 删除SysUser
func (e *SysUser) Remove(c *dto.SysUserById, p *actions.DataPermission, cb *casbin.SyncedEnforcer) error {
	var err error
	var data models.SysUser

	// 使用事务确保数据一致性
	tx := e.Orm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 先查询用户信息，获取用户名和角色信息
	err = tx.First(&data, c.GetId()).Error
	if err != nil {
		tx.Rollback()
		e.Log.Errorf("Error getting user info: %s", err)
		return err
	}

	// 删除用户与角色的关联关系
	err = tx.Where("user_id = ?", c.GetId()).Delete(&models.SysUserRole{}).Error
	if err != nil {
		tx.Rollback()
		e.Log.Errorf("Error deleting user-role relations: %s", err)
		return err
	}

	// 删除用户
	db := tx.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, c.GetId())
	if err = db.Error; err != nil {
		tx.Rollback()
		e.Log.Errorf("Error found in  RemoveSysUser : %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("无权删除该数据")
	}

	// 删除Casbin中的用户角色关联（g, user_$d, roleKey）
	if cb != nil {
		// 使用user_{userId}格式作为Casbin中的用户标识
		userSubject := fmt.Sprintf("user_%d", data.UserId)
		// 使用RemoveFilteredGroupingPolicy一次性删除该用户的所有角色关联
		// 参数0表示过滤第一个字段（用户标识），删除所有 g, user_$d, * 的策略
		_, err = cb.RemoveFilteredGroupingPolicy(0, userSubject)
		if err != nil {
			tx.Rollback()
			e.Log.Errorf("Error removing casbin user-role relations: %s", err)
			return err
		}
	}

	// 提交事务
	err = tx.Commit().Error
	if err != nil {
		e.Log.Errorf("Error committing transaction: %s", err)
		return err
	}

	return nil
}

// UpdatePwd 修改SysUser对象密码
func (e *SysUser) UpdatePwd(id int, oldPassword, newPassword string, p *actions.DataPermission) error {
	var err error

	if newPassword == "" {
		return nil
	}
	c := &models.SysUser{}

	err = e.Orm.Model(c).
		Scopes(
			actions.Permission(c.TableName(), p),
		).Select("UserId", "Password", "Salt").
		First(c, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("无权更新该数据")
		}
		e.Log.Errorf("db error: %s", err)
		return err
	}
	var ok bool
	ok, err = pkg.CompareHashAndPassword(c.Password, oldPassword)
	if err != nil {
		e.Log.Errorf("CompareHashAndPassword error, %s", err.Error())
		return err
	}
	if !ok {
		err = errors.New("incorrect Password")
		e.Log.Warnf("user[%d] %s", id, err.Error())
		return err
	}
	c.Password = newPassword
	db := e.Orm.Model(c).Where("user_id = ?", id).
		Select("Password", "Salt").
		Updates(c)
	if err = db.Error; err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		err = errors.New("set password error")
		log.Warnf("db update error")
		return err
	}
	return nil
}

func (e *SysUser) GetProfile(c *dto.SysUserById, user *models.SysUser, roles *[]models.SysRole, posts *[]models.SysPost) error {
	// 先查询用户基本信息
	err := e.Orm.First(user, c.GetId()).Error
	if err != nil {
		return err
	}

	// 手动加载部门信息，避免使用外键/关联查询
	if user.DeptId != 0 {
		var dept models.SysDept
		if err = e.Orm.First(&dept, user.DeptId).Error; err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}
		} else {
			user.Dept = &dept
		}
	}

	// 加载角色和岗位信息
	// 查询用户角色关联关系
	var userRoles []models.SysUserRole
	err = e.Orm.Where("user_id = ?", user.UserId).Find(&userRoles).Error
	if err != nil {
		return err
	}

	// 获取角色ID列表
	roleIds := make([]int, 0, len(userRoles))
	for _, ur := range userRoles {
		roleIds = append(roleIds, ur.RoleId)
	}

	// 查询角色信息
	if len(roleIds) > 0 {
		err = e.Orm.Find(roles, roleIds).Error
		if err != nil {
			return err
		}
	}

	// 查询岗位信息
	err = e.Orm.Find(posts, user.PostIds).Error
	if err != nil {
		return err
	}

	return nil
}

// SetUserRole 设置用户角色
func (e *SysUser) SetUserRole(c *dto.SysUserRoleReq, cb *casbin.SyncedEnforcer) error {
	var roles []models.SysRole

	// ========== 阶段1：事务内完成所有数据库操作 ==========
	err := e.Orm.Transaction(func(tx *gorm.DB) error {
		// 1. 检查用户是否存在
		var user models.SysUser
		if err := tx.First(&user, c.UserId).Error; err != nil {
			e.Log.Errorf("User not found: %s", err)
			return err
		}

		// 2. 如果有新的角色，先验证角色是否存在
		if len(c.RoleIds) > 0 {
			if err := tx.Where("role_id in ?", c.RoleIds).Find(&roles).Error; err != nil {
				e.Log.Errorf("Query roles error: %s", err)
				return err
			}

			if len(roles) != len(c.RoleIds) {
				e.Log.Errorf("Some roles not found")
				return errors.New("部分角色不存在")
			}
		}

		// 3. 删除用户旧的角色关系
		if err := tx.Where("user_id = ?", c.UserId).Delete(&models.SysUserRole{}).Error; err != nil {
			e.Log.Errorf("Delete old user-role relations error: %s", err)
			return err
		}

		// 4. 创建新的用户角色关系
		if len(c.RoleIds) > 0 {
			userRoles := make([]models.SysUserRole, 0, len(c.RoleIds))
			for _, roleId := range c.RoleIds {
				ur := models.SysUserRole{
					UserId: c.UserId,
					RoleId: roleId,
				}
				ur.SetCreateBy(c.UpdateBy)
				ur.SetUpdateBy(c.UpdateBy)
				userRoles = append(userRoles, ur)
			}

			if err := tx.Create(&userRoles).Error; err != nil {
				e.Log.Errorf("Create user-role relations error: %s", err)
				return err
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	// ========== 阶段2：事务已提交，安全地同步 Casbin ==========
	if cb != nil {
		userSubject := fmt.Sprintf("user_%d", c.UserId)

		// 删除旧的用户角色关联
		if _, err := cb.RemoveFilteredGroupingPolicy(0, userSubject); err != nil {
			e.Log.Errorf("Remove casbin user-role relations error: %s", err)
			return err
		}

		// 添加新的用户角色关联
		for _, role := range roles {
			if _, err := cb.AddGroupingPolicy(userSubject, role.RoleKey); err != nil {
				e.Log.Errorf("Add casbin user-role relation error: %s", err)
				return err
			}
		}
	}

	return nil
}
