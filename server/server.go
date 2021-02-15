package main

import (
	"fmt"
	"log"
	"net"
	"os"
	pb "urlmap-api/pb"
	"urlmap-api/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var port string = fmt.Sprintf(":%s", os.Getenv("PORT"))

func main() {
	listenPort, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()

	service := &service.Redirection{}
	// service name is 'Redirection' that was defined in pb
	pb.RegisterRedirectionServer(server, service)
	reflection.Register(server)
	server.Serve(listenPort)
}
