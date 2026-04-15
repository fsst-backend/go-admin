package handler

import (
	"strings"

	"go-admin/common/constant"

	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"gorm.io/gorm"
)

type Login struct {
	Username string `form:"UserName" json:"username" binding:"required"`
	Password string `form:"Password" json:"password" binding:"required"`
	Code     string `form:"Code" json:"code" binding:"required"`
	UUID     string `form:"UUID" json:"uuid" binding:"required"`
	// LoginChannel 登录渠道：空或 admin=后台；cs_app=客服 App（单设备与后台隔离）
	LoginChannel string `form:"loginChannel" json:"loginChannel"`
}

// NormalizeLoginChannel 归一化登录渠道：空串为后台；已知别名归一；其余小写原样（与新表行一一对应，加端无需改表）
func NormalizeLoginChannel(s string) string {
	s = strings.TrimSpace(strings.ToLower(s))
	if s == "" {
		return constant.LoginChannelAdmin
	}
	if s == constant.LoginChannelCSApp {
		return constant.LoginChannelCSApp
	}
	if s == constant.LoginChannelAdmin {
		return constant.LoginChannelAdmin
	}
	return s
}

func (u *Login) GetUser(tx *gorm.DB) (user SysUser, err error) {
	err = tx.Table("sys_user").Where("username = ?  and status = ?", u.Username, constant.UserStatusNormal).First(&user).Error
	if err != nil {
		log.Errorf("get user error, %s", err.Error())
		return
	}
	_, err = pkg.CompareHashAndPassword(user.Password, u.Password)
	if err != nil {
		log.Errorf("user login error, %s", err.Error())
		return
	}
	return
}
