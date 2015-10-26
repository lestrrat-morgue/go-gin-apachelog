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
	return NewWithWriter(gin.DefaultWriter)
}

func NewWithWriter(out io.Writer) gin.HandlerFunc {
	return NewWithWriterAndLogger(out, al.CombinedLog.Clone())
}

func NewWithWriterAndLogger(out io.Writer, logger *al.ApacheLog) gin.HandlerFunc {
	logger.SetOutput(out)
	return func(c *gin.Context) {
		t := time.Now()

		// URL Paths *may* be modified, so keep it for later use
		path := c.Request.URL.Path

		// Process request
		c.Next()

		c.Request.URL.Path = path
		logger.LogLine(c.Request, c.Writer.Status(), c.Writer.Header(), time.Since(t))
	}
}