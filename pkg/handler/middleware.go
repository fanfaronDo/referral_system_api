package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const userctx = "userID"

func (h *Handler) userIdentity(ctx *gin.Context) {
	header := ctx.GetHeader("Authorization")

	if header == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": ErrHeaderAuthUndefined.Error() + "1"})
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": ErrHeaderAuthUndefined.Error() + "2"})
		return
	}
	if len(headerParts[1]) == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": ErrHeaderAuthUndefined.Error() + "3"})
		return
	}

	userid, err := h.service.AuthService.ParseToken(headerParts[1])
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": ErrInvalidToken.Error()})
		return
	}

	ctx.Set(userctx, userid)
}

func getUserId(ctx *gin.Context) uint {
	id, exists := ctx.Get(userctx)
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": ErrHeaderAuthUndefined.Error()})
		return 0
	}

	return id.(uint)
}
