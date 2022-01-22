package core

import "errors"

type pinger interface {
	Ping() error
}

// PingDB is the use case that find if a database is alive or not.
type PingDB struct {
	database pinger
}

// NewPingDB create a new instance of the use case ping DB.
func NewPingDB(pinger pinger) (*PingDB, error) {
	if pinger == nil {
		return nil, errors.New("pinger is nil")
	}

	return &PingDB{database: pinger}, nil
}

// Execute return if was possible ping to a database or not.
func (p *PingDB) Execute() error {
	return p.database.Ping()
}
