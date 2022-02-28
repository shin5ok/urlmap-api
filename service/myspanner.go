package service

import (
	"context"
	"errors"
	"io"

	database "cloud.google.com/go/spanner/admin/database/apiv1"
)

type spanner struct{}

type MyDB interface {
	Put(string) error
	Get(string) (*[]string, error)
	List() (*[]string, error)
}

func New() *spanner {
	return &spanner{}

}

func (s *spanner) Put() error {
	return errors.New("")
}

func (s *spanner) Get() (*[]string, error) {
	return &[]string{}, errors.New("")
}

func (s *spanner) List() (*[]string, error) {
	return &[]string{}, errors.New("")
}

func createClients(w io.Writer, db string) error {
	ctx := context.Background()

	adminClient, err := database.NewDatabaseAdminClient(ctx)
	if err != nil {
		return err
	}
	defer adminClient.Close()

	dataClient, err := spanner.NewClient(ctx, db)
	if err != nil {
		return err
	}
	defer dataClient.Close()

	_ = adminClient
	_ = dataClient

	return nil
}
