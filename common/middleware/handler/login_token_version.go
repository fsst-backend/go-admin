package handler

import (
	cmodels "go-admin/common/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// bumpLoginTokenVersion 指定渠道单设备版本 +1（无行则插入为 1）
func bumpLoginTokenVersion(db *gorm.DB, userID int, loginChannel string) (newVer int, err error) {
	name := db.Dialector.Name()
	if name == "mysql" || name == "postgres" || name == "sqlite" {
		row := cmodels.SysUserLoginToken{
			UserID:       userID,
			LoginChannel: loginChannel,
			TokenVersion: 1,
		}
		err = db.Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "user_id"}, {Name: "login_channel"}},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"token_version": gorm.Expr("token_version + 1"),
			}),
		}).Create(&row).Error
		if err != nil {
			return 0, err
		}
	} else {
		err = db.Transaction(func(tx *gorm.DB) error {
			var cnt int64
			if err := tx.Model(&cmodels.SysUserLoginToken{}).
				Where("user_id = ? AND login_channel = ?", userID, loginChannel).
				Count(&cnt).Error; err != nil {
				return err
			}
			if cnt == 0 {
				return tx.Create(&cmodels.SysUserLoginToken{
					UserID:       userID,
					LoginChannel: loginChannel,
					TokenVersion: 1,
				}).Error
			}
			return tx.Model(&cmodels.SysUserLoginToken{}).
				Where("user_id = ? AND login_channel = ?", userID, loginChannel).
				Update("token_version", gorm.Expr("COALESCE(token_version,0) + 1")).Error
		})
		if err != nil {
			return 0, err
		}
	}
	err = db.Model(&cmodels.SysUserLoginToken{}).
		Where("user_id = ? AND login_channel = ?", userID, loginChannel).
		Select("token_version").
		Scan(&newVer).Error
	return newVer, err
}
