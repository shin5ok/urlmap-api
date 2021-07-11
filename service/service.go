package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	pb "urlmap-api/pb"
)

type Redirection struct{}

func (s *Redirection) GetInfo(ctx context.Context, org *pb.OrgUrl) (*pb.RedirectData, error) {
	redirectdata := &pb.RedirectData{}
	redirectdata.Redirect = &pb.RedirectInfo{User: "kawanos", RedirectPath: "redirectingExamplePath"}
	redirectdata.Redirect.Org = "https://example.com/foobar"
	fmt.Println(redirectdata)
	if true {
		return redirectdata, nil
	}
	return nil, errors.New("Error")
}

func (s *Redirection) SetInfo(ctx context.Context, r *pb.RedirectData) (*pb.OrgUrl, error) {
	// just stub for a test
	db, err := sqlConnect()
	if err != nil {
		// return &pb.OrgUrl{}, nil
		log.Fatal(err)
	}
	redirect := Redirects{}
	redirect.RedirectPath = r.Redirect.RedirectPath
	redirect.User = r.Redirect.User
	redirect.Org = r.Redirect.Org
	redirect.Active = 1
	result := db.Create(&redirect)
	if result.Error != nil {
		return &pb.OrgUrl{}, result.Error
	}
	org := &pb.OrgUrl{Org: r.Redirect.Org}
	return org, nil
}
