package models

// CasbinRule sys_casbin_rule
type CasbinRule struct {
	ID    uint   `gorm:"primaryKey"`
	Ptype string `gorm:"size:32;uniqueIndex:idx_casbin"`
	V0    string `gorm:"size:64;uniqueIndex:idx_casbin"`  // subject
	V1    string `gorm:"size:128;uniqueIndex:idx_casbin"` // object (url / resource)
	V2    string `gorm:"size:32;uniqueIndex:idx_casbin"`  // action
	V3    string `gorm:"size:64;uniqueIndex:idx_casbin"`
	V4    string `gorm:"size:64;uniqueIndex:idx_casbin"`
	V5    string `gorm:"size:64;uniqueIndex:idx_casbin"`
}

func (CasbinRule) TableName() string {
	return "sys_casbin_rule"
}
