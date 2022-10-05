package bootstrap

import (
	"fmt"
	"os"

	"crud-user-mvc/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitMySQL() (*gorm.DB, error) {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DBNAME"),
	)

	db, error := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if error != nil {
		return nil, error
	}

	fmt.Println("Successfully connected to database!")

	err := db.AutoMigrate(&model.User{})

	return db, err
}
