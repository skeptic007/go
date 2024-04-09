package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgresStore struct {
	db *sql.DB
}

type UserInterface interface {
	GetUser() (*User, error)
	GetUsers() (*User, error)
	CreateUser(*User) error
}

func NewPostgresStore() (*PostgresStore, error) {
	//connStr := "user=postgres dbname=postgres password=Password sslmode=disable"
	//db, err := sql.Open("postgres", connStr)
	//fmt.Println("hell db")

	db, err := sql.Open("postgres", "postgres://gqvcemiw:fXTZChMfu9tHic7BGdcye6NuFsKVyGHz@tiny.db.elephantsql.com/gqvcemiw?sslmode=disable")

	fmt.Println(db, err)
	//fmt.Println(db, err)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	return s.createAccountTable()
}

func (s *PostgresStore) createAccountTable() error {
	// query := `create table if not exists user (
	// 	id serial primary key,
	// 	name varchar(100),
	// 	email varchar(255),
	// 	dob int
	// )`

	query := `CREATE TABLE IF NOT EXISTS cars (
		brand VARCHAR(255),
		model VARCHAR(255),
		year INT
	  );`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreateUser(user *User) error {
	query := `insert into user (name, email, dob) values ($1, $2, $3)`
	_, err := s.db.Query(
		query,
		user.Name,
		user.Email,
		user.Dob)

	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) GetUser(id int) (*User, error) {
	query := `select * from user where id = $1`

	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanIntoUser(rows)
	}
	return nil, nil
}

func (s *PostgresStore) GetUsers() ([]*User, error) {

	userData := []*User{}
	rows, err := s.db.Query("Select * from user")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		users, err := scanIntoUser(rows)
		if err != nil {
			return nil, err
		}
		userData = append(userData, users)
	}
	return userData, nil

}

func scanIntoUser(rows *sql.Rows) (*User, error) {
	userDetail := new(User)
	err := rows.Scan(
		&userDetail.ID,
		&userDetail.Name,
		&userDetail.Email,
		&userDetail.Dob)

	return userDetail, err

}
