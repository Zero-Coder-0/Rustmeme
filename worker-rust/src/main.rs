use dotenv::dotenv;
use redis::AsyncCommands;
use serde::{Deserialize, Serialize}; // Restored JSON handling
use std::env;
use std::error::Error;

// Define the shape of the job (Must match the Go publisher's struct)
#[derive(Serialize, Deserialize, Debug)]
struct MemeJob {
    id: String,
    prompt: String,
    template: String,
    status: String,
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn Error>> {
    // 1. Load Environment Variables
    dotenv().ok();
    println!("⚙️  Worker initializing...");

    // 2. Connect to Redis
    let redis_url = env::var("REDIS_URL").unwrap_or("redis://127.0.0.1:6379/".to_string());
    let client = redis::Client::open(redis_url)?;
    let mut con = client.get_multiplexed_async_connection().await?;

    println!("✅ Worker connected to Redis!");
    println!("👂 Waiting for jobs in 'video_queue'...");

    // 3. The Infinite Loop (The Heartbeat)
    loop {
        // BLPOP: Blocking Left Pop. It waits here until a job arrives.
        // It returns a tuple: (queue_name, data)
        // We use "video_queue" as the target queue
        let result: Option<(String, String)> = con.blpop("video_queue", 0.0).await?;

        if let Some((_queue, data)) = result {
            // We got mail!
            println!("📩 Received Raw Data: {}", data);

            // Parse the JSON
            match serde_json::from_str::<MemeJob>(&data) {
                Ok(job) => {
                    println!("🔨 Processing Job ID: {}", job.id);
                    println!("   Command: '{}'", job.prompt);
                    // This is where we will eventually generate the image/video
                }
                Err(e) => println!("❌ Failed to parse job: {}", e),
            }
        }
    }
}
