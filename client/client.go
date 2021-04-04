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
	mode := flag.String("mode", "mode", "set or get")
	flag.Parse()
	conn, err := grpc.Dial(*host, grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()
	client := pb.NewRedirectionClient(conn)

	if *mode == "set" {
		data := &pb.RedirectData{
			&pb.RedirectInfo{"kawanos", "https://example.com/", "https://example.jp/", "my_comment", true},
			&pb.RedirectData_ValidDate{"2020-01-01", "2020-01-02"},
		}
		// data := &pb.RedirectData{}

		if res, err := client.SetInfo(context.TODO(), data); err != nil {
			log.Printf("error::%#v \n", err)
		} else {
			log.Printf("result:%#v \n", res)
		}
	} else {
		orgurl := &pb.OrgUrl{Org: "https://example.com/"}

		if res, err := client.GetInfo(context.TODO(), orgurl); err != nil {
			log.Printf("error::%#v \n", err)
		} else {
			log.Printf("result:%#v \n", res)
		}
	}
}
