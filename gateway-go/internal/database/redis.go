package database

import (
    "context"
    "fmt"
    "log"

    "github.com/Zero-Coder-0/Rustmeme/gateway/internal/config"
    "github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

// ConnectRedis initializes the connection
func ConnectRedis() {
    addr := config.Get("REDIS_ADDR", "localhost:6379")

    RedisClient = redis.NewClient(&redis.Options{
        Addr: addr,
        DB:   0, // Default DB
    })

    // Test the connection immediately (The "Test Snippet" Rule)
    ctx := context.Background()
    _, err := RedisClient.Ping(ctx).Result()
    if err != nil {
        log.Fatalf("❌ Failed to connect to Redis: %v", err)
    }

    fmt.Println("✅ Connected to Redis successfully")
}
