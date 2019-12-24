package service

import (
	"github.com/cuongcb/go-authen/pkg/dtos"
	"github.com/cuongcb/go-authen/pkg/service/internal/dao/sql"
	"github.com/cuongcb/go-authen/pkg/service/internal/model"
)

type serviceContext struct {
	repo repository
}

var ctx serviceContext

type repository interface {
	Get(uint64) (*model.User, error)
	GetAll() ([]*model.User, error)
	Save(*model.User) (*model.User, error)
}

// Init ...
func Init() {
	repo, err := sql.New()
	if err != nil {
		panic(err)
	}

	ctx = serviceContext{
		repo: repo,
	}
}

// CreateUser ...
func CreateUser(email, password string) (*dtos.User, error) {
	user, err := ctx.repo.Save(&model.User{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return nil, err
	}

	return &dtos.User{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}

// GetUserList ...
func GetUserList() ([]*dtos.User, error) {
	users, err := ctx.repo.GetAll()
	if err != nil {
		return nil, err
	}

	userList := make([]*dtos.User, 0, len(users))
	for _, u := range users {
		userList = append(userList, &dtos.User{
			ID:       u.ID,
			Email:    u.Email,
			Password: u.Password})
		// userList[i] = &dtos.User{
		// 	ID:       users[i].ID,
		// 	Email:    users[i].Email,
		// 	Password: users[i].Password,
		// }
	}

	return userList, nil
}

// GetUser ...
func GetUser(id uint64) (*dtos.User, error) {
	user, err := ctx.repo.Get(id)
	if err != nil {
		return nil, err
	}

	return &dtos.User{
		ID:       id,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}
