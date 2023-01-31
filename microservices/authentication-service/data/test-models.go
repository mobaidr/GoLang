package data

import (
	"database/sql"
	"time"
)

type PostgresTestRepository struct {
	Conn *sql.DB
}

func NewPostgresTestRepository(db *sql.DB) *PostgresTestRepository {
	return &PostgresTestRepository{
		Conn: db,
	}
}

func (p PostgresTestRepository) GetAll() ([]*User, error) {
	users := []*User{}

	return users, nil
}

func (p PostgresTestRepository) GetByEmail(email string) (*User, error) {
	user := User{
		ID:        1,
		Email:     "user@user.com",
		FirstName: "FirstName",
		LastName:  "LastName",
		Password:  "",
		Active:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &user, nil
}

func (p PostgresTestRepository) GetOne(id int) (*User, error) {
	user := User{
		ID:        id,
		Email:     "user@user.com",
		FirstName: "FirstName",
		LastName:  "LastName",
		Password:  "",
		Active:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &user, nil
}

func (p PostgresTestRepository) Update(user User) error {
	return nil
}

func (p PostgresTestRepository) DeleteByID(id int) error {
	return nil
}

func (p PostgresTestRepository) Insert(user User) (int, error) {
	return 1, nil
}

func (p PostgresTestRepository) ResetPassword(password string, user User) error {
	return nil
}

func (p PostgresTestRepository) PasswordMatches(plainText string, user User) (bool, error) {
	return true, nil
}
