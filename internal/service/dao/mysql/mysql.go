package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/cuongcb/go-authen/internal/service/model"
)

// Dao ...
type Dao struct {
	db *sql.DB
}

const dataSourceName = "root:Bacuong304@@tcp(127.0.0.1:3306)/db?parseTime=true"

// New ...
func New() (*Dao, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	return &Dao{db: db}, nil
}

// Save ...
func (d *Dao) Save(u *model.User) (*model.User, error) {
	query := "INSERT INTO user VALUES(?, CURRENT_TIME(), CURRENT_TIME(), ?, ?)"
	result, err := d.db.Exec(query, nil, u.Email, u.Password)
	if err != nil {
		return nil, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:       uint64(lastID),
		Email:    u.Email,
		Password: u.Password,
	}, nil

}

// GetAll ...
func (d *Dao) GetAll() ([]*model.User, error) {
	query := "SELECT * FROM user"
	rows, err := d.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, err
		}

		return nil, sql.ErrNoRows
	}

	var users []*model.User
	for rows.Next() {
		u := &model.User{}
		if err := rows.Scan(&u.ID,
			&u.CreatedAt,
			&u.UpdatedAt,
			&u.Email,
			&u.Password); err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, rows.Err()
}

// Get ...
func (d *Dao) Get(id uint64) (*model.User, error) {
	query := "SELECT * FROM user WHERE id = ?"
	row := d.db.QueryRow(query, id)

	u := &model.User{}
	if err := row.Scan(&u.ID,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.Email,
		&u.Password); err != nil {
		return nil, err
	}

	return u, nil
}

// GetByMail ...
func (d *Dao) GetByMail(email string) (*model.User, error) {
	query := "SELECT * FROM user WHERE email = ?"
	row := d.db.QueryRow(query, email)

	u := &model.User{}
	if err := row.Scan(&u.ID,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.Email,
		&u.Password); err != nil {
		return nil, err
	}

	return u, nil
}
