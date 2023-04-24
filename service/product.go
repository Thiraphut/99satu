package service

import (
	"99living-go/database"
	"99living-go/model"

	"github.com/gofiber/fiber/v2"
)

func GetAlbums() (res model.Album,err error) {
	db := database.Connectdb
	db.Model(model.Album{})
	result:= db.Find(&res)
	return res,result.Error
}

// func CreateAlbums(album model.Album) (err error) {
// 	db, err := database.Connectdb
// 	if err != nil {
// 		return 0, err
// 	}
// 	if err != nil {
// 		return 0, err
// 	}
// 	return album.ID, err
// }

func CreateAlbums(c *fiber.Ctx) error {
	db := database.Connectdb
	album := new(model.Album)
	err := c.BodyParser(album)
	
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "massage": "Check your input", "data": err})
	}
	err = db.Create(&album).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "massage": "Could not create todo", "data": err})
	}
	return c.JSON(&album)
}