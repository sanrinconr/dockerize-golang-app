package controller_test

import (
	"net/http/httptest"
	"testing"

	"dockerize/cmd/controller"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctl := controller.Ping{}

	got := ctl.Ping(c)

	assert.Nil(t, got)
}
