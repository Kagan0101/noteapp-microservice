package data

import(
	"time"
)

type NoteData struct {
    ID        string    `json:"id"` // UUID
    UserID    string    `json:"user_id"` // notu kimin oluşturduğu
    Title     string    `json:"title"`
    Content   string    `json:"content"`
    Tags      []string  `json:"tags"` // ["iş", "gizli", "günlük"] gibi
    IsPinned  bool      `json:"is_pinned"` // sabitlenen notlar
    IsArchived bool     `json:"is_archived"` // arşivlenmiş notlar
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
