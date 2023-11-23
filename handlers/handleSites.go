package handlers

import (
	filestorage "github.com/Oudwins/stackifyer/pkg/fileStorage"
	"github.com/gofiber/fiber/v2"
)

func PostNewSite(c *fiber.Ctx) error {

	titlesFileH, err := c.FormFile("titles")
	if err != nil {
		return err
	}

	f, err := titlesFileH.Open()
	if err != nil {
		return err
	}
	defer f.Close()

	err = filestorage.Api.SaveFile(&f, "test.png")
	if err != nil {
		print("error: ", err)
		return err
	}

	return c.Render("partials/alerts/alertSuccess", fiber.Map{})
}
