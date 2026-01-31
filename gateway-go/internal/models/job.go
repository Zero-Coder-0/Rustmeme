package models

// MemeJob is the ticket we pass to Rust
type MemeJob struct {
    ID       string `json:"id"`
    Prompt   string `json:"prompt"`
    Template string `json:"template"` // e.g., "drake-hotline"
    Status   string `json:"status"`
}
