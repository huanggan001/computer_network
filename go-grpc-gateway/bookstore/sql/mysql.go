package sql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"grpc/bookstore/data"
	"time"
)

func InitMySql() (*gorm.DB, error) {
	dsn := "admin:123456@tcp(127.0.0.1:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	err = db.AutoMigrate(&data.Shelf{}, &data.Book{})
	if err != nil {
		return nil, err
	}

	return db, err
}
