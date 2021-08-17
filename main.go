package main

import (
	"fmt"
	"log"
	"net"
	"os"
	pb "urlmap-api/pb"
	"urlmap-api/service"

	_ "time/tzdata"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var port string = os.Getenv("PORT")
var version string = "0.61"

func main() {
	if port == "" {
		port = "8080"
	}
	listenPort, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()

	service := &service.Redirection{}
	// service name is 'Redirection' that was defined in pb
	pb.RegisterRedirectionServer(server, service)
	reflection.Register(server)
	fmt.Printf("Listening on %s\n", port)
	server.Serve(listenPort)
}
