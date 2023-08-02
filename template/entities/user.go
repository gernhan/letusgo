package entities

import (
	"encoding/json"
	"time"
)

type User struct {
	ID          json.Number `json:"id"`
	Username    string      `json:"username"`
	Password    string      `json:"password"`
	CreatedTime time.Time   `json:"created_time"`
	UpdatedTime time.Time   `json:"updated_time"`
	UserType    json.Number `json:"user_type"`
	DateOfBirth time.Time   `json:"date_of_birth"`
}
