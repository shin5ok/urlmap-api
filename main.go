package main

import (
	"fmt"
	"net"
	"os"
	"time"
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
var version string = "2022011600"

func init() {
	log.Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	zerolog.LevelFieldName = "severity"
	zerolog.TimestampFieldName = "timestamp"
	zerolog.TimeFieldFormat = time.RFC3339Nano
}

func main() {
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

	log.Info().Msgf("Version of %s is Starting...\n", version)
	if port == "" {
		port = "8080"
	}
	listenPort, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	service := &service.Redirection{}
	// service name is 'Redirection' that was defined in pb
	pb.RegisterRedirectionServer(server, service)
	reflection.Register(server)
	fmt.Printf("Listening on %s\n", port)
	server.Serve(listenPort)
}
