package apachelog

import (
	"io"
	"time"

	"github.com/gin-gonic/gin"
	al "github.com/lestrrat/go-apache-logformat"
)

// New creates a new middleware that writes to `gin.DefaultWriter` and
// uses go-apache-logformat's CombinedLog format
func New() gin.HandlerFunc {
	return NewWithWriterAndLogger(gin.DefaultWriter, al.CombinedLog.Clone())
}

func NewWithWriterAndLogger(out io.Writer, logger *al.ApacheLog) gin.HandlerFunc {
	logger.SetOutput(out)
	return func(c *gin.Context) {
		t := time.Now()
		// Process request
		c.Next()

		logger.LogLine(c.Request, c.Writer.Status(), c.Writer.Header(), time.Since(t))
	}
}