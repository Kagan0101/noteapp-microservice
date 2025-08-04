package data

import(
	"time"
)

type UserData struct {
    ID        string    `json:"id"` // UUID
    Username  string    `json:"username"`
    Email     string    `json:"email"`
    Password  string    `json:"password"` // hashli saklanır
    CreatedAt time.Time `json:"created_at"`
}
