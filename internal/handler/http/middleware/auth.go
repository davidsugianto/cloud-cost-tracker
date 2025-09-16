package middleware

import (
	"net/http"
	"strings"

	"github.com/davidsugianto/cloud-cost-tracker/internal/pkg/auth"
	"github.com/davidsugianto/cloud-cost-tracker/internal/pkg/logger"
	"github.com/gin-gonic/gin"
)

func JWTAuth(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := extractBearerToken(c.GetHeader("Authorization"))
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid token"})
			return
		}

		claims, err := auth.ValidateToken(token, jwtSecret)
		if err != nil {
			logger.Warn(c.Request.Context(), "invalid token", map[string]interface{}{
				"error": err.Error(),
			})
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		c.Set("user_id", claims.UserID)
		c.Next()
	}
}

func extractBearerToken(authHeader string) string {
	if authHeader == "" {
		return ""
	}
	if strings.HasPrefix(authHeader, "Bearer ") {
		return strings.TrimPrefix(authHeader, "Bearer ")
	}
	return authHeader
}
