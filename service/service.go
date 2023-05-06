package service

import (
	"context"
	"fmt"
	"os"
	"time"

	pb "github.com/shin5ok/urlmap-api/pb"

	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/rs/zerolog"
	log "github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DbParams struct {
	Dbms   string
	Dbuser string
	Dbpass string
	Dbname string
	Dbhost string
}

type Redirection struct {
	DbParams
}

var (
	dbConn  *gorm.DB
	Project = os.Getenv("PROJECT")
)

func init() {
	log.Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	zerolog.LevelFieldName = "severity"
	zerolog.TimestampFieldName = "timestamp"
	zerolog.TimeFieldFormat = time.RFC3339Nano
}

func New(dbParams DbParams) Redirection {
	return Redirection{dbParams}
}

type DBOperator interface {
	makeConn() *gorm.DB
}

func (s *Redirection) makeConn() *gorm.DB {
	if dbConn != nil {
		log.Info().Msg("using a stored connection")
		return dbConn
	}
	log.Info().Msg("init db connection")
	db, err := SqlConnect(Project, s.DbParams)
	if err != nil {
		log.Fatal().Err(err)
	}
	dbConn = db
	return db
}

func (s *Redirection) GetInfoByUser(ctx context.Context, user *pb.User) (*pb.ArrayRedirectData, error) {
	u := user.User
	db := s.makeConn()

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
	db := s.makeConn()

	var result pb.OrgUrl
	// Field name in where args should be actual column name, not struct field
	status := db.Table("redirects").Debug().
		Select("redirects.org as Org , users.notify_to as NotifyTo , users.slack_url as SlackUrl, users.email as Email").
		Joins("join users on redirects.user = users.username").
		Where("redirect_path = ?", p).
		Scan(&result)

	if status.Error != nil {
		log.Error().Err(status.Error)
		return &pb.OrgUrl{}, status.Error
	}
	fmt.Println(&result)
	grpc_ctxtags.Extract(ctx).Set("result", &result)

	// return &pb.OrgUrl{Org: result.Org, NotifyTo: result.NotifyTo}, nil
	return &result, nil

}

func (s *Redirection) SetInfo(ctx context.Context, r *pb.RedirectData) (*pb.OrgUrl, error) {
	db := s.makeConn()
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
	db := s.makeConn()
	user := Users{Username: r.User, NotifyTo: r.NotifyTo}
	db.Debug().Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "username"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"username": r.User, "notify_to": r.NotifyTo}),
	}).Create(&user)
	return &pb.User{User: user.Username, NotifyTo: user.NotifyTo}, nil
}

func (s *Redirection) RemoveUser(ctx context.Context, r *pb.User) (*emptypb.Empty, error) {
	db := s.makeConn()
	user := Users{}
	db.Debug().
		Where("UserName = ?", r.User).
		Delete(&user)
	return &emptypb.Empty{}, nil
}

func (s *Redirection) ListUsers(ctx context.Context, empty *emptypb.Empty) (*pb.Users, error) {
	var userlist []*pb.User
	users := &pb.Users{}
	db := s.makeConn()
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

func (s *Redirection) PingPongMessage(ctx context.Context, message *pb.Message) (*pb.Message, error) {
	message.ShowModeOneof = &pb.Message_Mode{Mode: "pong"}
	return message, nil
}
