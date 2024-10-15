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
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": ErrHeaderAuthUndefined.Error()})
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": ErrHeaderAuthUndefined.Error()})
		return
	}
	if len(headerParts[1]) == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": ErrHeaderAuthUndefined.Error()})
		return
	}

	userid, err := h.service.ServiceAuth.ParseToken(headerParts[1])
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
