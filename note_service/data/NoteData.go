package data

import(
	"time"
)

type NoteData struct {
    ID        string    `json:"id"` 
    UserID    string    `json:"user_id"` 
    Title     string    `json:"title"`
    Content   string    `json:"content"`
    Tags      []string  `json:"tags"` 
    IsPinned  bool      `json:"is_pinned"` 
    IsArchived bool     `json:"is_archived"` 
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
