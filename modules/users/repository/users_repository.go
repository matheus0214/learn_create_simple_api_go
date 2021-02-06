package repository

import "matheus0214/learn_gin/modules/users/models"

var usersRepository []*models.UserModel

// UsersRepository interface to represent repository user
type UsersRepository interface {
	Save(u *models.UserModel) *models.UserModel
	FindByEmail(email string) (*models.UserModel, error)
}

// UsersRepositoryImpl implementation users reposiotry
type UsersRepositoryImpl struct{}

// Save create a new instance from user
func (ur *UsersRepositoryImpl) Save(u *models.UserModel) *models.UserModel {
	usersRepository = append(usersRepository, u)

	return u
}

// FindByEmail search user by pass email like param
func (ur *UsersRepositoryImpl) FindByEmail(email string) (*models.UserModel, error) {
	var searchUser *models.UserModel

	for _, val := range usersRepository {
		if val.Email == email {
			searchUser = val
		}
	}

	if searchUser == nil {
		return nil, models.ErrUserDoesNotFound
	}

	return searchUser, nil
}

// NewUsersRepository instance from users repository
func NewUsersRepository() *UsersRepositoryImpl {
	return &UsersRepositoryImpl{}
}
