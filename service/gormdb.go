package service

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/shin5ok/envorsecretm"
	"gorm.io/gorm"

	"gorm.io/driver/mysql"
	_ "gorm.io/gorm/clause"
)

var c = envorsecretm.Config{Project}

var (
	DBMS     = "mysql"
	DBUSER   = c.Get("DBUSER")
	DBPASS   = c.Get("DBPASSWORD")
	DBNAME   = c.Get("DBNAME")
	HOST     = c.Get("DBHOST")
	PROTOCOL = fmt.Sprintf("tcp(%s:3306)", HOST)
)

func sqlConnect(project string) (database *gorm.DB, err error) {

	PARAMS := "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	CONNECT := DBUSER + ":" + DBPASS + "@" + PROTOCOL + "/" + DBNAME + PARAMS
	return gorm.Open(mysql.Open(CONNECT), &gorm.Config{})
}

type Redirects struct {
	gorm.Model
	User         string `gorm:"foreignKey:UserName"`
	RedirectPath string
	Org          string
	Host         string
	Comment      string
	Active       int
	BeginAt      *time.Time // it will insert NULL when no value is specified
	EndAt        *time.Time // it will insert NULL when no value is specified
}

type Users struct {
	// gorm.Model
	Username string
	NotifyTo string
}
