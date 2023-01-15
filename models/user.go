package models

type User struct {
	ID        string `json:"user_id,omitempty"`
	Name      string `json:"name"`
	BirthDay  string `json:"birthday"`
	Gender    string `json:"gender"`
	PhotoURL  string `json:"photo_url"`
	Time      int64  `json:"current_time"`
	Active    bool   `json:"active,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}

type UserSignup struct {
	Name     string `json:"name" binding:"required"`
	BirthDay string `json:"birthday" binding:"required"`
	Gender   string `json:"gender" binding:"required"`
	PhotoURL string `json:"photo_url" binding:"required"`
}
