package handler

import (
	"github.com/fanfaronDo/referral_system_api/pkg/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(ctx *gin.Context) {
	var user model.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.AuthService.CreateUser(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"userID": user.ID})
}

func (h *Handler) signIn(ctx *gin.Context) {
	var user model.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.service.AuthService.GenerateToken(user.Username, user.Password)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": ErrUserNotRegistered.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"token": token})
}

func (h *Handler) signUpWithReferralCode(ctx *gin.Context) {
	var user model.User
	referrercode := ctx.Param(code)

	if referrercode == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": ErrReferralCodeIsRequired.Error()})
		return
	}

	if len(referrercode) != MaxLengthReferralCode {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": ErrInvalidReferrerCode.Error()})
		return
	}

	refcode, _ := h.service.ReferralCodeService.GetReferralCode(referrercode)

	err := h.service.ReferralCodeService.CheckReferralCode(refcode)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound,
			gin.H{"error": err.Error()})
		return
	}

	if err = ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = h.service.AuthService.CreateUser(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.service.ReferralService.CreateReferral(refcode, user.ID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"userID": user.ID})
}
