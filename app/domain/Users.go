package domain

// Users is a model
type Users struct {
	ID          int     `json:"ID"`
	UserID      string  `json:"user_id`
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	DisplayName string  `json:"display_name"`
	Email       *string `json:"email"`
	CreatedAt   int64   `json:"created_at"`
	UpdatedAt   int64   `json:"updated_at"`
}

// UsersForGet is a model
type UsersForGet struct {
	ID        int     `json:"ID"`
	UserID    string  `json:"user_id`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Email     *string `json:"email"`
}

// BuildForGet is a function
func (u *Users) BuildForGet() UsersForGet {
	user := UsersForGet{}
	user.ID = u.ID
	user.FirstName = u.FirstName
	user.LastName = u.LastName
	if u.Email != nil {
		user.Email = u.Email
	} else {
		empty := ""
		user.Email = &empty
	}
	return user
}
