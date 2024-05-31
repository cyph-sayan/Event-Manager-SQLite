package models

import (
	"events-management/database"
	"events-management/utility"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user *User) SaveUser() error {
	pass, err := utility.HashPassword(user.Password)
	if err != nil {
		return err
	}
	query := `
		INSERT INTO users (email, password) values (?,?)
	`
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(user.Email, pass)
	if err != nil {
		return err
	}
	user.Id, err = result.LastInsertId()
	return err
}

func (user *User) ValidateUser() (int64, error) {
	query := `
		SELECT id, password from users where email = ?
	`
	result := database.DB.QueryRow(query, user.Email)
	var retrievedPassWord string
	var id int64
	err := result.Scan(&id, &retrievedPassWord)
	if err != nil {
		return 0, err
	}

	err = utility.ValidatePassword(user.Password, retrievedPassWord)

	if err != nil {
		return 0, err
	}

	return id, nil

}
