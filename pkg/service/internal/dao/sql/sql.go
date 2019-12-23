package sql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type sqlDao struct {
	db *sql.DB
}

const dataSourceName = "root:root@tcp(127.0.0.1:3306)/db"

// New ...
func New() (*sqlDao, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	return &sqlDao{db: db}, nil
}

func (s *sqlDao) Save(email, password string) error {
	query := "INSERT INTO user VALUES(?, ?)"
	_, err := s.db.Exec(query, email, password)
	return err
}
