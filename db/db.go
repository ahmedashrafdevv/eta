package db

import (
	"eta/config"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

var (
	DB *gorm.DB
)

func New() (*gorm.DB, error) {
	conStr := fmt.Sprintf("sqlserver://%s:%s@%s:1433?database=%s", config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), config.Config("DB_NAME"))
	DB, err := gorm.Open("mssql", conStr)
	if err != nil {
		fmt.Println("Failed to connect to external database")
		return nil, err
	}
	DB.LogMode(true)
	fmt.Println("Connection Opened to Database")
	return DB, nil
}
