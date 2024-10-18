package service

import (
	"github.com/fanfaronDo/referral_system_api/pkg/model"
)

type ReferralCodeActiveValidation struct {
	referralCodeActive *model.ReferralCode
}

func NewReferralCodeActiveValidation(referralCodeActive *model.ReferralCode) *ReferralCodeActiveValidation {
	return &ReferralCodeActiveValidation{referralCodeActive}
}

func (v *ReferralCodeActiveValidation) IsExists() bool {
	return v.referralCodeActive != nil
}

func (v *ReferralCodeActiveValidation) IsTimeAlive() bool {
	return v.referralCodeActive.IsActive
}
