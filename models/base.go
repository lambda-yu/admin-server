package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"time"
)

var db *gorm.DB

type ModelBase struct {
	ID        int `gorm:"primary_key" json:"id"`
	CreatedAt time.Time
	UpdateAt  time.Time
}

func Setup() {
	var err error
	db, err = gorm.Open(sqlite.Open("manager.db"), &gorm.Config{})

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	db.AutoMigrate(&User{})

}

func (m ModelBase) AftetUpdate(db *gorm.DB) (err error) {
	return db.Model(m).Where("ID = ?", m.ID).Update("UpdateAt", time.Now().Unix()).Error
}
