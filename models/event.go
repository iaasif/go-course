package models

import "time"

type Event struct {
	ID          int
	Name        string
	Description string
	Localtion   string
	DateTime    time.Time
	UserID      int
}

var events = []Event{}

func (e Event) Save() {
	// later added to a database
	events = append(events, e)
}
