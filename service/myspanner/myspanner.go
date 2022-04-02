package myspanner

import (
	"context"
	"errors"
	"fmt"

	"cloud.google.com/go/spanner"
	"github.com/shin5ok/urlmap-api/service"
)

type myDB service.MyDB

type spannerConfig struct {
	db     string
	client *spanner.Client
}

func New(db string) myDB {
	config := &spannerConfig{}
	config.db = db
	ctx := context.Background()
	client, err := spanner.NewClient(ctx, db)
	if err != nil {
		panic(fmt.Sprintf("cannot create a spanner client: %v", err))
	}
	config.client = client
	return config
}

func (s *spannerConfig) Put(params *[]string) error {
	return errors.New("")
}

func (s *spannerConfig) Get(query string) (*[]string, error) {
	return &[]string{}, errors.New("")
}

func (s *spannerConfig) List() (*[]string, error) {
	return &[]string{}, errors.New("")
}
