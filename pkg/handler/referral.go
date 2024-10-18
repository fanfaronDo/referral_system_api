package handler

import (
	"fmt"
	"github.com/fanfaronDo/referral_system_api/pkg/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const (
	code                  = "code"
	email                 = "email"
	MaxLengthReferralCode = 8
)

type referralCodeController struct {
	ExpirationTime string `json:"expiration_time"`
}

func (h *Handler) createReferralCode(ctx *gin.Context) {
	var referralCode model.ReferralCode
	var refCodeController referralCodeController

	referralCode.UserId = getUserId(ctx)
	if err := ctx.ShouldBindJSON(&refCodeController); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	duration, err := time.ParseDuration(refCodeController.ExpirationTime)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": err.Error() + "with expiration_time: " + refCodeController.ExpirationTime})
		return
	}

	referralCode.ExpirationTime = duration
	fmt.Println(referralCode)
	codecreate, err := h.service.ReferralCodeService.CreateReferralCode(&referralCode)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": codecreate})
}

func (h *Handler) deleteReferralCode(ctx *gin.Context) {
	refcode := ctx.Param(code)
	if refcode == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": ErrReferralCodeIsRequired.Error()})
		return
	}
	if len(refcode) != MaxLengthReferralCode {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": ErrInvalidReferrerCode.Error()})
		return
	}

	userid := getUserId(ctx)

	err := h.service.ReferralCodeService.DeleteReferralCode(userid, refcode)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.AbortWithStatus(http.StatusOK)
}

func (h *Handler) getReferralCodeByEmail(ctx *gin.Context) {
	emailuser := ctx.Query(email)
	if emailuser == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": ErrEmailRequired.Error()})
		return
	}

	userid := getUserId(ctx)
	referrer, err := h.service.ReferralCodeService.GetReferralCodeByEmail(userid, emailuser)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": referrer.Code})
}
