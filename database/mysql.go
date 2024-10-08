package database

import (
	"fmt"
	"log"
	"os"
	projectModel "portfolioAPI/project/models"
	"sync"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	database *gorm.DB
	once     sync.Once
)

func initMySql() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

  databaseURL := os.Getenv("DB_DSN")
	db, err := gorm.Open(mysql.Open(databaseURL), &gorm.Config{
    NamingStrategy: schema.NamingStrategy{
      SingularTable: true,
      NoLowerCase: true,
    },
  })
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

  db.AutoMigrate(&projectModel.Project{}, &projectModel.Tag{}, &projectModel.ProjectStatus{})

	fmt.Println("Database connection successful!")
	database = db

}
func GetDBClient() *gorm.DB {
	once.Do(initMySql)
	return database
}
