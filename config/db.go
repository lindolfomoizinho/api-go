package config

import (
	"sync"

	"github.com/lindolfomoizinho/api-go/schemas"
)

type Datastore struct {
	M map[string]schemas.User
	*sync.RWMutex
}

func NewDatastore() *Datastore {
	return &Datastore{
		M:       make(map[string]schemas.User),
		RWMutex: &sync.RWMutex{},
	}
}
