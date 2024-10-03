package repository

import (
	"database/sql"
	"fmt"
	"github.com/IndySeh/go-crud-application/pkg/logging"
	"github.com/IndySeh/go-crud-application/pkg/types"
	"log/slog"
	"time"
)

func FetchUsersFromDB(db *sql.DB) ([]*types.User, error) {
	query := "Select * from users"

	logging.DbLogger.Info("Executing query", slog.String("query", query))

	startTime := time.Now()

	rows, err := db.Query(query)

	if err != nil {
		logging.DbLogger.Error("Error executing query", slog.String("query", query), slog.String("error", err.Error()))
		return nil, fmt.Errorf("error in executing query: %v", err)
	}

	defer rows.Close()

	duration := time.Since(startTime)
	logging.DbLogger.Info("Query Executed", slog.String("query", query), slog.Duration("duration", duration))

	// Slice to hold the pointers of users
	var users []*types.User

	for rows.Next() {
		user := &types.User{}
		err := rows.Scan(&user.Id, &user.Email, &user.Name)
		if err != nil {
			return nil, fmt.Errorf("error: %v", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error: %v", err.Error())
	}

	return users, nil
}

func FetchUserFromDB(db *sql.DB, id int) (*types.User, error) {

	query := "select * from users where id = ?"
	row := db.QueryRow(query, id)

	user := &types.User{}
	err := row.Scan(&user.Id, &user.Name, &user.Email)
	if err != nil {
		return nil, fmt.Errorf("error: %v", err.Error())
	}
	return user, nil
}

func DeleteUserFromDB(db *sql.DB, Id int) error {
	user := &types.User{}

	findUser := `SELECT * FROM users where id = ?`
	err := db.QueryRow(findUser, Id).Scan(&user.Id, &user.Name, &user.Email)
	if err != nil {
		return fmt.Errorf("user not found %v", err)
	}

	deleteUser := `DELETE FROM users WHERE id = ?`
	res, err := db.Exec(deleteUser, Id)
	if err != nil {
		return fmt.Errorf("error in deleting user")
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error in deleting rows")
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows were affected ")
	}
	return nil
}

func InsertUserInDB(db *sql.DB, name, email string) error {
	query := `INSERT INTO users (name, email) VALUES (?,?)`
	_, err := db.Exec(query, name, email)
	if err != nil {
		return fmt.Errorf("error inserting data %v", err.Error())
	}
	return nil
}

func UpdateUserInDB(db *sql.DB, user *types.User) error {
	query := `UPDATE users SET name = ?, email = ? WHERE id = ?`
	_, err := db.Exec(query, user.Name, user.Email, user.Id)
	if err != nil {
		return err
	}
	return nil
}

// TODO Check User before going ahead with request.
func UserExist()  {
	
}