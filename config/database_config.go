package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Database *gorm.DB


func ConnectToDatabase() {
	dsn := "user=postgres password=Ak01888714929 dbname=crud port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	Database, err = gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
		
	}),&gorm.Config{
		Logger: logger.Default.LogMode(logger.Info) ,
	})
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("Connected to database")
	}
}