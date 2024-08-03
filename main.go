package main

import (
    "log"
    "notes/config"
    "notes/database"
    "notes/routes"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/swagger"
    _ "notes/docs" // This is necessary for swag to find your docs
)

// @title Fiber Swagger Example API
// @version 1.0
// @description This is a sample server for a notes app.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
    cfg := config.LoadConfig()

    database.MigrateDB(cfg)

    db := database.GetDBConnection(cfg)
    defer db.Close()

    app := fiber.New()

    routes.SetupRoutes(app, db)

    // Setup Swagger
    app.Get("/swagger/*", swagger.HandlerDefault)

    log.Fatal(app.Listen(":8080"))
}
