package data

import(
	"time"
)

type UserData struct {
    ID        string    `json:"id"` 
    Username  string    `json:"username"`
    Email     string    `json:"email"`
    Password  string    `json:"password"` 
    CreatedAt time.Time `json:"created_at"`
}
