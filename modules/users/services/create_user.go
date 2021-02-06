package services

import "github.com/matheus0214/projects/learn_gin/modules/users/models"

var usersRepository []*models.UserModel

// CreateUser create user service
func CreateUser(u *models.UserModel) (*models.UserModel, error) {
	existUser := false
	// Check user exist
	for _, user := range usersRepository {
		if user.Email == u.Email {
			existUser = true
			break
		}
	}

	if existUser {
		return nil, models.ErrInvalidFieldEmailErrorUser
	}

	usersRepository = append(usersRepository, u)

	return u, nil
}
