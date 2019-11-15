package db

import (
	"fmt"

	"FirstGoProject/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// New Creates a new DB named example.db
func New() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./example.db")
	if err != nil {
		fmt.Println("storage err: ", err)
	}
	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)
	return db
}

// TODO: Learn this eventually
// func TestDB() *gorm.DB {
// 	db, err := gorm.Open("sqlite3", "./../example_test.db")
// 	if err != nil {
// 		fmt.Println("storage err: ", err)
// 	}
// 	db.DB().SetMaxIdleConns(3)
// 	db.LogMode(false)
// 	return db
// }

// func DropTestDB() error {
// 	if err := os.Remove("./../example_test.db"); err != nil {
// 		return err
// 	}
// 	return nil
// }

//TODO: Learn what this does
// AutoMigrate ...
func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.Champ{},
		&model.Origin{},
		// &model.User{},
		// &model.Follow{},
		// &model.Article{},
		// &model.Comment{},
		// &model.Tag{},
	)
}
