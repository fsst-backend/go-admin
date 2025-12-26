package dto

import (
	"go-admin/app/admin/models"
	common "go-admin/common/models"

	"go-admin/common/dto"
)

// SysMenuGetPageReq 列表或者搜索使用结构体
type SysMenuGetPageReq struct {
	dto.OffsetLimitPagination `search:"-"`  // 使用 offset/limit 分页
	Title          string `form:"title" search:"type:contains;column:title;table:sys_menu" comment:"菜单名称"` // 菜单名称
	Status         string `form:"status" search:"type:exact;column:status;table:sys_menu" comment:"状态"`    // 状态
}

func (m *SysMenuGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SysMenuInsertReq struct {
	MenuId         int    `json:"menuId" comment:"编码"`            // 编码
	MenuName     string `json:"menuName" comment:"菜单name"`   //菜单name
	Title        string `json:"title" comment:"显示标题"`        //显示标题
	MenuType     string `json:"menuType" comment:"菜单类型"`     //菜单类型
	MenuPath     string `json:"menuPath" comment:"菜单路径"`     //菜单类型 /0/1/7
	Path         string `json:"path" comment:"路径"`           //路径 /camellia/sys_user
	Perm         string `json:"perm" comment:"权限标识"`         //权限标识
	Component    string `json:"component" comment:"组件"`      //组件
	Icon         string `json:"icon" comment:"图标"`           //图标
	SortValue    int    `json:"sortValue" comment:"排序值"`     //排序值
	IsExternal   bool   `json:"isExternal" comment:"是否外部链接"` //是否外部链接
	ExternalLink string `json:"externalLink" comment:"外部链接"` //外部链接
	TextBadge    string `json:"textBadge" comment:"文本徽章"`    //文本徽章
	ActivePath   string `json:"activePath" comment:"激活路径"`   //激活路径
	Status       string `json:"status" comment:"状态"`         //状态
	KeepAlive    bool   `json:"keepAlive" comment:"是否缓存"`    //是否缓存
	IsHide       bool   `json:"isHide" comment:"是否隐藏"`       //是否隐藏
	IsIframe     bool   `json:"isIframe" comment:"是否iframe"` //是否iframe
	ShowBadge    bool   `json:"showBadge" comment:"是否显示徽章"`  //是否显示徽章
	FixedTab     bool   `json:"fixedTab" comment:"是否固定标签页"`  //是否固定标签页
	IsHideTab    bool   `json:"isHideTab" comment:"是否隐藏标签页"` //是否隐藏标签页
	IsFullPage   bool   `json:"isFullPage" comment:"是否全屏页面"` //是否全屏页面
	ParentId     int    `json:"parentId" comment:"上级菜单"`     //上级菜单
	PermissionCode string `json:"permissionCode" comment:"权限Code"`   //权限Code
	common.ControlBy
}

func (s *SysMenuInsertReq) Generate(model *models.SysMenu) {
	if s.MenuId != 0 {
		model.MenuId = s.MenuId
	}
	model.MenuName = s.MenuName
	model.Title = s.Title
	model.MenuType = s.MenuType
	model.MenuPath = s.MenuPath
	model.Path = s.Path
	model.Perm = s.Perm
	model.Component = s.Component
	model.Icon = s.Icon
	model.SortValue = s.SortValue
	model.IsExternal = s.IsExternal
	model.ExternalLink = s.ExternalLink
	model.TextBadge = s.TextBadge
	model.ActivePath = s.ActivePath
	model.Status = s.Status
	model.KeepAlive = s.KeepAlive
	model.IsHide = s.IsHide
	model.IsIframe = s.IsIframe
	model.ShowBadge = s.ShowBadge
	model.FixedTab = s.FixedTab
	model.IsHideTab = s.IsHideTab
	model.IsFullPage = s.IsFullPage
	model.ParentId = s.ParentId
	model.PermissionCode = s.PermissionCode
	if s.CreateBy != 0 {
		model.CreateBy = s.CreateBy
	}
	if s.UpdateBy != 0 {
		model.UpdateBy = s.UpdateBy
	}
}

func (s *SysMenuInsertReq) GetId() interface{} {
	return s.MenuId
}

type SysMenuUpdateReq struct {
	MenuId         int    `json:"menuId" comment:"编码"`            // 编码
	MenuName     string `json:"menuName" comment:"菜单name"`   //菜单name
	Title        string `json:"title" comment:"显示标题"`        //显示标题
	MenuType     string `json:"menuType" comment:"菜单类型"`     //菜单类型
	MenuPath     string `json:"menuPath" comment:"菜单类型"`     //菜单类型
	Path         string `json:"path" comment:"路径"`           //路径
	Perm         string `json:"perm" comment:"权限标识"`         //权限标识
	Component    string `json:"component" comment:"组件"`      //组件
	Icon         string `json:"icon" comment:"图标"`           //图标
	SortValue    int    `json:"sortValue" comment:"排序值"`     //排序值
	IsExternal   bool   `json:"isExternal" comment:"是否外部链接"` //是否外部链接
	ExternalLink string `json:"externalLink" comment:"外部链接"` //外部链接
	TextBadge    string `json:"textBadge" comment:"文本徽章"`    //文本徽章
	ActivePath   string `json:"activePath" comment:"激活路径"`   //激活路径
	Status       string `json:"status" comment:"状态"`         //状态
	KeepAlive    bool   `json:"keepAlive" comment:"是否缓存"`    //是否缓存
	IsHide       bool   `json:"isHide" comment:"是否隐藏"`       //是否隐藏
	IsIframe     bool   `json:"isIframe" comment:"是否iframe"` //是否iframe
	ShowBadge    bool   `json:"showBadge" comment:"是否显示徽章"`  //是否显示徽章
	FixedTab     bool   `json:"fixedTab" comment:"是否固定标签页"`  //是否固定标签页
	IsHideTab    bool   `json:"isHideTab" comment:"是否隐藏标签页"` //是否隐藏标签页
	IsFullPage   bool   `json:"isFullPage" comment:"是否全屏页面"` //是否全屏页面
	ParentId     int    `json:"parentId" comment:"上级菜单"`     //上级菜单
	PermissionCode string `json:"permissionCode" comment:"权限Code"`   //权限Code
	common.ControlBy
}

func (s *SysMenuUpdateReq) Generate(model *models.SysMenu) {
	if s.MenuId != 0 {
		model.MenuId = s.MenuId
	}
	model.MenuName = s.MenuName
	model.Title = s.Title
	model.MenuType = s.MenuType
	model.MenuPath = s.MenuPath
	model.Path = s.Path
	model.Perm = s.Perm
	model.Component = s.Component
	model.Icon = s.Icon
	model.SortValue = s.SortValue
	model.IsExternal = s.IsExternal
	model.ExternalLink = s.ExternalLink
	model.TextBadge = s.TextBadge
	model.ActivePath = s.ActivePath
	model.Status = s.Status
	model.KeepAlive = s.KeepAlive
	model.IsHide = s.IsHide
	model.IsIframe = s.IsIframe
	model.ShowBadge = s.ShowBadge
	model.FixedTab = s.FixedTab
	model.IsHideTab = s.IsHideTab
	model.IsFullPage = s.IsFullPage
	model.ParentId = s.ParentId
	model.PermissionCode = s.PermissionCode
	if s.CreateBy != 0 {
		model.CreateBy = s.CreateBy
	}
	if s.UpdateBy != 0 {
		model.UpdateBy = s.UpdateBy
	}
}

func (s *SysMenuUpdateReq) GetId() interface{} {
	return s.MenuId
}

type SysMenuGetReq struct {
	Id int `form:"id"`
}

func (s *SysMenuGetReq) GetId() interface{} {
	return s.Id
}

type SysMenuDeleteReq struct {
	Ids []int `json:"ids"`
	common.ControlBy
}

func (s *SysMenuDeleteReq) GetId() interface{} {
	return s.Ids
}

type MenuLabel struct {
	Id       int         `json:"id,omitempty" gorm:"-"`
	Label    string      `json:"label,omitempty" gorm:"-"`
	Children []MenuLabel `json:"children,omitempty" gorm:"-"`
}

type MenuRole struct {
	models.SysMenu
	IsSelect bool `json:"is_select" gorm:"-"`
}

type SelectRole struct {
	RoleId int `uri:"roleId"`
}
