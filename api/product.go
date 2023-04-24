package api

import (
	"99living-go/service"

	"github.com/gofiber/fiber/v2"
)

func GetAlbums(c *fiber.Ctx) error {
	album ,err:= service.GetAlbums()
	if err!=nil{
		return c.JSON(err)
		// fmt.Println(album)
		// fmt.Println(err)
	}
	return c.JSON(album)
}

// func CreateAlbums(c *fiber.Ctx) (err error) {
// 	album ,err:= service.CreateAlbums()
// 	album := new(model.Album)
// 	if err := c.BodyParser(&album); err != nil {
// 		return c.JSON(err)
// }
// 	return c.JSON(album)
// }