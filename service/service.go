package service

import (
	"context"
	"errors"
	"fmt"
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

func (s *Redirection) SetInfo(ctx context.Context, org *pb.RedirectData) (*pb.OrgUrl, error) {
	// just stub for a test
	if true {
		return &pb.OrgUrl{}, nil
	}
	return nil, errors.New("Error")
}
