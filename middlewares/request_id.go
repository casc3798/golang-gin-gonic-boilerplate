package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//RequestIDMiddleware ...
//Generate a unique ID and attach it to each request for future reference or use
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := uuid.NewString()
		c.Writer.Header().Set("X-Request-Id", uuid)
		c.Next()
	}
}