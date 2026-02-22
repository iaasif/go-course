package models

import "time"

type Event struct {
	ID          int    `json:id`
	Name        string `json:name`
	Description string
	Localtion   string
	DateTime    time.Time
	UserID      int
}
