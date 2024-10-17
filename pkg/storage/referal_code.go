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

func (r *ReferralCode) GetReferralCodeByEmail(userID uint, email string) (model.ReferralCode, error) {
	var code model.ReferralCode
	err := r.db.Table("users u").
		Select("r.id as id, r.code as code, r.is_active as is_active, r.expiration_time as expiration_time, r.user_id user_id").
		Joins("JOIN referral_codes r ON u.id = r.user_id").
		Where("u.id = ? AND u.username = ? AND r.is_active = ?", userID, email, true).
		Scan(&code).Error

	if code.UserId == 0 {
		return model.ReferralCode{}, ErrActiveReferralCodeNotFound
	}

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
		return nil, ErrActiveReferralCodeNotFound
	}

	return &referralConde, nil
}

func (r *ReferralCode) DeleteReferralCode(userID uint, code string) error {
	var referrerID uint
	err := r.db.Table("referral_codes r").
		Select("r.user_id").
		Where("code = ? AND user_id = ?", code, userID).
		Scan(&referrerID).Error

	if err != nil || referrerID == 0 {
		return ErrUserCodeNotFound
	}

	tx := r.db.Where("code = ?", code).Delete(&model.ReferralCode{})

	return tx.Error
}
