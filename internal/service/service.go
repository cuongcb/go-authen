package service

import (
	"time"

	"github.com/cuongcb/go-authen/internal/dtos"
	"github.com/cuongcb/go-authen/internal/service/cache"
	"github.com/cuongcb/go-authen/internal/service/dao/mysql"
	"github.com/cuongcb/go-authen/internal/service/model"
)

type serviceContext struct {
	repo  reposer
	cache cacher
}

var ctx serviceContext

type reposer interface {
	Get(uint64) (*model.User, error)
	GetByMail(string) (*model.User, error)
	GetAll() ([]*model.User, error)
	Save(*model.User) (*model.User, error)
}

type cacher interface {
	Set(string, string, time.Duration) error
	Get(string) (string, error)
}

// Init ...
func Init() {
	repo, err := mysql.New()
	if err != nil {
		panic(err)
	}

	cache, err := cache.New()
	if err != nil {
		panic(err)
	}

	ctx = serviceContext{
		repo:  repo,
		cache: cache,
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

// GetUserByMail ...
func GetUserByMail(email string) (*dtos.User, error) {
	user, err := ctx.repo.GetByMail(email)
	if err != nil {
		return nil, err
	}

	return &dtos.User{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}
