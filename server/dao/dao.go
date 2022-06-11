package dao

import (
	"database/sql"
	"fmt"
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB
var MDB *sql.DB
var RDB *redis.Client

func InitDB() {
	dsn := "root:123@tcp(127.0.0.1:3306)/chess?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return
	}
	DB = db
}

func InitMDB() {
	db, err := sql.Open("mysql", "root:045226@tcp(localhost:3306)/chess?charset=utf8mb4&parseTime=True")
	if err != nil {
		fmt.Println("failed", err)
		panic(err)
	}
	MDB = db
}

func InitRDB() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := RDB.Ping().Result()
	if err != nil {
		log.Fatal(err)
		return
	}
}
