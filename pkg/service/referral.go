package service

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/fanfaronDo/referral_system_api/pkg/model"
	"github.com/fanfaronDo/referral_system_api/pkg/storage"
	"time"
)

type Referral struct {
	storage *storage.Storage
}

func NewReferral(storage *storage.Storage) *Referral {
	return &Referral{storage}
}

func (r *Referral) CreateReferralCode(referralCode *model.ReferralCode) (string, error) {
	var codeActive string

	referralCodeActive, err := r.storage.ReferralCodeStorage.
		GetReferralCodeByUserIdWithStatusActive(referralCode.UserId)

	if err != nil {
		return "", err
	}

	referralCodeValidator := NewReferralCodeActiveValidation(referralCodeActive)

	if referralCodeValidator.IsExists() && referralCodeValidator.IsTimeAlive() {
		return referralCodeValidator.referralCodeActive.Code, nil
	}

	email, err := r.storage.ReferralStorage.GetEmailById(referralCode.UserId)
	if err != nil {
		return "", ErrUserNotFound
	}
	codeActive = r.generateReferralCode(email)
	referralCode.IsActive = true
	referralCode.Code = codeActive
	if err = r.storage.ReferralCodeStorage.CreateReferralCode(referralCode); err != nil {
		return "", err
	}

	return codeActive, nil
}

func (r *Referral) DeleteReferralCode(userID uint, code string) error {
	return r.storage.ReferralCodeStorage.DeleteReferralCode(userID, code)
}

func (r *Referral) GetReferralCodeByEmail(userID uint, email string) (model.ReferralCode, error) {
	referrerCode, err := r.storage.GetReferralCodeByEmail(userID, email)
	if err != nil {
		return model.ReferralCode{}, err
	}
	return referrerCode, nil
}

func (r *Referral) generateReferralCode(email string) string {
	hash := sha256.New()
	_, err := hash.Write([]byte(email + time.Now().String()))
	if err != nil {
		fmt.Println("Error writing to hash:", err)
		return ""
	}

	hashBytes := hash.Sum(nil)
	referralCode := hex.EncodeToString(hashBytes)

	return referralCode[:8]
}
