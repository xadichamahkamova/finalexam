package authorization

import (
	"log"
	"net/http"
	"strings"

	t "api-gateway/internal/https/token"
	"api-gateway/logger"
	"api-gateway/internal/cache" 

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	Cache *cache.RedisMethod 
}

func (a *AuthMiddleware) MiddleWare() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		logger.Info("Authorization header:", token)
		url := ctx.Request.URL.Path
		logger.Info("Request URL:", url)

		if strings.Contains(url, "swagger") || url == "/api/v1/users" || url == "/api/v1/users/login" {
			ctx.Next()
			return
		}

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Authorization header is missing",
			})
			return
		}

		if !strings.HasPrefix(token, "Bearer ") {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Authorization token is missing Bearer prefix",
			})
			return
		}

		token = strings.TrimPrefix(token, "Bearer ")

		claims, err := t.ExtractClaim(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": err.Error(),
			})
			return
		}
		log.Println(claims)

		user_id, ok := claims["user_id"].(string)
		if !ok || user_id == "" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "ID not found in token",
			})
			return
		}

		user_email, ok := claims["user_email"].(string)
		if !ok || user_id == "" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "ID not found in token",
			})
			return
		}

	
		err = a.Cache.HoldOnUserID(user_id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to save user ID to Redis",
			})
			return
		}

		logger.Info("USER ID:", user_id, "USER EMAIL: ", user_email)
		
		ctx.Set("user_id", user_id)
		ctx.Set("user_email", user_email)
		
		ctx.Next()
	}
}
