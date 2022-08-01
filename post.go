package userDB

import (
	"database/sql"
	"fmt"
)

// Create a new user that is part of a `ban` and has the specified `emails`
func NewUser(db *sql.DB, name string, emails []string, ban int) error {
	// Gebruiker toevoegen
	insertQuery := fmt.Sprintf("insert into Users (Name) values (%s)", name)
	result, err := db.Exec(insertQuery)
	if err != nil {
		return err
	}

	// Get user id
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	// Email(s) toevoegen
	for _, email := range emails {
		err := AddUserEmail(db, id, email)
		if err != nil {
			return err
		}
	}

	// Ban toevoegen
	err = AddUserBan(db, id, ban)

	return err
}

// Remove a user from the database, including all its emails and ban entries
func RemoveUser(dq *sql.DB, userId int64) error {
	panic("unimplemented")
}

// Add an email to a user
func AddUserEmail(db *sql.DB, userId int64, email string) error {
	query := fmt.Sprintf("insert into Emails (Email_Id, Email) (values %d, %s)", userId, email)
	_, err := db.Exec(query)
	return err
}

// Remove an email from a user
func RemoveUserEmail(db *sql.DB, userId int64, email string) error {
	panic("unimplemented")
}

// Add a ban to a user
func AddUserBan(db *sql.DB, userId int64, ban int) error {
	query := fmt.Sprintf("insert into User_Ban (User_Id, Ban_Id) values (%d, %d)", userId, ban)
	_, err := db.Exec(query)
	return err
}

// Remove a ba from a user
func RemoveUserBan(db *sql.DB, userId, int64, ban int) error {
	panic("unimplementted")
}
