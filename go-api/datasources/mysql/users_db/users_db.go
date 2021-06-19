package users_db

import (
	"fmt"
	"os"

	"github.com/mrmtsu/go_api/domain/articles"
	"github.com/mrmtsu/go_api/domain/users"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DbConnect() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database_name := os.Getenv("DB_DATABASE_NAME")

	dns := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database_name + "?charset=utf8"
	database, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}

	DB = database

	database.AutoMigrate(&users.User{}, &articles.Article{})
	fmt.Println("ok")
}
