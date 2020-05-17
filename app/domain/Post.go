package domain

import "time"

// type NullTime struct {
// 	Time  time.Time
// 	Valid bool
// }

// Post is a struct
type Post struct {
	ID            int        `json:"ID"`
	UserID        string     `json:"uid"`
	ScheduledDate time.Time  `json:"scheduled_date"`
	DeadlineDate  time.Time  `json:"deadline_date"`
	Title         string     `json:"title"`
	Contents      string     `json:"contents"`
	Fee           int        `json:"fee"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
}
