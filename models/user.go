package models

import "app/db"

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	// later added to a database
	query := `
	INSERT INTO user (email, password) 
	VALUES (?, ?)
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(u.Email, u.Password)

	if err != nil {
		return err
	}
	userID, err := result.LastInsertId()

	u.ID = userID
	return err
}
