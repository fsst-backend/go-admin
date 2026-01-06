package models

import "go-admin/common/models"

type SysMenu struct {
	MenuId         int           `json:"menuId" gorm:"column:menu_id;type:int;primaryKey;autoIncrement"`
	MenuName       string        `json:"menuName" gorm:"column:menu_name;type:varchar(128);"`
	Title          string        `json:"title" gorm:"column:title;type:varchar(128);"`
	MenuType       string        `json:"menuType" gorm:"column:menu_type;type:varchar(10);"` // M: Menu/Directory   C: Component/Page  F: Function/Button
	MenuPath       string        `json:"menuPath" gorm:"column:menu_path;type:varchar(128);"`
	Path           string        `json:"path" gorm:"column:path;type:varchar(128);"`
	Perm           string        `json:"perm" gorm:"column:perm;type:varchar(255);"`
	Component      string        `json:"component" gorm:"column:component;type:varchar(255);"`
	Icon           string        `json:"icon" gorm:"column:icon;type:varchar(128);"`
	SortValue      int           `json:"sortValue" gorm:"column:sort_value;type:int;"`
	IsExternal     bool          `json:"isExternal" gorm:"column:is_external;type:tinyint;DEFAULT:0;"`
	ExternalLink   string        `json:"externalLink" gorm:"column:external_link;type:varchar(255);"`
	TextBadge      string        `json:"textBadge" gorm:"column:text_badge;type:varchar(10);"`
	ActivePath     string        `json:"activePath" gorm:"column:active_path;type:varchar(100);"`
	Status         string        `json:"status" gorm:"column:status;type:varchar(1);"`
	KeepAlive      bool          `json:"keepAlive" gorm:"column:keep_alive;type:tinyint;DEFAULT:1;"`
	IsHide         bool          `json:"isHide" gorm:"column:is_hide;type:tinyint;DEFAULT:0;"`
	IsIframe       bool          `json:"isIframe" gorm:"column:is_iframe;type:tinyint;DEFAULT:0;"`
	ShowBadge      bool          `json:"showBadge" gorm:"column:show_badge;type:tinyint;DEFAULT:1;"`
	FixedTab       bool          `json:"fixedTab" gorm:"column:fixed_tab;type:tinyint;DEFAULT:0;"`
	IsHideTab      bool          `json:"isHideTab" gorm:"column:is_hide_tab;type:tinyint;DEFAULT:0;"`
	IsFullPage     bool          `json:"isFullPage" gorm:"column:is_full_page;type:tinyint;DEFAULT:0;"`
	ParentId       int           `json:"parentId" gorm:"column:parent_id;type:int;comment:父菜单ID"`
	PermissionCode string        `json:"permissionCode" gorm:"column:permission_code;type:varchar(128);comment:关联权限表code"`
	Children       []SysMenu     `json:"children,omitempty" gorm:"-"`
	Permission     SysPermission `json:"permission,omitempty" gorm:"-"`
	models.ControlBy
	models.ModelTime
}

// SysMenu字段常量定义 - 用于GORM查询和函数调用
const (
	SysMenuMenuId       = "menu_id"
	SysMenuMenuName     = "menu_name"
	SysMenuTitle        = "title"
	SysMenuMenuType     = "menu_type"
	SysMenuMenuPath     = "menu_path"
	SysMenuPath         = "path"
	SysMenuPerm         = "perm"
	SysMenuComponent    = "component"
	SysMenuIcon         = "icon"
	SysMenuSortValue    = "sort_value"
	SysMenuIsExternal   = "is_external"
	SysMenuExternalLink = "external_link"
	SysMenuTextBadge    = "text_badge"
	SysMenuActivePath   = "active_path"
	SysMenuStatus       = "status"
	SysMenuKeepAlive    = "keep_alive"
	SysMenuIsHide       = "is_hide"
	SysMenuIsIframe     = "is_iframe"
	SysMenuShowBadge    = "show_badge"
	SysMenuFixedTab     = "fixed_tab"
	SysMenuIsHideTab    = "is_hide_tab"
	SysMenuIsFullPage   = "is_full_page"
	SysMenuParentId     = "parent_id"
)

type SysMenuSlice []SysMenu

func (x SysMenuSlice) Len() int           { return len(x) }
func (x SysMenuSlice) Less(i, j int) bool { return x[i].SortValue < x[j].SortValue }
func (x SysMenuSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func (*SysMenu) TableName() string {
	return "sys_menu"
}

func (e *SysMenu) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysMenu) GetId() interface{} {
	return e.MenuId
}
