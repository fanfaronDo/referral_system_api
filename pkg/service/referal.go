package service

import (
	"github.com/fanfaronDo/referral_system_api/pkg/model"
	"github.com/fanfaronDo/referral_system_api/pkg/storage"
)

type Referral struct {
	storage *storage.Storage
}

func NewReferral(storage *storage.Storage) *Referral {
	return &Referral{storage}
}

func (r *Referral) CreateReferral(referralCode *model.ReferralCode, userID uint) error {
	var referral model.Referral
	referral.ReferrerId = referralCode.UserId
	referral.ReferredId = userID

	if err := r.storage.ReferralCodeStorage.UpdateReferralCodeStatus(referralCode, false); err != nil {
		return err
	}

	return r.storage.ReferralStorage.CreateReferral(&referral)
}

func (r *Referral) GetReferrersById(referrerId uint) ([]model.ReferralCode, error) {
	return r.storage.GetReferrersById(referrerId)
}
