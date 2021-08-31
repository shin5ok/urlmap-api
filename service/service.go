package service

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
	pb "urlmap-api/pb"

	"github.com/shin5ok/envorsecretm"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Redirection struct{}

var dbConn *gorm.DB
var Project = os.Getenv("PROJECT")

var c = envorsecretm.Config{Project}

type dbParams struct {
	dbms   string
	dbuser string
	dbpass string
	dbname string
	dbhost string
}

var v = dbParams{
	dbms:   "mysql",
	dbuser: c.Get("DBUSER"),
	dbpass: c.Get("DBPASSWORD"),
	dbname: c.Get("DBNAME"),
	dbhost: c.Get("DBHOST"),
}

func (v dbParams) makeConn() *gorm.DB {
	if dbConn != nil {
		log.Println("using a stored connection")
		return dbConn
	}
	log.Println("init db connection")
	db, err := sqlConnect(Project, v)
	if err != nil {
		log.Fatal(err)
	}
	dbConn = db
	return db
}

func (s *Redirection) GetInfoByUser(ctx context.Context, user *pb.User) (*pb.ArrayRedirectData, error) {
	u := user.User
	db := v.makeConn()

	pbResults := &pb.ArrayRedirectData{}
	resultSlice := []*pb.RedirectData{}

	// from ./gormdb.go as the same package
	var results []Redirects
	// Field name in where args should be actual column name, not struct field
	status := db.Where("user = ?", u).Find(&results)
	if status.Error != nil {
		log.Println(status.Error)
		return &pb.ArrayRedirectData{}, status.Error
	}

	if status.RowsAffected == 0 {
		log.Println("0 rows returned")
		return &pb.ArrayRedirectData{}, status.Error
	}

	for _, v := range results {

		resultSlice = append(resultSlice, &pb.RedirectData{
			Redirect: &pb.RedirectInfo{
				User:         v.User,
				Org:          v.Org,
				Comment:      v.Comment,
				Active:       int32(v.Active),
				RedirectPath: v.RedirectPath,
			},
		})
	}

	pbResults.Redirects = resultSlice
	return pbResults, nil
}

func (s *Redirection) GetOrgByPath(ctx context.Context, path *pb.RedirectPath) (*pb.OrgUrl, error) {
	p := path.Path
	db := v.makeConn()

	type RedirectOrg struct {
		Org      string
		NotifyTo string
	}
	var result RedirectOrg
	// Field name in where args should be actual column name, not struct field
	status := db.Table("redirects").Select("redirects.org, users.notify_to").Joins("join users on redirects.user = users.username").Where("redirect_path = ?", p).Scan(&result)

	if status.Error != nil {
		log.Println(status.Error)
		return &pb.OrgUrl{}, status.Error
	}
	fmt.Println(result)

	return &pb.OrgUrl{Org: result.Org, NotifyTo: result.NotifyTo}, nil

}

func (s *Redirection) SetInfo(ctx context.Context, r *pb.RedirectData) (*pb.OrgUrl, error) {
	db := v.makeConn()
	redirect := Redirects{}
	redirect.RedirectPath = r.Redirect.RedirectPath
	redirect.User = r.Redirect.User
	redirect.Org = r.Redirect.Org
	redirect.Comment = r.Redirect.Comment
	redirect.Active = 1
	t := time.Now()
	redirect.BeginAt = &t
	status := db.Create(&redirect)
	if status.Error != nil {
		log.Printf("%+v", redirect)
		log.Println(status.Error)
		return &pb.OrgUrl{}, status.Error
	}
	org := &pb.OrgUrl{Org: r.Redirect.Org}
	return org, nil
}

func (s *Redirection) SetUser(ctx context.Context, r *pb.User) (*pb.User, error) {
	db := v.makeConn()
	user := Users{Username: r.User, NotifyTo: r.NotifyTo}
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "username"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"username": r.User, "notify_to": r.NotifyTo}),
	}).Create(&user)
	return &pb.User{User: user.Username, NotifyTo: user.NotifyTo}, nil
}

func (s *Redirection) RemoveUser(ctx context.Context, r *pb.User) (*emptypb.Empty, error) {
	db := v.makeConn()
	redirect := Redirects{}
	redirect.User = r.User
	status := db.Delete(&redirect)
	if status.Error != nil {
		log.Printf("%+v", redirect)
		log.Println(status.Error)
		return nil, status.Error
	}
	return nil, nil
}
