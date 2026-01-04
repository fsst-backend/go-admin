package models

type SysMenu struct {
	MenuId         int    `json:"menuId" gorm:"column:menu_id;primaryKey;autoIncrement"`
	MenuName       string `json:"menuName" gorm:"column:menu_name;size:128;"`
	Title          string `json:"title" gorm:"column:title;size:128;"`
	MenuType       string `json:"menuType" gorm:"column:menu_type;size:10;"` // M: Menu/Directory   C: Component/Page  F: Function/Button
	MenuPath       string `json:"menuPath" gorm:"column:menu_path;size:128;"`
	Path           string `json:"path" gorm:"column:path;size:128;"`
	Perm           string `json:"perm" gorm:"column:perm;size:255;"`
	Component      string `json:"component" gorm:"column:component;size:255;"`
	Icon           string `json:"icon" gorm:"column:icon;size:128;"`
	SortValue      int    `json:"sortValue" gorm:"column:sort_value;size:4;"`
	IsExternal     bool   `json:"isExternal" gorm:"column:is_external;size:1;DEFAULT:0;"`
	ExternalLink   string `json:"externalLink" gorm:"column:external_link;size:255;"`
	TextBadge      string `json:"textBadge" gorm:"column:text_badge;size:10;"`
	ActivePath     string `json:"activePath" gorm:"column:active_path;size:100;"`
	Status         string `json:"status" gorm:"column:status;size:1;"`
	KeepAlive      bool   `json:"keepAlive" gorm:"column:keep_alive;size:1;DEFAULT:1;"`
	IsHide         bool   `json:"isHide" gorm:"column:is_hide;size:1;DEFAULT:0;"`
	IsIframe       bool   `json:"isIframe" gorm:"column:is_iframe;size:1;DEFAULT:0;"`
	ShowBadge      bool   `json:"showBadge" gorm:"column:show_badge;size:1;DEFAULT:1;"`
	FixedTab       bool   `json:"fixedTab" gorm:"column:fixed_tab;size:1;DEFAULT:0;"`
	IsHideTab      bool   `json:"isHideTab" gorm:"column:is_hide_tab;size:1;DEFAULT:0;"`
	IsFullPage     bool   `json:"isFullPage" gorm:"column:is_full_page;size:1;DEFAULT:0;"`
	ParentId       int    `json:"parentId" gorm:"column:parent_id;type:bigint;comment:父菜单ID"`
	PermissionCode string `json:"permissionCode" gorm:"column:permission_code;size:128;comment:关联权限表code"`
	ControlBy
	ModelTime
}

func (SysMenu) TableName() string {
	return "sys_menu"
}
