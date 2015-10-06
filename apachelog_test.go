package apachelog

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestApacheLog(t *testing.T) {
	g := gin.New()

	buf := &bytes.Buffer{}
	g.Use(NewWithWriter(buf))
	g.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	s := httptest.NewServer(g)
	defer s.Close()

	res, err := http.Get(s.URL)
	if !assert.NoError(t, err) {
		return
	}

	if !assert.Equal(t, res.StatusCode, http.StatusOK) {
		return
	}

	assert.NotEmpty(t, buf.String())
}