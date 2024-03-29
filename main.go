package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	pb "github.com/shin5ok/urlmap-api/pb"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"

	"github.com/shin5ok/urlmap-api/service"

	_ "time/tzdata"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/pereslava/grpc_zerolog"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/shin5ok/shoutouthostnamegcp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	health "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

var (
	port      string = os.Getenv("PORT")
	version   string = "2022061000"
	appPort   string = "8080"
	promPort  string = "18080"
	projectID string
	dbParams  service.DbParams
)

type healthCheck struct{}

func init() {
	log.Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	zerolog.LevelFieldName = "severity"
	zerolog.TimestampFieldName = "timestamp"
	zerolog.TimeFieldFormat = time.RFC3339Nano

	shoutouthostnamegcp.SetSigHandler(os.Getenv("SLACK_URL"), os.Getenv("SLACK_CHANNEL"))

	dbParams.Dbuser = os.Getenv("DBUSER")
	dbParams.Dbpass = os.Getenv("DBPASS")
	dbParams.Dbname = os.Getenv("DBNAME")
	dbParams.Dbhost = os.Getenv("DBHOST")

	projectID = os.Getenv("GOOGLE_CLOUD_PROJECT")
}

func main() {
	// logger for gRPC to zerolog
	// https://pkg.go.dev/github.com/pereslava/grpc_zerolog#section-readme
	serverLogger := log.Level(zerolog.TraceLevel)
	grpc_zerolog.ReplaceGrpcLogger(zerolog.New(os.Stderr).Level(zerolog.ErrorLevel))

	tp := tpExporter(projectID, "urlmap-api")
	ctx := context.Background()
	defer tp.ForceFlush(ctx)
	otel.SetTracerProvider(tp)

	interceptorOpt := otelgrpc.WithTracerProvider(otel.GetTracerProvider())

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc_zerolog.NewPayloadUnaryServerInterceptor(serverLogger),
			grpc_prometheus.UnaryServerInterceptor,
			otelgrpc.UnaryServerInterceptor(interceptorOpt),
		),
		grpc.ChainStreamInterceptor(
			grpc_zerolog.NewPayloadStreamServerInterceptor(serverLogger),
			grpc_zerolog.NewStreamServerInterceptor(serverLogger),
			grpc_prometheus.StreamServerInterceptor,
			otelgrpc.StreamServerInterceptor(interceptorOpt),
		),
	)

	serverLogger.Info().Msgf("Version of %s is Starting...\n", version)
	if port == "" {
		port = appPort
	}
	listenPort, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		serverLogger.Fatal().Msg(err.Error())
	}

	service := service.New(dbParams)
	// service name is 'Redirection' that was defined in pb
	pb.RegisterRedirectionServer(server, &service)

	var h = &healthCheck{}
	health.RegisterHealthServer(server, h)

	grpc_prometheus.EnableHandlingTimeHistogram()
	grpc_prometheus.Register(server)
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		if err := http.ListenAndServe(":"+promPort, nil); err != nil {
			panic(err)
		}
		serverLogger.Info().Msgf("Prometheus exporter listening on :%s\n", promPort)
	}()

	reflection.Register(server)
	serverLogger.Info().Msgf("gGRPC server listening on :%s\n", port)
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
