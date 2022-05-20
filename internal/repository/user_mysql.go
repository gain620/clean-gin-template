package repository

import (
	"database/sql"
)

type userMysqlRepo struct {
	conn *sql.DB
}

func New(conn *sql.DB) *userMysqlRepo {
	return &userMysqlRepo{
		conn: conn,
	}
}

// TODO : Add default pagination logic
func (u *userMysqlRepo) GetUsers() error {
	//query := `SELECT id, name, age, updated_at, created_at
	//					FROM users ORDER BY created_at LIMIT ? `
	//
	//
	return nil
}
