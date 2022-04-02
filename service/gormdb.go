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

	PROTOCOL := fmt.Sprintf("tcp(%s:3306)", p.dbhost)
	PARAMS := "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	CONNECT := p.dbuser + ":" + p.dbpass + "@" + PROTOCOL + "/" + p.dbname + PARAMS
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
