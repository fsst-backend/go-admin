package models

import (
	"strings"

	"go-admin/common/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type SysUser struct {
	UserId   int    `gorm:"column:user_id;type:int;primaryKey;autoIncrement;comment:编码"  json:"userId"`
	UUID     string `json:"uuid" gorm:"column:uuid;type:varchar(255);comment:UUID"`
	Username string `json:"username" gorm:"column:username;type:varchar(64);comment:用户名"`
	Password string `json:"-" gorm:"column:password;type:varchar(128);comment:密码"`
	NickName string `json:"nickName" gorm:"column:nick_name;type:varchar(128);comment:昵称"`
	Phone    string `json:"phone" gorm:"column:phone;type:varchar(11);comment:手机号"`
	Salt     string `json:"-" gorm:"column:salt;type:varchar(255);comment:加盐"`
	Avatar   string `json:"avatar" gorm:"column:avatar;type:varchar(255);comment:头像"`
	Sex      string `json:"sex" gorm:"column:sex;type:varchar(255);comment:性别"`
	Email    string `json:"email" gorm:"column:email;type:varchar(128);comment:邮箱"`
	DeptId   int    `json:"deptId" gorm:"column:dept_id;type:int;comment:部门"`
	PostId   int    `json:"postId" gorm:"column:post_id;type:int;comment:岗位"`
	Remark       string `json:"remark" gorm:"column:remark;type:varchar(255);comment:备注"`
	Status       int    `json:"status" gorm:"column:status;type:tinyint;comment:状态"`
	TokenVersion int    `json:"-" gorm:"column:token_version;type:int;default:0;comment:登录令牌版本，新登录递增后旧token失效"`
	models.ControlBy
	models.ModelTime
}

func (*SysUser) TableName() string {
	return "sys_user"
}

// Encrypt 加密
func (e *SysUser) Encrypt() (err error) {
	if e.Password == "" {
		return nil
	}

	// 已经是 bcrypt hash，直接跳过
	if strings.HasPrefix(e.Password, "$2a$") ||
		strings.HasPrefix(e.Password, "$2b$") ||
		strings.HasPrefix(e.Password, "$2y$") {
		return nil
	}

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(e.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	e.Password = string(hash)
	return nil
}

func (e *SysUser) BeforeCreate(_ *gorm.DB) error {
	return e.Encrypt()
}

func (e *SysUser) BeforeUpdate(_ *gorm.DB) error {
	var err error
	if e.Password != "" {
		err = e.Encrypt()
	}
	return err
}
