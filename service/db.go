package main

import (
	"log"
	"os"
	"sort"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	// _ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	gorm.Model
	Text   string
	Status string
}

var dbFile string = os.Getenv("DBFILE")
var isInit bool = false

//DB初期化
func dbInit() error {
	if isInit {
		return nil
	}
	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		log.Printf("cannot open %s\n", err)
		return err
	}
	{
		err := db.AutoMigrate(&Todo{})
		if err != nil {
			log.Fatal(err)
		}
	}
	isInit = true
	return nil
}

//DB追加
func dbInsert(text string, status string) error {
	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		log.Println("cannot open")
		return err
	}
	db.Create(&Todo{Text: text, Status: status})
	//defer db.Close()()
	return nil
}

//DB更新
func dbUpdate(id int, text string, status string) error {
	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		log.Println("cannot open")
		return err
	}
	var todo Todo
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	db.Save(&todo)
	// db.Close()
	return nil
}

//DB削除
func dbDelete(id int) error {
	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		log.Println("cannot open")
		return err
	}
	var todo Todo
	db.First(&todo, id)
	db.Delete(&todo)
	// db.Close()
	return nil
}

//DB全取得
func dbGetAll() ([]Todo, error) {
	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		log.Println("cannot open")
		return nil, err
	}
	var todos []Todo
	db.Order("created_at desc").Find(&todos)

	// Sorting todos[]
	// https://qiita.com/Sekky0905/items/2d5ccd6d076106e9d21c
	sort.Slice(todos, func(i, j int) bool { return todos[i].ID < todos[j].ID })
	return todos, nil
}

//DB一つ取得
func dbGetOne(id int) (Todo, error) {
	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		log.Println("cannot open")
		return Todo{}, err
	}
	var todo Todo
	{
		err := db.First(&todo, id).Error
		// db.Close()
		if err != nil {
			log.Println(err)
			return Todo{}, err
		}
	}
	return todo, nil
}
