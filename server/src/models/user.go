package models

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/shailesz/cli-chat-golang-server/src/helpers"
	"github.com/shailesz/cli-chat-golang/src/controllers"
)

type User struct {
	Username string
	Password string
}

func (u *User) Login(p *pgxpool.Pool) int {
	var user User
	var hp string

	hp = helpers.Sha256(u.Password)

	const query = `SELECT username, password FROM users 
	WHERE username=$1 AND password=$2`

	row := p.QueryRow(context.TODO(), query, u.Username, hp)
	if row != nil {
		err := row.Scan(&user.Username, &user.Password)

		if err != nil {
			return 404
		}
	} else {
		return 404
	}

	log.Println("username, password: ", user.Username, user.Password)

	return 200
}

func (u *User) Signup(p *pgxpool.Pool) int {
	var user User
	var hp string

	hp = helpers.Sha256(u.Password)

	controllers.CreateUser(user.Username, hp)

	return 200
}
