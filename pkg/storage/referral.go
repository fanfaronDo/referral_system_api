package storage

import (
	"github.com/fanfaronDo/referral_system_api/pkg/model"
	"gorm.io/gorm"
)

type Referral struct {
	db *gorm.DB
}

func NewReferral(db *gorm.DB) *Referral {
	return &Referral{
		db: db,
	}
}

func (r *Referral) CreateReferral(referral *model.Referral) error {
	return r.db.Create(referral).Error
}

func (r *Referral) CreateReferralCode(code *model.ReferralCode) error {
	return r.db.Create(code).Error
}

func (r *Referral) GetReferralCodeByEmail(email string) (model.ReferralCode, error) {
	var code model.ReferralCode
	err := r.db.Table("users u").
		Select("r.id as id, r.code as code, r.is_active as is_active, r.expiration_time as expiration_time, r.user_id user_id").
		Joins("JOIN referral_codes r ON u.id = r.user_id").
		Where("u.username = ?", email).
		Scan(&code).Error
	return code, err
}

func (r *Referral) GetReferrersById(referrerId int) ([]model.ReferralCode, error) {

	var referredUsers []model.Referral
	if err := r.db.Where("referrer_id = ?", referrerId).Find(&referredUsers).Error; err != nil {
		return nil, err
	}

	userIds := make([]int, 0, len(referredUsers))
	for _, referral := range referredUsers {
		userIds = append(userIds, referral.ReferredId)
	}

	var referralCodes []model.ReferralCode
	if err := r.db.Where("user_id IN ?", userIds).Find(&referralCodes).Error; err != nil {
		return nil, err
	}

	return referralCodes, nil
}

func (r *Referral) GetReferralCode(code string) (model.ReferralCode, error) {
	var referralCode model.ReferralCode
	err := r.db.Where("code = ?", code).First(&referralCode).Error
	return referralCode, err
}

func (r *Referral) GetEmailById(userId int) (string, error) {
	var email string
	err := r.db.Table("users").Where("id = ?", userId).First(&email).Error
	return email, err
}

func (r *Referral) DeleteReferralCode(codeID uint) error {
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
