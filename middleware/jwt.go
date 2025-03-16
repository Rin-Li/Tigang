package middleware

import (
	"Tigang/pkg/util"
	"time"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(401, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				c.JSON(401, gin.H{
					"message": "Unauthorized",
				})
				c.Abort()
				return
			} else if time.Now().Unix() > claims.ExpiresAt.Time.Unix() {
				c.JSON(401, gin.H{
					"message": "Pass time",
				})
				c.Abort()
				return
			}
		}
		// Valid , continue
		c.Next()
	}
}
