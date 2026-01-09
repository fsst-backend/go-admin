package models

import "go-admin/common/models"

type SysMenu struct {
	MenuId         int    `json:"menuId" gorm:"column:menu_id;type:int;primaryKey;autoIncrement"`
	MenuName       string `json:"menuName" gorm:"column:menu_name;type:varchar(128);"`
	Title          string `json:"title" gorm:"column:title;type:varchar(128);"`
	MenuType       string `json:"menuType" gorm:"column:menu_type;type:varchar(10);"` // M: Menu/Directory   C: Component/Page  F: Function/Button
	MenuPath       string `json:"menuPath" gorm:"column:menu_path;type:varchar(128);"`
	Path           string `json:"path" gorm:"column:path;type:varchar(128);"`
	Perm           string `json:"perm" gorm:"column:perm;type:varchar(255);"`
	Component      string `json:"component" gorm:"column:component;type:varchar(255);"`
	Icon           string `json:"icon" gorm:"column:icon;type:varchar(128);"`
	SortValue      int    `json:"sortValue" gorm:"column:sort_value;type:int;"`
	IsExternal     bool   `json:"isExternal" gorm:"column:is_external;type:tinyint;DEFAULT:0;"`
	ExternalLink   string `json:"externalLink" gorm:"column:external_link;type:varchar(255);"`
	TextBadge      string `json:"textBadge" gorm:"column:text_badge;type:varchar(10);"`
	ActivePath     string `json:"activePath" gorm:"column:active_path;type:varchar(100);"`
	Status         string `json:"status" gorm:"column:status;type:varchar(1);"`
	KeepAlive      bool   `json:"keepAlive" gorm:"column:keep_alive;type:tinyint;DEFAULT:1;"`
	IsHide         bool   `json:"isHide" gorm:"column:is_hide;type:tinyint;DEFAULT:0;"`
	IsIframe       bool   `json:"isIframe" gorm:"column:is_iframe;type:tinyint;DEFAULT:0;"`
	ShowBadge      bool   `json:"showBadge" gorm:"column:show_badge;type:tinyint;DEFAULT:1;"`
	FixedTab       bool   `json:"fixedTab" gorm:"column:fixed_tab;type:tinyint;DEFAULT:0;"`
	IsHideTab      bool   `json:"isHideTab" gorm:"column:is_hide_tab;type:tinyint;DEFAULT:0;"`
	IsFullPage     bool   `json:"isFullPage" gorm:"column:is_full_page;type:tinyint;DEFAULT:0;"`
	ParentId       int    `json:"parentId" gorm:"column:parent_id;type:int;comment:父菜单ID"`
	PermissionCode string `json:"permissionCode" gorm:"column:permission_code;type:varchar(128);comment:关联权限表code"`
	models.ControlBy
	models.ModelTime
}

func (SysMenu) TableName() string {
	return "sys_menu"
}
