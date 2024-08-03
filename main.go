package main

import (
    "log"
    "notes/config"
    "notes/database"
    "notes/routes"
    "github.com/gofiber/fiber/v2"
)

func main() {
    cfg := config.LoadConfig()

    database.MigrateDB(cfg)

    db := database.GetDBConnection(cfg)
    defer db.Close()

    app := fiber.New()

    routes.SetupRoutes(app, db)

    log.Fatal(app.Listen(":8080"))
}
