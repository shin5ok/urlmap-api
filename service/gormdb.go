package service

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	DBMS     = "mysql"
	DBUSER   = os.Getenv("DBUSER")
	DBPASS   = os.Getenv("DBPASSWORD")
	DBNAME   = os.Getenv("DBNAME")
	HOST     = os.Getenv("DBHOST")
	PROTOCOL = fmt.Sprintf("tcp(%s:3306)", HOST)
)

func sqlConnect() (database *gorm.DB, err error) {

	PARAMS := "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	CONNECT := DBUSER + ":" + DBPASS + "@" + PROTOCOL + "/" + DBNAME + PARAMS
	return gorm.Open(DBMS, CONNECT)
}

type Redirects struct {
	gorm.Model
	User         string
	RedirectPath string
	Org          string
	Host         string
	Comment      string
	Active       int
}
