package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	pb "urlmap-api/pb"

	"google.golang.org/grpc"
)

func main() {
	host := flag.String("host", "localhost:8080", "host you want to connect")
	mode := flag.String("mode", "get", "set or get")
	flag.Parse()
	conn, err := grpc.Dial(*host, grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()
	client := pb.NewRedirectionClient(conn)

	if *mode == "set" {
		data := &pb.RedirectData{}
		data.Redirect = &pb.RedirectInfo{
			User:         "kawanos",
			Org:          "https://example.com/",
			RedirectPath: "takosuke",
			Comment:      "my_comment",
			Active:       true}
		// &pb.RedirectData_ValidDate{"2020-01-01", "2020-01-02"},

		if res, err := client.SetInfo(context.TODO(), data); err != nil {
			log.Printf("error::%#v \n", err)
		} else {
			log.Printf("result:%#v \n", res)
		}
	} else {
		orgurl := &pb.OrgUrl{Org: "https://example.com/"}

		if res, err := client.GetInfo(context.TODO(), orgurl); err != nil {
			log.Printf("error:%#v \n", err)
		} else {
			fmt.Printf("Redirect.User:%s, Redirect.Org:%s, Redirect.Redirect: %s\n", res.Redirect.User, res.Redirect.Org, res.Redirect.RedirectPath)
		}
	}
}
