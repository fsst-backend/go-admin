package apis

import (
	"errors"
	"go-admin/app/admin/models"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	"go-admin/common/global"
	"go-admin/common/mycasbin"
)

type SysUser struct {
	api.Api
}

// GetPage
// @Summary 列表用户信息数据
// @Description 获取JSON
// @Tags 用户
// @Accept application/json
// @Product application/json
// @Param username query string false "用户名"
// @Param nickName query string false "昵称"
// @Param phone query string false "手机号"
// @Param status query string false "状态"
// @Param limit query int false "页条数"
// @Param offset query int false "页码"
// @Success 200 {string} {object} response.Response{message=[]models.SysUser} "{"code": 0, "message": [...]}"
// @Router /lotus/api/v1/sys-user [get]
// @Security Bearer
func (e SysUser) GetPage(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserGetPageReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	list := make([]models.SysUser, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.PageOK(list, int(count), req.GetOffset(), req.GetLimit(), "查询成功")
}

// Get
// @Summary 获取用户
// @Description 获取JSON
// @Tags 用户
// @Success 200 {object} response.Response{message=models.SysUser} "{"code": 0, "message": [...]}
// @Router /lotus/api/v1/sys-user/get [get]
// @Security Bearer
func (e SysUser) Get(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserById{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.SysUser
	//数据权限检查
	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(http.StatusUnprocessableEntity, err, "查询失败")
		return
	}
	e.OK(object, "查询成功")
}

// Insert
// @Summary 创建用户
// @Description 获取JSON
// @Tags 用户
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysUserInsertReq true "用户数据"
// @Success 200 {object} response.Response{message=string} "{"code": 0, "message": [...]}"
// @Router /lotus/api/v1/sys-user [post]
// @Security Bearer
func (e SysUser) Insert(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserInsertReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	if err := CheckUsername(req.Username); err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	var hash []byte
	// 密码加密
	if hash, err = bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost); err != nil {
		req.Password = string(hash)
	}

	// 获取Casbin enforcer
	cb := sdk.Runtime.GetCasbinKey(c.Request.Host)

	err = s.Insert(&req, cb)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update
// @Summary 修改用户数据
// @Description 获取JSON
// @Tags 用户
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysUserUpdateReq true "body"
// @Success 200 {object} response.Response{message=string} "{"code": 0, "message": [...]}"
// @Router /lotus/api/v1/sys-user [put]
// @Security Bearer
func (e SysUser) Update(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserUpdateReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	if err := CheckUsername(req.Username); err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req.SetUpdateBy(user.GetUserId(c))

	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req.GetId(), "更新成功")
}

// Delete
// @Summary 删除用户数据
// @Description 删除数据
// @Tags 用户
// @Success 200 {object} response.Response{message=string} "{"code": 0, "message": [...]}
// @Router /lotus/api/v1/sys-user [delete]
// @Security Bearer
func (e SysUser) Delete(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserById{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// 设置编辑人
	req.SetUpdateBy(user.GetUserId(c))

	// 数据权限检查
	p := actions.GetPermissionFromContext(c)

	// 获取Casbin enforcer
	cb := sdk.Runtime.GetCasbinKey(c.Request.Host)

	err = s.Remove(&req, p, cb)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req.GetId(), "删除成功")
}

// InsetAvatar
// @Summary 修改头像
// @Description 获取JSON
// @Tags 个人中心
// @Accept application/json
// @Param data body dto.UpdateSysUserAvatarReq true "body"
// @Success 200 {object} response.Response{message=string} "{"code": 0, "message": [...]}}"
// @Router /lotus/api/v1/user/avatar [post]
// @Security Bearer
func (e SysUser) InsetAvatar(c *gin.Context) {
	s := service.SysUser{}
	req := dto.UpdateSysUserAvatarReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// 数据权限检查
	p := actions.GetPermissionFromContext(c)
	// 使用当前登录用户ID更新头像
	currentUserId := user.GetUserId(c)
	err = s.UpdateAvatar(currentUserId, &req, p)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req.Avatar, "修改成功")
}

// UpdateStatus 修改用户状态
// @Summary 修改用户状态
// @Description 获取JSON
// @Tags 用户
// @Accept  application/json
// @Product application/json
// @Param data body dto.UpdateSysUserStatusReq true "body"
// @Success 200 {object} response.Response{message=string} "{"code": 0, "message": [...]}"
// @Router /lotus/api/v1/user/status [put]
// @Security Bearer
func (e SysUser) UpdateStatus(c *gin.Context) {
	s := service.SysUser{}
	req := dto.UpdateSysUserStatusReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req.SetUpdateBy(user.GetUserId(c))

	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	err = s.UpdateStatus(&req, p)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req.GetId(), "更新成功")
}

// ResetPwd 重置用户密码
// @Summary 重置用户密码
// @Description 获取JSON
// @Tags 用户
// @Accept  application/json
// @Product application/json
// @Param data body dto.ResetSysUserPwdReq true "body"
// @Success 200 {object} response.Response{message=string} "{"code": 0, "message": [...]}"
// @Router /lotus/api/v1/user/pwd/reset [put]
// @Security Bearer
func (e SysUser) ResetPwd(c *gin.Context) {
	s := service.SysUser{}
	req := dto.ResetSysUserPwdReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req.SetUpdateBy(user.GetUserId(c))

	//数据权限检查
	p := actions.GetPermissionFromContext(c)
	var hash []byte
	if hash, err = bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost); err != nil {
		req.Password = string(hash)
	}

	err = s.ResetPwd(&req, p)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req.GetId(), "更新成功")
}

// UpdatePwd
// @Summary 修改密码
// @Description 获取JSON
// @Tags 用户
// @Accept  application/json
// @Product application/json
// @Param data body dto.PassWord true "body"
// @Success 200 {object} response.Response{message=string} "{"code": 0, "message": [...]}"
// @Router /lotus/api/v1/user/pwd/set [put]
// @Security Bearer
func (e SysUser) UpdatePwd(c *gin.Context) {
	s := service.SysUser{}
	req := dto.PassWord{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// 数据权限检查
	p := actions.GetPermissionFromContext(c)
	var hash []byte
	if hash, err = bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost); err != nil {
		req.NewPassword = string(hash)
	}

	err = s.UpdatePwd(user.GetUserId(c), req.OldPassword, req.NewPassword, p)
	if err != nil {
		e.Logger.Error(err)
		e.Error(http.StatusForbidden, err, "密码修改失败")
		return
	}

	e.OK(nil, "密码修改成功")
}

// GetProfile
// @Summary 获取个人中心用户
// @Description 获取JSON
// @Tags 个人中心
// @Success 200 {object} response.Response{message=gin.H} "{"code": 0, "message": [...]}"
// @Router /lotus/api/v1/user/profile [get]
// @Security Bearer
func (e SysUser) GetProfile(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserById{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req.Id = user.GetUserId(c)

	sysUser := models.SysUser{}
	roles := make([]models.SysRole, 0)
	posts := make([]models.SysPost, 0)
	err = s.GetProfile(&req, &sysUser, &roles, &posts)
	if err != nil {
		e.Logger.Errorf("get user profile error, %s", err.Error())
		e.Error(500, err, "获取用户信息失败")
		return
	}
	e.OK(gin.H{
		"user":  sysUser,
		"roles": roles,
		"posts": posts,
	}, "查询成功")
}

// GetInfo
// @Summary 获取个人信息
// @Description 获取JSON
// @Tags 个人中心
// @Success 200 {object} response.Response{message=dto.SysUserInfoResp} "{"code": 0, "message": [...]}"
// @Router /lotus/api/v1/getinfo [get]
// @Security Bearer
func (e SysUser) GetInfo(c *gin.Context) {
	req := dto.SysUserById{}
	s := service.SysUser{}
	r := service.SysRole{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&r.Service).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	p := actions.GetPermissionFromContext(c)
	req.Id = user.GetUserId(c)

	// 获取用户信息
	sysUser := models.SysUser{}
	err = s.Get(&req, p, &sysUser)
	if err != nil {
		e.Error(http.StatusUnauthorized, err, "登录失败")
		return
	}

	// 查询用户的角色
	var userRoles []models.SysUserRole
	err = e.Orm.Where("user_id = ?", sysUser.UserId).Find(&userRoles).Error
	if err != nil {
		e.Error(http.StatusInternalServerError, err, "查询用户角色失败")
		return
	}

	// 查询角色详细信息
	roles := make([]string, 0)
	var permissions []string
	var buttons []string
	isSuperAdmin := false

	if len(userRoles) > 0 {
		roleIds := make([]int, 0, len(userRoles))
		for _, ur := range userRoles {
			roleIds = append(roleIds, ur.RoleId)
		}

		var roleList []models.SysRole
		err = e.Orm.Where("role_id in ?", roleIds).Find(&roleList).Error
		if err != nil {
			e.Error(http.StatusInternalServerError, err, "查询角色信息失败")
			return
		}

		// 收集角色名称和检查是否为 SuperAdmin
		for _, role := range roleList {
			roles = append(roles, role.RoleName)
			if role.RoleKey == mycasbin.SuperAdmin {
				isSuperAdmin = true
			}
		}
	}

	// 设置权限
	if isSuperAdmin {
		// SuperAdmin 拥有所有权限
		permissions = []string{"*:*:*"}
		buttons = []string{"*:*:*"}
	} else {
		// 普通用户:合并所有角色的权限
		permissionMap := make(map[string]bool)
		for _, ur := range userRoles {
			list, _ := r.GetPremissonByRoleId(ur.RoleId, c.Request.Host)
			for _, perm := range list {
				permissionMap[perm] = true
			}
		}

		// 将 map 转换为切片
		permissions = make([]string, 0, len(permissionMap))
		for perm := range permissionMap {
			permissions = append(permissions, perm)
		}
		buttons = permissions // buttons 和 permissions 保持一致
	}

	// 构建响应数据
	resp := dto.SysUserInfoResp{
		Roles:        roles,
		Permissions:  permissions,
		Buttons:      buttons,
		Introduction: "I am a super administrator",
		Avatar:       "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
		UserName:     sysUser.Username,
		UserId:       sysUser.UserId,
		UUID:         sysUser.UUID,
		DeptId:       sysUser.DeptId,
		Alias:        sysUser.NickName,
		Phone:        sysUser.Phone,
		Email:        sysUser.Email,
	}
	if sysUser.Avatar != "" {
		resp.Avatar = sysUser.Avatar
	}
	e.OK(resp, "")
}

// SetUserRole
// @Summary 设置用户角色
// @Description 为用户分配角色
// @Tags 用户
// @Accept application/json
// @Product application/json
// @Param data body dto.SysUserRoleReq true "用户角色授权请求"
// @Success 200 {object} response.Response{message=string} "{"code": 0, "message": [...]}"
// @Router /lotus/api/v1/sys-user/role [put]
// @Security Bearer
func (e SysUser) SetUserRole(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserRoleReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// 设置操作人
	req.SetUpdateBy(user.GetUserId(c))

	// 获取Casbin enforcer
	cb := sdk.Runtime.GetCasbinKey(c.Request.Host)

	err = s.SetUserRole(&req, cb)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, "设置用户角色失败")
		return
	}
	_, err = global.LoadPolicy(c)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, "设置用户角色成功，但刷新策略失败")
		return
	}
	e.OK(req.UserId, "设置用户角色成功")
}

func CheckUsername(username string) error {
	// 长度校验
	if len(username) < 4 || len(username) > 32 {
		return errors.New("用户名长度必须在 4~32 位之间")
	}

	// 字符范围校验
	allowed := regexp.MustCompile(`^[a-zA-Z0-9@_.]+$`)
	if !allowed.MatchString(username) {
		return errors.New("用户名只能包含字母、数字、@、_、.")
	}

	// 不能以 . 或 _ 开头或结尾
	if username[0] == '.' || username[0] == '_' ||
		username[len(username)-1] == '.' || username[len(username)-1] == '_' {
		return errors.New("用户名不能以 . 或 _ 开头或结尾")
	}

	// 不允许连续符号
	if regexp.MustCompile(`(\.\.|__)`).MatchString(username) {
		return errors.New("用户名不能包含连续的 . 或 _")
	}

	return nil
}
