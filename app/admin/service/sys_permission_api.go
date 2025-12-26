package service

import (
	"errors"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	cDto "go-admin/common/dto"

	"github.com/go-admin-team/go-admin-core/sdk/service"
)

type SysPermissionApi struct {
	service.Service
}

// GetPage 获取 SysPermissionApi 列表
func (e *SysPermissionApi) GetPage(c *dto.SysPermissionApiGetPageReq, list *[]models.SysPermissionApi, count *int64) error {
	err := e.Orm.
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.PaginateOffsetLimit(c.GetLimit(), c.GetOffset()),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("Service GetSysPermissionApiPage error:%s", err)
		return err
	}
	return nil
}

// Get 获取单个 SysPermissionApi 对象
func (e *SysPermissionApi) Get(d *dto.SysPermissionApiGetReq, model *models.SysPermissionApi) error {
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
		e.Log.Errorf("Service GetSysPermissionApi error: %s", err)
		_ = e.AddError(err)
		return err
	}
	return nil
}

// Insert 创建 SysPermissionApi 对象
func (e *SysPermissionApi) Insert(c *dto.SysPermissionApiInsertReq) error {
	var data models.SysPermissionApi
	c.Generate(&data)
	if err := e.Orm.Create(&data).Error; err != nil {
		e.Log.Errorf("Service InsertSysPermissionApi error:%s", err)
		return err
	}
	return nil
}

// Update 修改 SysPermissionApi 对象
func (e *SysPermissionApi) Update(c *dto.SysPermissionApiUpdateReq) error {
	var model models.SysPermissionApi
	if err := e.Orm.First(&model, c.GetId()).Error; err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	c.Generate(&model)
	db := e.Orm.Save(&model)
	if err := db.Error; err != nil {
		e.Log.Errorf("Service UpdateSysPermissionApi error:%s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除 SysPermissionApi
func (e *SysPermissionApi) Remove(d *dto.SysPermissionApiDeleteReq) error {
	var data models.SysPermissionApi
	db := e.Orm.Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveSysPermissionApi error:%s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
