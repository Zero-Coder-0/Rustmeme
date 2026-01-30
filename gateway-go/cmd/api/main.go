package main

import (
    "log"

    "github.com/Zero-Coder-0/Rustmeme/gateway/internal/config"
    "github.com/Zero-Coder-0/Rustmeme/gateway/internal/database"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
    // 1. Load Config
    config.LoadEnv()

    // 2. Connect to Memory (Redis)
    database.ConnectRedis()

    // 3. Initialize App
    app := fiber.New(fiber.Config{
        AppName: "ZME-X Gateway",
    })
    app.Use(logger.New())

    app.Get("/health", func(c *fiber.Ctx) error {
        return c.Status(200).JSON(fiber.Map{
            "status": "online",
            "db":     "connected",
        })
    })

    // 4. Start Server
    port := config.Get("PORT", "3000")
    log.Fatal(app.Listen(":" + port))
}
