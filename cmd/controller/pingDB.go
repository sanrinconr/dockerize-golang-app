package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type executor interface {
	Execute() error
}

// PingDB controller to validate health of the app.
type PingDB struct {
	pinger executor
}

// NewPingDB create a controller.
func NewPingDB(executor executor) (*PingDB, error) {
	if executor == nil {
		return nil, errors.New("nil executor")
	}

	return &PingDB{pinger: executor}, nil
}

// PingDB validate the health of the database.
func (p *PingDB) PingDB(c *gin.Context) error {
	if err := p.pinger.Execute(); err != nil {
		return NewError(http.StatusServiceUnavailable, err)
	}

	c.Status(http.StatusOK)

	return nil
}
