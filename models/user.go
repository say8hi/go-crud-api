package models

type User struct {
	ID       int    `json:"id,omitempty"`
	Username string `json:"username"`
	FullName string `json:"full_name,omitempty"`
}
