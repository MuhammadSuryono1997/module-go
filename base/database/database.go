package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/MuhammadSuryono1997/framework-okta/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func (dbConfig DBConfig) GetConnectionString() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

var database *gorm.DB

func Init(config DBConfig) {
	db, err := gorm.Open("mysql", config.GetConnectionString())
	if err != nil {
		fmt.Println("error ", err)
	}
	database = db
	pingTicker := time.NewTicker(15 * time.Second)
	pingDone := make(chan bool)
	go func() {
		for {
			select {
			case <-pingDone:
				return
			case <-pingTicker.C:
				b := pingDb(db.DB())
				if !b {
					pingDone <- true
				}
			}
		}
	}()
}

func GetDb() *gorm.DB {
	return database
}

func pingDb(db *sql.DB) bool {
	er := db.Ping()
	if er != nil {
		log.Print("mysql error ping", er)
		return false
	} else {
		log.Print("mysql success ping")
		return true
	}
}

func CreateConnection() {

	port, err := strconv.Atoi(os.Getenv("DATABASE_POST"))

	if err != nil {
		fmt.Println(string(utils.ColorYellow()), err)
	}

	dbConfig := DBConfig{
		Host:     os.Getenv("DATABASE_HOST"),
		Port:     port,
		User:     os.Getenv("DATABASE_USERNAME"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		DBName:   os.Getenv("DATABASE_NAME"),
	}

	Init(dbConfig)
}
