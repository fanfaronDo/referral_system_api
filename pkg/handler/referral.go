package handler

import (
	"github.com/fanfaronDo/referral_system_api/pkg/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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

	code, err := h.service.ReferralService.CreateReferralCode(&referralCode)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": code})
}
