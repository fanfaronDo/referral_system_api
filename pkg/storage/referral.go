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

func (r *Referral) GetReferrersById(referrerId uint) ([]model.ReferralCode, error) {

	var referredUsers []model.Referral
	if err := r.db.Where("referrer_id = ?", referrerId).Find(&referredUsers).Error; err != nil {
		return nil, err
	}

	userIds := make([]uint, 0, len(referredUsers))
	for _, referral := range referredUsers {
		userIds = append(userIds, referral.ReferredId)
	}

	var referralCodes []model.ReferralCode
	if err := r.db.Where("user_id IN ?", userIds).Find(&referralCodes).Error; err != nil {
		return nil, err
	}

	return referralCodes, nil
}

func (r *Referral) GetReferralCodeByEmail(userID uint, email string) (model.ReferralCode, error) {
	var code model.ReferralCode
	err := r.db.Table("users u").
		Select("r.id as id, r.code as code, r.is_active as is_active, r.expiration_time as expiration_time, r.user_id user_id").
		Joins("JOIN referrals r ON u.id = r.referrer_id").
		Where("r.referred_id = ? AND u.username = ?", userID, email).
		Scan(&code).Error

	if code.UserId == 0 {
		return model.ReferralCode{}, ErrActiveReferralCodeNotFound
	}

	return code, err
}

func (r *Referral) GetEmailById(userId uint) (string, error) {
	var email string
	var user model.User
	err := r.db.Table("users").Where("id = ?", userId).First(&user).Error
	if err != nil {
		return "", err
	}
	email = user.Username
	return email, err
}
