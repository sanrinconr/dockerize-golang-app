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

type executorMock func() error

func (e executorMock) Execute() error {
	return e()
}

var defaultExecutorMock executorMock = func() error {
	return nil
}

func TestNew_fail(t *testing.T) {
	want := errors.New("nil executor")

	_, got := controller.NewPingDB(nil)

	assert.Equal(t, got, want)
}

func TestPingDB_success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	ctl, err := controller.NewPingDB(defaultExecutorMock)
	assert.Nil(t, err)
	err = ctl.PingDB(c)
	assert.Nil(t, err)

	t.Run("success_code", func(t *testing.T) {
		want := http.StatusOK

		got := w.Code

		assert.Equal(t, got, want)
	})

	t.Run("success_body", func(t *testing.T) {
		want := ""

		got := w.Body.String()

		assert.Equal(t, got, want)
	})
}

func TestPingDB_fail(t *testing.T) {
	t.Run("error_type", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		var executorMock executorMock = func() error {
			return errors.New("executor error")
		}
		ctl, err := controller.NewPingDB(executorMock)
		assert.Nil(t, err)
		want := controller.Error{
			Code:    http.StatusServiceUnavailable,
			Message: http.StatusText(http.StatusServiceUnavailable),
			Cause:   "executor error",
		}

		got := ctl.PingDB(c)

		assert.Equal(t, got, want)
	})
}