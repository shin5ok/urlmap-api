package service

import (
	"fmt"
	"log"
	"os"
	"time"

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
	gorm.Model
	UserName string
	NotifyTo string
}

/* for test */
func main() {
	db, err := sqlConnect()
	if err != nil {
		log.Fatal(err)
	}
	db.Create(&Redirects{
		User:         "tako",
		RedirectPath: "ika",
		Org:          "https://www.example.tv",
		Active:       1,
	})
	results := []*Redirects{}
	error := db.Find(&results).Error
	if error != nil || len(results) == 0 {
		return
	}
	for i, r := range results {
		fmt.Printf("%d: %s, %s, %s, %d\n", i, r.User, r.RedirectPath, r.Org, r.Active)

	}
}
