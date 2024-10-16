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
	email, err := r.storage.ReferralStorage.GetEmailById(referralCode.UserId)
	if err != nil {
		return "", err
	}

	code := r.generateReferralCode(email)
	// create referral

	return code, nil
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
