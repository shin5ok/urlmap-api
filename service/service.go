package service

import (
	"context"
	"log"
	pb "urlmap-api/pb"

	"github.com/jinzhu/gorm"
)

type Redirection struct{}

var dbConn *gorm.DB

func init() {}
func makeConn() *gorm.DB {
	if dbConn != nil {
		log.Println("using a stored connection")
		return dbConn
	}
	log.Println("init db connection")
	db, err := sqlConnect()
	if err != nil {
		log.Fatal(err)
	}
	dbConn = db
	return db
}

func (s *Redirection) GetInfoByUser(ctx context.Context, user *pb.User) (*pb.RedirectData, error) {
	u := user.User
	db := makeConn()

	type Redirects struct {
		Org     string
		User    string
		Active  int
		Host    string
		Comment string
	}
	var result Redirects
	// Field name in where args should be actual column name, not struct field
	status := db.Where("user = ?", u).First(&result)
	if status.Error != nil {
		log.Println(status.Error)
		return &pb.RedirectData{}, status.Error
	}
	log.Println(result)
	return &pb.RedirectData{
		Redirect: &pb.RedirectInfo{
			User:    result.User,
			Org:     result.Org,
			Host:    result.Host,
			Comment: result.Comment,
			Active:  int32(result.Active),
		},
	}, nil
}

func (s *Redirection) GetOrgByPath(ctx context.Context, path *pb.RedirectPath) (*pb.OrgUrl, error) {
	p := path.Path
	db := makeConn()

	type Redirects struct {
		Org string
	}
	var result Redirects
	// Field name in where args should be actual column name, not struct field
	status := db.Where("redirect_path = ?", p).First(&result)
	if status.Error != nil {
		log.Println(status.Error)
		return &pb.OrgUrl{}, status.Error
	}
	log.Println(result)

	return &pb.OrgUrl{Org: result.Org}, nil

}

func (s *Redirection) SetInfo(ctx context.Context, r *pb.RedirectData) (*pb.OrgUrl, error) {
	db := makeConn()
	redirect := Redirects{}
	redirect.RedirectPath = r.Redirect.RedirectPath
	redirect.User = r.Redirect.User
	redirect.Org = r.Redirect.Org
	redirect.Active = 1
	status := db.Create(&redirect)
	if status.Error != nil {
		log.Println(status.Error)
		return &pb.OrgUrl{}, status.Error
	}
	org := &pb.OrgUrl{Org: r.Redirect.Org}
	return org, nil
}
