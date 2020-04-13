package store

import (
	"database/sql"
	"github.com/remicaumette/zaap.sh/zaap-services/pkg/core"
)

type ApplicationStore struct {
	conn *sql.DB
}

func NewApplicationStore(conn *sql.DB) *ApplicationStore {
	return &ApplicationStore{conn}
}

func (s *ApplicationStore) FindApplicationById() (*core.Application, error) {
	return nil, nil
}
