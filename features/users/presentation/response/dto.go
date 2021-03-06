package response

import "time"

type User struct {
	ID           uint      `json:"id" form:"id"`
	Name         string    `json:"name" form:"name"`
	Email        string    `json:"email" form:"email"`
	CreatedAt    time.Time `json:"created_at" form:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" form:"updated_at"`
}
