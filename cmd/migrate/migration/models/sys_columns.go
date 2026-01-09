package models

import "go-admin/common/models"

type SysColumns struct {
	ColumnId           int    `gorm:"column:column_id;type:int;primaryKey;autoIncrement" json:"columnId"`
	TableId            int    `gorm:"column:table_id;type:int;" json:"tableId"`
	ColumnName         string `gorm:"column:column_name;type:varchar(128);" json:"columnName"`
	ColumnComment      string `gorm:"column:column_comment;type:varchar(128);" json:"columnComment"`
	ColumnType         string `gorm:"column:column_type;type:varchar(128);" json:"columnType"`
	GoType             string `gorm:"column:go_type;type:varchar(128);" json:"goType"`
	GoField            string `gorm:"column:go_field;type:varchar(128);" json:"goField"`
	JsonField          string `gorm:"column:json_field;type:varchar(128);" json:"jsonField"`
	IsPk               string `gorm:"column:is_pk;type:varchar(4);" json:"isPk"`
	IsIncrement        string `gorm:"column:is_increment;type:varchar(4);" json:"isIncrement"`
	IsRequired         string `gorm:"column:is_required;type:varchar(4);" json:"isRequired"`
	IsInsert           string `gorm:"column:is_insert;type:varchar(4);" json:"isInsert"`
	IsEdit             string `gorm:"column:is_edit;type:varchar(4);" json:"isEdit"`
	IsList             string `gorm:"column:is_list;type:varchar(4);" json:"isList"`
	IsQuery            string `gorm:"column:is_query;type:varchar(4);" json:"isQuery"`
	QueryType          string `gorm:"column:query_type;type:varchar(128);" json:"queryType"`
	HtmlType           string `gorm:"column:html_type;type:varchar(128);" json:"htmlType"`
	DictType           string `gorm:"column:dict_type;type:varchar(128);" json:"dictType"`
	Sort               int    `gorm:"column:sort;type:int;" json:"sort"`
	List               string `gorm:"column:list;type:varchar(1);" json:"list"`
	Pk                 bool   `gorm:"column:pk;type:tinyint;" json:"pk"`
	Required           bool   `gorm:"column:required;type:tinyint;" json:"required"`
	SuperColumn        bool   `gorm:"column:super_column;type:tinyint;" json:"superColumn"`
	UsableColumn       bool   `gorm:"column:usable_column;type:tinyint;" json:"usableColumn"`
	Increment          bool   `gorm:"column:increment;type:tinyint;" json:"increment"`
	Insert             bool   `gorm:"column:insert;type:tinyint;" json:"insert"`
	Edit               bool   `gorm:"column:edit;type:tinyint;" json:"edit"`
	Query              bool   `gorm:"column:query;type:tinyint;" json:"query"`
	Remark             string `gorm:"column:remark;type:varchar(255);" json:"remark"`
	FkTableName        string `gorm:"column:fk_table_name;type:varchar(255);" json:"fkTableName"`
	FkTableNameClass   string `gorm:"column:fk_table_name_class;type:varchar(255);" json:"fkTableNameClass"`
	FkTableNamePackage string `gorm:"column:fk_table_name_package;type:varchar(255);" json:"fkTableNamePackage"`
	FkLabelId          string `gorm:"column:fk_label_id;type:varchar(255);" json:"fkLabelId"`
	FkLabelName        string `gorm:"column:fk_label_name;type:varchar(255);" json:"fkLabelName"`
	models.ModelTime
	models.ControlBy
}

func (SysColumns) TableName() string {
	return "sys_columns"
}
