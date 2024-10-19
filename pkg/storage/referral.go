package storage

import (
	"fmt"
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

func (r *Referral) GetReferrersById(referrerId uint) ([]model.ReferralInfo, error) {

	var referralInfo []model.ReferralInfo

	err := r.db.
		Table("referrals r").
		Select("u.username as username, r.created_at as created_at").
		Joins("JOIN users u ON r.referred_id = u.id").
		Where("r.referrer_id = ?", referrerId).
		Find(&referralInfo).Error

	if err != nil {
		return nil, err
	}

	fmt.Println(referralInfo)

	return referralInfo, nil
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
