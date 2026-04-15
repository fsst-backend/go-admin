package models

// TbDemo 示例表，与历史 migrate 模型一致
type TbDemo struct {
	Model
	Name string `json:"name" gorm:"column:name;type:varchar(128);comment:名称"`
	ModelTime
	ControlBy
}

func (TbDemo) TableName() string {
	return "tb_demo"
}
