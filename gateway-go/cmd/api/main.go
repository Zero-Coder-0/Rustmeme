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

        // ... inside main() ...

    // NEW: The Generate Route
    app.Post("/generate", func(c *fiber.Ctx) error {
        // 1. Read User Input
        type Request struct {
            Prompt string `json:"prompt"`
        }
        var req Request
        if err := c.BodyParser(&req); err != nil {
            return c.Status(400).SendString("Invalid Input")
        }

        // 2. Create the Job Ticket
        job := models.MemeJob{
            ID:       "12345", // Hardcoded for testing
            Prompt:   req.Prompt,
            Template: "default_template",
            Status:   "pending",
        }

        // 3. Send to Redis
        if err := database.PushJobToQueue(job); err != nil {
            return c.Status(500).SendString("Failed to queue job")
        }

        return c.JSON(fiber.Map{
            "message": "Job sent to the Muscle!",
            "job_id":  job.ID,
        })
    })

    // ... app.Listen ...



    // 4. Start Server
    port := config.Get("PORT", "3000")
    log.Fatal(app.Listen(":" + port))
}
