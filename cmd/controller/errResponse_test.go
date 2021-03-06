package controller_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"dockerize/cmd/controller"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctl := controller.ErrResponse{}
	want := controller.NewError(http.StatusInternalServerError, errors.New("test Error"))

	got := ctl.Error(c)

	assert.Equal(t, want, got)
}
