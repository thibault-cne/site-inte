package middlewares

import (
	"backend/routes/auth"
	users_services "backend/services/users.services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		gu, err := auth.GetUser(ctx)
		if err != nil {
			ctx.Set("Logged", false)
			return
		}

		user, err := users_services.GetUser(gu.UserID)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error getting user"})
		}

		ctx.Set("Logged", true)
		ctx.Set("UserID", gu.UserID)
		ctx.Set("User", user)
	}
}
