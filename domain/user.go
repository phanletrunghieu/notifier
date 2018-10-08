package domain

// User .
type User struct {
	Model
	Name string `json:"name,omitempty"`
}
