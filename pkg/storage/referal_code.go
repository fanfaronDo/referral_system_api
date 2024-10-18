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

func (r *ReferralCode) GetReferralCode(code string) (*model.ReferralCode, error) {
	var referralCode model.ReferralCode
	if err := r.db.Where("code=? AND is_active=?", code, true).First(&referralCode).Error; err != nil {
		return nil, err
	}
	if referralCode.Code == "" {
		return nil, ErrReferralCodeNotFound
	}

	return &referralCode, nil
}

func (r *ReferralCode) GetReferralCodeByUserIdWithStatusActive(userID uint) (*model.ReferralCode, error) {
	var referralConde model.ReferralCode
	err := r.db.Where("user_id = ? AND is_active = ?", userID, true).Find(&referralConde).Error
	if err != nil {
		return nil, ErrActiveReferralCodeNotFound
	}

	return &referralConde, nil
}

func (r *ReferralCode) UpdateReferralCodeStatus(referralCode *model.ReferralCode, status bool) error {
	if err := r.db.First(referralCode, referralCode.ID).Error; err != nil {
		return err
	}
	referralCode.IsActive = status
	if err := r.db.Save(referralCode).Error; err != nil {
		return err
	}
	return nil
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
