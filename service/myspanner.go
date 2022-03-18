package service

import (
	"context"
	"errors"
	"fmt"

	"cloud.google.com/go/spanner"
)

type spannerConfig struct {
	db     string
	client *spanner.Client
}

type MyDB interface {
	Put(string) error
	Get(string) (*[]string, error)
	List() (*[]string, error)
}

func New(db string) *spannerConfig {
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

func (s *spannerConfig) Put(*[]string) error {
	return errors.New("")
}

func (s *spannerConfig) Get() (*[]string, error) {
	return &[]string{}, errors.New("")
}

func (s *spannerConfig) List() (*[]string, error) {
	return &[]string{}, errors.New("")
}
