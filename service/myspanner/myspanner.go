package myspanner

import (
	"context"
	"errors"
	"fmt"

	"cloud.google.com/go/spanner"
	"github.com/rs/zerolog/log"
	"github.com/shin5ok/urlmap-api/service"
	"google.golang.org/api/iterator"
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
	defer client.Close()
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
	ctx := context.Background()
	_, err := s.client.Apply(ctx, m)
	if err != nil {
		return err
	}
	return nil
}

func (s *spannerConfig) Get(query string) ([]string, error) {

	stmt := spanner.Statement{
		SQL: `SELECT user,redirect_path,org from redirects
					WHERE redirect_path = @redirectPath`,
		Params: map[string]interface{}{
			"redirectPath": query,
		},
	}
	var results []string
	ctx := context.Background()
	iter := s.client.Single().Query(ctx, stmt)
	defer iter.Stop()
	for {
		row, err := iter.Next()
		if err == iterator.Done {
			log.Info().Msgf("%+v", results)
			return results, nil
		}
		if err != nil {
			return results, err
		}
		var user, redirectPath, org string
		if err := row.Columns(&user, &redirectPath, &org); err != nil {
			return results, err
		}
		results = append(results, redirectPath)
	}
}

func (s *spannerConfig) List() ([]string, error) {
	return []string{}, errors.New("")
}
