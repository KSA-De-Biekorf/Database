package userDB

import (
	"database/sql"
	"fmt"
)

// Database entry
type Entry struct {
	Id    int
	Name  string
	BanId int
	Ban   string
	Email string
}

// Fetch a query string (not exec; returns the query)
func fetchQuery(db *sql.DB, query string) ([]Entry, error) {
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []Entry
	for rows.Next() {
		var id int
		var name string
		var banId int
		var ban string
		var email string
		if err := rows.Scan(&id, &name, &banId, &ban, &email); err != nil {
			return nil, err
		}
		entry := Entry{id, name, banId, ban, email}
		entries = append(entries, entry)
	}

	return entries, nil
}

// Fetch every user from the database
func FetchAllUsers(db *sql.DB) ([]Entry, error) {
	query := "SELECT Users.User_Id as id, Users.Name as name, Bannen.Ban_id as ban_id, Bannen.Naam as ban, Emails.Email as email FROM Users " +
		"INNER JOIN User_Ban ON User_Ban.User_Id = Users.User_Id " +
		"INNER JOIN Bannen ON User_Ban.Ban_Id = Bannen.Ban_Id " +
		"INNER JOIN Emails ON Emails.Email_id = Users.User_Id"

	return fetchQuery(db, query)
}

// Fetch all users in a `ban`. Use the `ban` enum
func FetchBan(db *sql.DB, ban string) ([]Entry, error) {
	query := "SELECT Users.User_Id as id, Users.Name as name, Bannen.Ban_id as ban_id, Bannen.Naam as ban, Emails.Email as email FROM Users " +
		"INNER JOIN User_Ban ON User_Ban.User_Id = Users.User_Id " +
		"INNER JOIN Bannen ON User_Ban.Ban_Id = Bannen.Ban_Id " +
		fmt.Sprintf("AND Bannen.Naam = '%s' ", ban) +
		"INNER JOIN Emails ON Emails.Email_id = Users.User_Id"

	return fetchQuery(db, query)
}
