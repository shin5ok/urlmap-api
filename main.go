package main

import (
	"fmt"
	"log"
	"net"
	"os"
	pb "urlmap-api/pb"
	"urlmap-api/service"

	_ "time/tzdata"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
)

var port string = os.Getenv("PORT")
var version string = "1.06"

func main() {
	if port == "" {
		port = "8080"
	}
	listenPort, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalln(err)
	}

	zap, _ := zap.NewProduction()
	zap_opt := grpc_zap.WithLevels(
		func(c codes.Code) zapcore.Level {
			return zapcore.InfoLevel
		},
	)
	server := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(zap, zap_opt),
		),
	)

	service := &service.Redirection{}
	// service name is 'Redirection' that was defined in pb
	pb.RegisterRedirectionServer(server, service)
	reflection.Register(server)
	fmt.Printf("Listening on %s\n", port)
	server.Serve(listenPort)
}
