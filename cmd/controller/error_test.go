package controller_test

import (
	"errors"
	"net/http"
	"testing"

	"dockerize/cmd/controller"
	"github.com/stretchr/testify/assert"
)

func TestNewError(t *testing.T) {
	t.Parallel()

	t.Run("fe", func(t *testing.T) {
		t.Parallel()
		err := controller.NewError(http.StatusOK, errors.New("pruebas"))
		want := "pruebas"

		got := err.Error()

		assert.Equal(t, want, got)
	})
}
