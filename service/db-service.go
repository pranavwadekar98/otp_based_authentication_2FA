package service

import (
	"errors"
	"fmt"

	"example.com/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DBService interface {
	Save(entity.SignUpData) error
	Find(entity.SignUpData) error
	// CloseDB()
}

type database struct {
	conn *gorm.DB
}

func New() DBService {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	db.AutoMigrate(&entity.SignUpData{})
	return &database{
		conn: db,
	}
}

// func (db *database) CloseDB() {
// 	err := db.conn.Close()
// 	if err != nil {
// 		panic("Failed to close the database")
// 	}
// }

func (db *database) Save(data entity.SignUpData) error {
	err := db.conn.Create(&data)
	fmt.Println("--->", err)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (db *database) Find(data entity.SignUpData) error {
	model := entity.SignUpData{}
	err := db.conn.Where("phone = ? AND password = ?", data.Email, data.Password).Find(&model)
	if err.Error != nil {
		return err.Error
	}
	if model.Email == "" && model.Password == "" {
		return errors.New("User does not exists")
	}
	return nil
}
