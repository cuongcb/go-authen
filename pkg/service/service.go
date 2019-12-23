package service

import (
	"github.com/cuongcb/go-authen/pkg/service/internal/dao/sql"
)

type serviceContext struct {
	repo repository
}

var ctx serviceContext

type repository interface {
	Save(string, string) error
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
func CreateUser(email, password string) error {
	return ctx.repo.Save(email, password)
}
