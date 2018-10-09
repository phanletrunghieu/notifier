package domain

// Device .
type Device struct {
	Model
	Token  string `json:"token,omitempty"`
	Name   string `json:"name"`
	UserID UUID   `sql:",type:uuid" json:"user_id"`
}
