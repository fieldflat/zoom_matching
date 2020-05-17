package domain

import "time"

// Participant is a struct
type Participant struct {
	ID        int        `json:"ID"`
	UserID    string     `json:"uid"`
	PostID    int        `json:"post_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
