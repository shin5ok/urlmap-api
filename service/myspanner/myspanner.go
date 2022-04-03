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
	ctx    context.Context
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
	config.ctx = ctx
	return config
}

func (s *spannerConfig) Put(params []interface{}) error {
	m := []*spanner.Mutation{
		/*
			spanner.InsertOrUpdate("Singers", singerColumns, []interface{}{1, "Marc", "Richards"}),
			spanner.InsertOrUpdate("Singers", singerColumns, []interface{}{2, "Catalina", "Smith"}),
		*/
		spanner.InsertOrUpdate(service.RedirectTableName, service.RedirectTableColumn, params),
	}
	_, err := s.client.Apply(s.ctx, m)
	if err != nil {
		return err
	}
	return nil
}

func (s *spannerConfig) Get(query string) ([]string, error) {
	return []string{}, errors.New("")
}

func (s *spannerConfig) List() ([]string, error) {
	return []string{}, errors.New("")
}
