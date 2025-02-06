package presentation

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Middleware функция возвращает middleware, который логирует информацию о запросах
func Middleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()

        c.Next()

        LoggerService.Info(fmt.Sprintf("Completed %s %s with %d in %v",
            c.Request.Method,
            c.Request.URL.Path,
            c.Writer.Status(),
            time.Since(start)))
    }
}
