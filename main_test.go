package main

import (
	"context"
	"log"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	pb "github.com/shin5ok/urlmap-api/pb"
	service "github.com/shin5ok/urlmap-api/service"
)

var lis *bufconn.Listener

const bufSize = 1024 * 1024

// https://github.com/castaneai/grpc-testing-with-bufconn/blob/master/server/server_test.go
func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()

	server := service.Redirection{}
	pb.RegisterRedirectionServer(s, &server)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

}

func bufDialer(ctx context.Context, address string) (net.Conn, error) {
	return lis.Dial()
}

func TestSetUser(t *testing.T) {

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "localhost", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewRedirectionClient(conn)

	var notifyTo string = "test@example.com"
	var user string = "foo"
	resp, err := client.SetUser(ctx, &pb.User{User: user, NotifyTo: notifyTo})
	if err != nil {
		t.Error(err)
	}

	if resp.GetUser() != user {
		t.Errorf("not Match %s != %s", resp.GetUser(), user)
	}

}
