package service

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func sqlConnect() (database *gorm.DB, err error) {
	DBMS := "mysql"
	DBUSER := os.Getenv("DBUSER")
	DBPASS := os.Getenv("DBPASSWORD")
	DBNAME := os.Getenv("DBNAME")
	HOST := os.Getenv("DBHOST")
	PROTOCOL := fmt.Sprintf("tcp(%s:3306)", HOST)

	CONNECT := DBUSER + ":" + DBPASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	log.Println(CONNECT)
	return gorm.Open(DBMS, CONNECT)
}

type Redirects struct {
	User         string `json:user`
	RedirectPath string `json:redirect_path`
	Org          string `json:org`
	Host         string `json:host`
	Comment      string `json:comment`
	Active       int    `json:active`
}

func main() {
	db, err := sqlConnect()
	if err != nil {
		log.Fatal(err)
	}
	db.Create(&Redirects{
		User:         "tako",
		RedirectPath: "shuya",
		Org:          "https://www.example.tv",
	})
}
