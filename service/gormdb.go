package service

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"

	"gorm.io/driver/mysql"
	_ "gorm.io/gorm/clause"
)

func SqlConnect(project string, p DbParams) (database *gorm.DB, err error) {

	protocol := fmt.Sprintf("tcp(%s:3306)", p.Dbhost)
	params := "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	connect := p.Dbuser + ":" + p.Dbpass + "@" + protocol + "/" + p.Dbname + params
	return gorm.Open(mysql.Open(connect), &gorm.Config{})
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
