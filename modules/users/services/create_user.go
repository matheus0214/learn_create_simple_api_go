package services

import (
	"github.com/matheus0214/projects/learn_gin/modules/users/models"
	"github.com/matheus0214/projects/learn_gin/modules/users/repository"
)

// CreateUser create user service
func CreateUser(u *models.UserModel) (*models.UserModel, error) {
	usersRepo, _ := repository.NewUsersRepository().FindByEmail(u.Email)

	if usersRepo != nil {
		return nil, models.ErrUserEmailAlreadyInUse
	}

	newUser := repository.NewUsersRepository().Save(u)

	return newUser, nil
}
