package server

import (
	"os"

	"dockerize/cmd/controller"
	"dockerize/cmd/core"
	"dockerize/cmd/infrastructure"
	"github.com/gin-gonic/gin"
)

type handler func(ctx *gin.Context) error
type dependencies struct {
	postgres *infrastructure.Postgres
}

func resolver() dependencies {
	return dependencies{
		postgres: resolvePostgres(),
	}
}

/* CONTROLLERS */

func (d *dependencies) Ping() handler {
	ctl := controller.Ping{}

	return ctl.Ping
}

func (d *dependencies) PingDB() handler {
	ctl, err := controller.NewPingDB(d.pingDB())
	if err != nil {
		panic(err)
	}

	return ctl.PingDB
}

func (d *dependencies) Error() handler {
	ctl := controller.ErrResponse{}

	return ctl.Error
}

/* USE CASES. */
func (d *dependencies) pingDB() *core.PingDB {
	pingDB, err := core.NewPingDB(d.postgres)
	if err != nil {
		panic(err)
	}

	return pingDB
}
/* INFRASTRUCTURE. */
func resolvePostgres() *infrastructure.Postgres {
	url := os.Getenv("DATABASE_URL")

	postgres, err := infrastructure.NewPosgres(url)
	if err != nil {
		panic(err)
	}

	return postgres
}