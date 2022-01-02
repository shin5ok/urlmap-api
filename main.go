package main

import (
	"fmt"
	"net"
	"os"
	pb "urlmap-api/pb"
	"urlmap-api/service"

	_ "time/tzdata"

	"github.com/pereslava/grpc_zerolog"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var port string = os.Getenv("PORT")
var version string = "20210102"

func main() {
	log.Info().Msg(fmt.Sprintf("Version of %s is Starting...\n", version))
	if port == "" {
		port = "8080"
	}
	listenPort, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	// logger for gRPC to zerolog
	// https://pkg.go.dev/github.com/pereslava/grpc_zerolog#section-readme
	serverLogger := log.Level(zerolog.TraceLevel)
	grpc_zerolog.ReplaceGrpcLogger(zerolog.New(os.Stderr).Level(zerolog.ErrorLevel))

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc_zerolog.NewPayloadUnaryServerInterceptor(serverLogger),
			grpc_zerolog.NewPayloadUnaryServerInterceptor(serverLogger),
		),
		grpc.ChainStreamInterceptor(
			grpc_zerolog.NewPayloadStreamServerInterceptor(serverLogger),
			grpc_zerolog.NewStreamServerInterceptor(serverLogger),
		),
	)

	service := &service.Redirection{}
	// service name is 'Redirection' that was defined in pb
	pb.RegisterRedirectionServer(server, service)
	reflection.Register(server)
	fmt.Printf("Listening on %s\n", port)
	server.Serve(listenPort)
}
