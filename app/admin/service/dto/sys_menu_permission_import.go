package dto

type MenuPermissionIO struct {
	Menus       []MenuIO       `json:"menus"`
	Permissions []PermissionIO `json:"permissions"`
}

type MenuIO struct {
	MenuType       string   `json:"menuType" comment:"菜单类型"`   //菜单类型
	Path           string   `json:"path" comment:"路径"`         //url /camellia/sys_user
	Component      string   `json:"component" comment:"路由"`    //
	Perm           string   `json:"perm" comment:"权限标识"`       //权限标识
	MenuName       string   `json:"menuName" comment:"菜单name"` //菜单name
	Title          string   `json:"title" comment:"显示标题"`      //显示标题
	PermissionCode string   `json:"permission_code"`
	Icon           string   `json:"icon"`
	SortValue      int      `json:"sortValue"`
	IsExternal     bool     `json:"isExternal"`
	ExternalLink   string   `json:"externalLink" `
	TextBadge      string   `json:"textBadge"`
	ActivePath     string   `json:"activePath"`
	Status         string   `json:"status"`
	KeepAlive      bool     `json:"keepAlive"`
	IsHide         bool     `json:"isHide"`
	IsIframe       bool     `json:"isIframe"`
	ShowBadge      bool     `json:"showBadge"`
	FixedTab       bool     `json:"fixedTab"`
	IsHideTab      bool     `json:"isHideTab"`
	IsFullPage     bool     `json:"isFullPage"`
	Children       []MenuIO `json:"children"`
}

type ApiIO struct {
	Method string `json:"method"`
	URL    string `json:"url"`
}

type PermissionIO struct {
	Code string  `json:"code"`
	Name string  `json:"name"`
	Type string  `json:"type"`
	Apis []ApiIO `json:"apis"`
}
