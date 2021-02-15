package main

import (
	"context"
	"flag"
	"log"
	pb "urlmap-api/pb"

	"google.golang.org/grpc"
)

func main() {
	host := flag.String("host", "host", "host you want to connect")
	flag.Parse()
	conn, err := grpc.Dial(*host, grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()
	client := pb.NewRedirectionClient(conn)
	orgurl := &pb.OrgUrl{Org: "https://example.com/"}

	if res, err := client.GetInfo(context.TODO(), orgurl); err != nil {
		log.Printf("error::%#v \n", err)
	} else {
		log.Printf("result:%#v \n", res)
	}
}
