package services

import (
	"matheus0214/learn_gin/modules/users/models"
	"matheus0214/learn_gin/modules/users/repository"
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
