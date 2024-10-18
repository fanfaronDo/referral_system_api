package handler

import (
	"github.com/fanfaronDo/referral_system_api/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	route := gin.New()

	route.GET("/ping", func(ginCtx *gin.Context) {
		ginCtx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	auth := route.Group("/auth")
	{
		auth.POST("/signup", h.signUp)
		auth.POST("/signin", h.signIn)
		auth.POST("/signup/:code", h.signUpWithReferralCode)
	}

	api := route.Group("/api", h.userIdentity)
	{
		api.POST("/referral-code", h.createReferralCode)
		api.GET("/referral-code", h.getReferralCodeByEmail)
		api.DELETE("/referral-code/:code", h.deleteReferralCode)
		api.GET("/referral-code/:id", h.getReferralsByIdReferrer)
	}

	return route
}
