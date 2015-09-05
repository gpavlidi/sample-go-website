package models

import (
	"database/sql"
	"errors"
	"fmt"
)

type User struct {
	Id        string
	FirstName string
	LastName  string
}

type Users []User

//func NewUser(db *sql.DB, firstName, lastName string) *User {
//	return &User{FirstName: firstName, LastName: lastName}
//}

func NewOrUpdateUser(db *sql.DB, user *User) error {
	var err error
	var res sql.Result
	if user.Id == "" {
		res, err = db.Exec("INSERT INTO users(firstName, lastName) VALUES(?,?)", user.FirstName, user.LastName)
	} else {
		res, err = db.Exec("UPDATE users SET firstName=?, lastName=? WHERE id=?", user.FirstName, user.LastName, user.Id)
	}
	if err != nil {
		return err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return err
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowCnt != 1 {
		return errors.New(fmt.Sprintf("Row count (%s) and/or lastId (%s) are wrong.", lastId, rowCnt))
	}

	return nil
}

func GetAllUsers(db *sql.DB) (Users, error) {
	var users = Users{}

	rows, err := db.Query("select id, firstName, lastName from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user = User{}
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserById(db *sql.DB, id string) (*User, error) {
	var user = &User{Id: id}
	err := db.QueryRow("SELECT firstName, lastName FROM users WHERE id = ?", id).Scan(&user.FirstName, &user.LastName)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func DeleteUserById(db *sql.DB, id string) error {
	res, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return err
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowCnt != 1 {
		return errors.New(fmt.Sprintf("Row count (%s) and/or lastId (%s) are wrong.", lastId, rowCnt))
	}

	return nil
}
