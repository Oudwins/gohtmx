package main

import (
	"log"

	"github.com/Oudwins/stackifyer/db"
	"github.com/Oudwins/stackifyer/handlers"
	filestorage "github.com/Oudwins/stackifyer/pkg/fileStorage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Missing .env file")
	}

	if err := db.Init(); err != nil {
		log.Fatal(("Init DB"))
	}

	if err := filestorage.Init(); err != nil {
		log.Fatal(err)
	}

	htmlEngine := html.New("www/web", ".html")

	app := fiber.New(fiber.Config{
		Views:        htmlEngine,
		ErrorHandler: handlers.DefaultErrorHandler,
		// BodyLimit default 10mb
	})

	app.Use(handlers.WithAuthenticated)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Index Page",
		}, "layouts/base")
	})

	// Auth
	app.Post("/signup", handlers.SignUp)
	app.Post("/signin", handlers.SignIn)
	app.Get("/signup", func(c *fiber.Ctx) error {
		return c.Render("signup", fiber.Map{
			"Title": "Sign Up",
		}, "layouts/base")
	})
	app.Get("/signin", func(c *fiber.Ctx) error {
		return c.Render("signin", fiber.Map{
			"Title": "Sign In",
		}, "layouts/base")
	})

	// Restricted
	app.Use(handlers.Restricted)
	app.Post("/signout", handlers.SignOut)

	app.Get("/dashboard", func(c *fiber.Ctx) error {
		return c.Render("dashboard/index", fiber.Map{}, "layouts/base")
	})
	app.Post("/sites", handlers.PostNewSite)
	// files
	fileEnpoints := app.Group("/files")
	fileEnpoints.Post("/images", handlers.PostUploadImages)

	app.Listen(":3000")
}
