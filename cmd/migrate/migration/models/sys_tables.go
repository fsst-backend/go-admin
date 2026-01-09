package models

import "go-admin/common/models"

type SysTables struct {
	TableId             int    `gorm:"column:table_id;type:int;primaryKey;autoIncrement" json:"tableId"`                 //表编码
	TBName              string `gorm:"column:table_name;type:varchar(255);" json:"tableName"`                            //表名称
	TableComment        string `gorm:"column:table_comment;type:varchar(255);" json:"tableComment"`                      //表备注
	ClassName           string `gorm:"column:class_name;type:varchar(255);" json:"className"`                            //类名
	TplCategory         string `gorm:"column:tpl_category;type:varchar(255);" json:"tplCategory"`                        //
	PackageName         string `gorm:"column:package_name;type:varchar(255);" json:"packageName"`                        //包名
	ModuleName          string `gorm:"column:module_name;type:varchar(255);" json:"moduleName"`                          //go文件名
	ModuleFrontName     string `gorm:"column:module_front_name;type:varchar(255);comment:前端文件名;" json:"moduleFrontName"` //前端文件名
	BusinessName        string `gorm:"column:business_name;type:varchar(255);" json:"businessName"`                      //
	FunctionName        string `gorm:"column:function_name;type:varchar(255);" json:"functionName"`                      //功能名称
	FunctionAuthor      string `gorm:"column:function_author;type:varchar(255);" json:"functionAuthor"`                  //功能作者
	PkColumn            string `gorm:"column:pk_column;type:varchar(255);" json:"pkColumn"`
	PkGoField           string `gorm:"column:pk_go_field;type:varchar(255);" json:"pkGoField"`
	PkJsonField         string `gorm:"column:pk_json_field;type:varchar(255);" json:"pkJsonField"`
	Options             string `gorm:"column:options;type:varchar(255);" json:"options"`
	TreeCode            string `gorm:"column:tree_code;type:varchar(255);" json:"treeCode"`
	TreeParentCode      string `gorm:"column:tree_parent_code;type:varchar(255);" json:"treeParentCode"`
	TreeName            string `gorm:"column:tree_name;type:varchar(255);" json:"treeName"`
	Tree                bool   `gorm:"column:tree;type:tinyint;default:0;" json:"tree"`
	Crud                bool   `gorm:"column:crud;type:tinyint;default:1;" json:"crud"`
	Remark              string `gorm:"column:remark;type:varchar(255);" json:"remark"`
	IsDataScope         int    `gorm:"column:is_data_scope;type:tinyint;" json:"isDataScope"`
	IsActions           int    `gorm:"column:is_actions;type:tinyint;" json:"isActions"`
	IsAuth              int    `gorm:"column:is_auth;type:tinyint;" json:"isAuth"`
	IsLogicalDelete     string `gorm:"column:is_logical_delete;type:varchar(1);" json:"isLogicalDelete"`
	LogicalDelete       bool   `gorm:"column:logical_delete;type:tinyint;" json:"logicalDelete"`
	LogicalDeleteColumn string `gorm:"column:logical_delete_column;type:varchar(128);" json:"logicalDeleteColumn"`
	models.ModelTime
	models.ControlBy
}

func (SysTables) TableName() string {
	return "sys_tables"
}
