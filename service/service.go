package service

import (
	"context"
	"fmt"
	"os"
	"time"
	pb "urlmap-api/pb"

	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	log "github.com/rs/zerolog/log"
	"github.com/shin5ok/envorsecretm"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type dbParams struct {
	dbms   string
	dbuser string
	dbpass string
	dbname string
	dbhost string
}

type Redirection struct{}

var dbConn *gorm.DB
var Project = os.Getenv("PROJECT")
var c = envorsecretm.Config{ProjectId: Project}

var v = dbParams{
	dbms:   "mysql",
	dbuser: c.Get("DBUSER"),
	dbpass: c.Get("DBPASS"),
	dbname: c.Get("DBNAME"),
	dbhost: c.Get("DBHOST"),
}

func (v dbParams) makeConn() *gorm.DB {
	if dbConn != nil {
		log.Info().Msg("using a stored connection")
		return dbConn
	}
	log.Info().Msg("init db connection")
	db, err := sqlConnect(Project, v)
	if err != nil {
		log.Fatal().Err(err)
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
	status := db.Debug().Where("user = ?", u).Find(&results)
	if status.Error != nil {
		log.Error().Err(status.Error)
		return &pb.ArrayRedirectData{}, status.Error
	}

	if status.RowsAffected == 0 {
		log.Info().Msg("0 rows returned")
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

	grpc_ctxtags.Extract(ctx).Set("results", resultSlice)

	pbResults.Redirects = resultSlice
	return pbResults, nil
}

func (s *Redirection) GetOrgByPath(ctx context.Context, path *pb.RedirectPath) (*pb.OrgUrl, error) {
	p := path.Path
	db := v.makeConn()

	var result pb.OrgUrl
	// Field name in where args should be actual column name, not struct field
	status := db.Table("redirects").Debug().
		Select("redirects.org as Org , users.notify_to as NotifyTo").
		Joins("join users on redirects.user = users.username").
		Where("redirect_path = ?", p).
		Scan(&result)

	if status.Error != nil {
		log.Error().Err(status.Error)
		return &pb.OrgUrl{}, status.Error
	}
	fmt.Println(result)
	grpc_ctxtags.Extract(ctx).Set("result", result)

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
	status := db.Debug().Create(&redirect)
	if status.Error != nil {
		log.Error().Msgf("%+v", redirect)
		log.Error().Err(status.Error)
		return &pb.OrgUrl{}, status.Error
	}
	org := &pb.OrgUrl{Org: r.Redirect.Org}
	return org, nil
}

func (s *Redirection) SetUser(ctx context.Context, r *pb.User) (*pb.User, error) {
	db := v.makeConn()
	user := Users{Username: r.User, NotifyTo: r.NotifyTo}
	db.Debug().Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "username"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"username": r.User, "notify_to": r.NotifyTo}),
	}).Create(&user)
	return &pb.User{User: user.Username, NotifyTo: user.NotifyTo}, nil
}

func (s *Redirection) RemoveUser(ctx context.Context, r *pb.User) (*emptypb.Empty, error) {
	db := v.makeConn()
	user := Users{}
	db.Debug().
		Where("UserName = ?", r.User).
		Delete(&user)
	return &emptypb.Empty{}, nil
}

func (s *Redirection) ListUsers(ctx context.Context, empty *emptypb.Empty) (*pb.Users, error) {
	var userlist []*pb.User
	users := &pb.Users{}
	db := v.makeConn()
	status := db.Table("users").
		Debug().
		// userlist has 'User' but table has 'username', so need to use 'as' SQL sentence
		Select("username as user", "notify_to").
		Scan(&userlist)
	log.Info().Msgf("%+v", userlist)
	users.Users = userlist

	if status.Error != nil {
		log.Error().Err(status.Error)
		return &pb.Users{}, status.Error
	}

	return users, nil
}
