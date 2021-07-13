package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"strings"
	pb "urlmap-api/pb"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

func main() {
	host := flag.String("host", "localhost:8080", "host you want to connect")
	path := flag.String("path", "", "redirect path")
	mode := flag.String("mode", "get", "set or get")
	user := flag.String("user", "", "info data")
	flag.Parse()
	conn, err := grpc.Dial(*host, grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()
	client := pb.NewRedirectionClient(conn)

	if *mode == "set" {
		data := &pb.RedirectData{}
		randPath := strings.Split(uuid.New().String(), "-")[0]
		data.Redirect = &pb.RedirectInfo{
			User:         "kawanos",
			Org:          "https://example.com/" + randPath,
			RedirectPath: randPath,
			Comment:      "my_comment",
			Active:       1}
		// &pb.RedirectData_ValidDate{"2020-01-01", "2020-01-02"},

		if res, err := client.SetInfo(context.TODO(), data); err != nil {
			log.Printf("error::%#v \n", err)
		} else {
			log.Printf(randPath)
			log.Printf("result:%#v \n", res)
		}
	} else if *mode == "info" {
		u := &pb.User{User: *user}

		if res, err := client.GetInfoByUser(context.TODO(), u); err != nil {
			log.Printf("error:%#v \n", err)
		} else {
			// format specified "%+v" to dump
			// fmt.Printf("%+v\n", res)
			j, _ := json.Marshal(res)
			fmt.Println(string(j))
		}

	} else {
		path := &pb.RedirectPath{Path: *path}

		if res, err := client.GetOrgByPath(context.TODO(), path); err != nil {
			log.Printf("error:%#v \n", err)
		} else {
			// fmt.Printf("Redirect.User:%s, Redirect.Org:%s, Redirect.Redirect: %s\n", res.Redirect.User, res.Redirect.Org, res.Redirect.RedirectPath)
			fmt.Println(res)
		}
	}
}
