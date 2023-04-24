package database

import (
	"99living-go/model"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)



var Connectdb *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DB_LIVING")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	Connectdb = db
	db.AutoMigrate(&model.Product{},&model.Album{},&model.AlbumExtraImage{})
	
}
