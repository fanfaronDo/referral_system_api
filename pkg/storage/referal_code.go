package storage

import (
	"github.com/fanfaronDo/referral_system_api/pkg/model"
	"gorm.io/gorm"
)

type ReferralCode struct {
	db *gorm.DB
}

func NewReferralCode(db *gorm.DB) *ReferralCode {
	return &ReferralCode{
		db: db,
	}
}

func (r *ReferralCode) CreateReferralCode(code *model.ReferralCode) error {
	return r.db.Create(code).Error
}

func (r *ReferralCode) GetReferralCodeByEmail(email string) (model.ReferralCode, error) {
	var code model.ReferralCode
	err := r.db.Table("users u").
		Select("r.id as id, r.code as code, r.is_active as is_active, r.expiration_time as expiration_time, r.user_id user_id").
		Joins("JOIN referral_codes r ON u.id = r.user_id").
		Where("u.username = ?", email).
		Scan(&code).Error
	return code, err
}

func (r *ReferralCode) GetReferralCode(code string) (model.ReferralCode, error) {
	var referralCode model.ReferralCode
	err := r.db.Where("code = ?", code).First(&referralCode).Error
	return referralCode, err
}

func (r *ReferralCode) GetReferralCodeByUserIdWithStatusActive(userID uint) (*model.ReferralCode, error) {
	var referralConde model.ReferralCode
	err := r.db.Where("user_id = ? AND is_active = ?", userID, true).Find(&referralConde).Error
	if err != nil {
		return nil, err
	}

	return &referralConde, nil
}

func (r *ReferralCode) DeleteReferralCode(codeID uint) error {
	var referrerID uint
	err := r.db.Table("referral_codes r").
		Select("r.user_id").
		Where("id = ?", codeID).
		Scan(&referrerID).Error
	if err != nil {
		return err
	}

	tx := r.db.Where("id = ?", codeID).Delete(&model.ReferralCode{})
	tx.Where("id = ?", referrerID).Delete(&model.ReferralCode{})

	return tx.Error
}
