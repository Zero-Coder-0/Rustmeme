package database

import (
    "context"
    "encoding/json"
    "github.com/Zero-Coder-0/Rustmeme/gateway/internal/models"
)

// PushJobToQueue sends the job to the "meme_queue" list
func PushJobToQueue(job models.MemeJob) error {
    ctx := context.Background()
    
    // Convert the struct to JSON bytes
    data, err := json.Marshal(job)
    if err != nil {
        return err
    }

    // RPUSH adds the element to the tail of the list
    return RedisClient.RPush(ctx, "meme_queue", data).Err()
}
