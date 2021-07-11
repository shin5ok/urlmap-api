package service

import (
	"fmt"
	"log"
	"math/rand"
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

func genRand(i int) string {
	src := []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	x := make([]rune, i)
	for n := range x {
		x[n] = src[rand.Intn(len(src))]
	}
	rand := string(x)
	fmt.Println(rand)
	return rand
}