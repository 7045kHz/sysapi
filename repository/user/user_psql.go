package userRepository

import (
	"sysapi/models"
	"database/sql"
	"fmt"
	"log"
)

type UserRepository struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (u UserRepository) Signup(db *sql.DB, user models.User) models.User {
	stmt := "insert into users (email, password,token,level) values($1, $2, $3 ,$4) ON CONFLICT (email) DO UPDATE SET password = $2 RETURNING id;"
	err := db.QueryRow(stmt, user.Email, user.Password, "", 0).Scan(&user.ID)
	logFatal(err)

	user.Password = ""
	fmt.Printf("UserRep: %v\n", user)
	return user
}

func (u UserRepository) Login(db *sql.DB, user models.User) (models.User, error) {
	row := db.QueryRow("select * from users where email=$1", user.Email)
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Level, &user.Token)

	if err != nil {
		return user, err
	}

	return user, nil
}
