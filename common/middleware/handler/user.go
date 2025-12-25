package handler

import (
	"go-admin/common/models"

	"gorm.io/gorm"
)

type SysUser struct {
	UserId   int    `gorm:"primaryKey;autoIncrement;comment:编码"  json:"userId"`
	UUID     string `json:"uuid" gorm:"size:255;comment:用户UUID"`
	Username string `json:"username" gorm:"size:64;comment:用户名"`
	Password string `json:"-" gorm:"size:128;comment:密码"`
	NickName string `json:"nickName" gorm:"size:128;comment:昵称"`
	Phone    string `json:"phone" gorm:"size:11;comment:手机号"`
	Salt     string `json:"-" gorm:"size:255;comment:加盐"`
	Avatar   string `json:"avatar" gorm:"size:255;comment:头像"`
	Sex      string `json:"sex" gorm:"size:255;comment:性别"`
	Email    string `json:"email" gorm:"size:128;comment:邮箱"`
	DeptId   int    `json:"deptId" gorm:"size:20;comment:部门"`
	PostId   int    `json:"postId" gorm:"size:20;comment:岗位"`
	Remark   string `json:"remark" gorm:"size:255;comment:备注"`
	Status   string `json:"status" gorm:"size:4;comment:状态"`
	DeptIds  []int  `json:"deptIds" gorm:"-"`
	PostIds  []int  `json:"postIds" gorm:"-"`
	RoleIds  []int  `json:"roleIds" gorm:"-"`
	//Dept     *SysDept `json:"dept"`
	models.ControlBy
	models.ModelTime
}

func (*SysUser) TableName() string {
	return "sys_user"
}

func (e *SysUser) AfterFind(_ *gorm.DB) error {
	e.DeptIds = []int{e.DeptId}
	e.PostIds = []int{e.PostId}
	return nil
}
