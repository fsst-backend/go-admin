package service

import (
	"errors"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	cDto "go-admin/common/dto"

	"github.com/go-admin-team/go-admin-core/sdk/service"
)

type SysRolePermission struct {
	service.Service
}

// GetPage 获取 SysRolePermission 列表
func (e *SysRolePermission) GetPage(c *dto.SysRolePermissionGetPageReq, list *[]models.SysRolePermission, count *int64) error {
	err := e.Orm.
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.PaginateOffsetLimit(c.GetLimit(), c.GetOffset()),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("Service GetSysRolePermissionPage error:%s", err)
		return err
	}
	return nil
}

// Get 获取单个 SysRolePermission 对象
func (e *SysRolePermission) Get(d *dto.SysRolePermissionGetReq, model *models.SysRolePermission) error {
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
		e.Log.Errorf("Service GetSysRolePermission error: %s", err)
		_ = e.AddError(err)
		return err
	}
	return nil
}

// Insert 创建 SysRolePermission 对象
func (e *SysRolePermission) Insert(c *dto.SysRolePermissionInsertReq) error {
	var data models.SysRolePermission
	c.Generate(&data)
	if err := e.Orm.Create(&data).Error; err != nil {
		e.Log.Errorf("Service InsertSysRolePermission error:%s", err)
		return err
	}
	return nil
}

// Update 修改 SysRolePermission 对象
func (e *SysRolePermission) Update(c *dto.SysRolePermissionUpdateReq) error {
	var model models.SysRolePermission
	if err := e.Orm.First(&model, c.GetId()).Error; err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	c.Generate(&model)
	db := e.Orm.Save(&model)
	if err := db.Error; err != nil {
		e.Log.Errorf("Service UpdateSysRolePermission error:%s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除 SysRolePermission
func (e *SysRolePermission) Remove(d *dto.SysRolePermissionDeleteReq) error {
	var data models.SysRolePermission
	db := e.Orm.Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveSysRolePermission error:%s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
