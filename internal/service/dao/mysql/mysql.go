package mysql

import (
	"database/sql"
	"fmt"
	"time"

	// standar mysql
	_ "github.com/go-sql-driver/mysql"

	"github.com/cuongcb/go-authen/internal/service/model"
)

// Dao ...
type Dao struct {
	db *sql.DB
}

const dataSourceName = "root:root@tcp(cont_mysql)/appdb?parseTime=true"

// New ...
func New() (*Dao, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	for db.Ping() != nil {
		fmt.Println("cannot ping to db")
		time.Sleep(1 * time.Second)
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

	hasRows := rows.Next()
	if !hasRows {
		if err := rows.Err(); err != nil {
			fmt.Println(err)
			return nil, err
		}

		return nil, sql.ErrNoRows
	}

	var users []*model.User
	for hasRows {
		u := &model.User{}
		if err := rows.Scan(&u.ID,
			&u.CreatedAt,
			&u.UpdatedAt,
			&u.Email,
			&u.Password); err != nil {
			fmt.Println(err)
			return nil, err
		}

		users = append(users, u)

		hasRows = rows.Next()
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
