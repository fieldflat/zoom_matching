package domain

import "time"

// Users is a model
type Users struct {
	ID          int        `json:"ID"`
	UserID      string     `json:"uid"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	DisplayName string     `json:"display_name"`
	Email       string     `json:"email"`
	Password    string     `json:"password"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

// UsersForGet is a model
type UsersForGet struct {
	ID        int    `json:"ID"`
	UserID    string `json:"uid"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// BuildForGet is a function
func (u *Users) BuildForGet() UsersForGet {
	user := UsersForGet{}
	user.ID = u.ID
	user.FirstName = u.FirstName
	user.LastName = u.LastName
	if u.Email != "" {
		user.Email = u.Email
	} else {
		user.Email = ""
	}
	return user
}
