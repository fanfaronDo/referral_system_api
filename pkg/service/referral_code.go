package service

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/fanfaronDo/referral_system_api/pkg/model"
	"github.com/fanfaronDo/referral_system_api/pkg/storage"
	"time"
)

type ReferralCode struct {
	storage *storage.Storage
}

func NewReferralCode(storage *storage.Storage) *ReferralCode {
	return &ReferralCode{storage}
}

func (r *ReferralCode) CreateReferralCode(referralCode *model.ReferralCode) (string, error) {
	var codeActive string

	referralCodeActive, _ := r.storage.ReferralCodeStorage.
		GetReferralCodeByUserIdWithStatusActive(referralCode.UserId)

	referralCodeValidator := NewReferralCodeActiveValidation(referralCodeActive)

	if referralCodeValidator.IsExists() {
		referralCodeActive.UpdateAliveTimeStatus()

		if referralCodeValidator.IsTimeAlive() {
			return referralCodeValidator.referralCodeActive.Code, nil
		}
		if referralCodeActive.ID != 0 {
			err := r.storage.ReferralCodeStorage.UpdateReferralCodeStatus(referralCodeActive, referralCodeActive.IsActive)
			if err != nil {
				return "", err
			}
		}
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

func (r *ReferralCode) GetReferralCodeByEmail(userID uint, email string) (model.ReferralCode, error) {
	return r.storage.GetReferralCodeByEmail(userID, email)
}

func (r *ReferralCode) GetReferralCode(code string) (*model.ReferralCode, error) {
	return r.storage.GetReferralCode(code)
}

func (r *ReferralCode) CheckReferralCode(referralCode *model.ReferralCode) error {
	referralCodeValid := NewReferralCodeActiveValidation(referralCode)
	if !referralCodeValid.IsExists() {
		return ErrReferrerCodeNotFound
	}

	referralCode.UpdateAliveTimeStatus()
	if !referralCode.IsActive {
		return ErrReferrerCodeIsOutOfDate
	}

	return nil
}

func (r *ReferralCode) UpdateReferralCodeStatus(referralCode *model.ReferralCode, status bool) error {
	return r.storage.ReferralCodeStorage.UpdateReferralCodeStatus(referralCode, status)
}

func (r *ReferralCode) DeleteReferralCode(userID uint, code string) error {
	return r.storage.ReferralCodeStorage.DeleteReferralCode(userID, code)
}

func (r *ReferralCode) generateReferralCode(email string) string {
	hash := sha256.New()
	_, err := hash.Write([]byte(email + time.Now().String()))
	if err != nil {
		fmt.Println("Error writing to hash:", err)
		return ""
	}

	hashBytes := hash.Sum(nil)
	referralCode := hex.EncodeToString(hashBytes)

	return referralCode[:MaxLengthReferralCode]
}
