package main

import (
	"context"
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
	"google.golang.org/grpc/codes"
	health "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

var port string = os.Getenv("PORT")
var version string = "2022011600"

type healthCheck struct{}

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

	serverLogger.Info().Msgf("Version of %s is Starting...\n", version)
	if port == "" {
		port = "8080"
	}
	listenPort, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		serverLogger.Fatal().Msg(err.Error())
	}

	service := &service.Redirection{}
	// service name is 'Redirection' that was defined in pb
	pb.RegisterRedirectionServer(server, service)

	var h = &healthCheck{}
	health.RegisterHealthServer(server, h)

	reflection.Register(server)
	serverLogger.Info().Msgf("Listening on %s\n", port)
	server.Serve(listenPort)
}

func (h *healthCheck) Check(context.Context, *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	return &health.HealthCheckResponse{
		Status: health.HealthCheckResponse_SERVING,
	}, nil
}

func (h *healthCheck) Watch(*health.HealthCheckRequest, health.Health_WatchServer) error {
	return status.Error(codes.Unimplemented, "No implementation for Watch")
}
