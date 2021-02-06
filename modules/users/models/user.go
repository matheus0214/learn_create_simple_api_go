package models

import (
	"errors"

	"github.com/matheus0214/projects/learn_gin/shared"
)

// UserModel modeling user model
type UserModel struct {
	Name  string `json:"name"`
	Age   uint   `json:"age"`
	Email string `json:"email"`
}

// Errors validations user
var (
	ErrInvalidFieldNameErrorUser  = errors.New("Field name should not be empty")
	ErrInvalidFieldEmailErrorUser = errors.New("Field email should not be empty")
	ErrUserEmailAlreadyInUse      = errors.New("Email already in use")
	ErrUserDoesNotFound           = errors.New("User does not found")
)

// ValidateEmail validation to check user email is empty
func (u *UserModel) ValidateEmail() error {
	if u.Email == "" {
		return ErrInvalidFieldEmailErrorUser
	}

	return nil
}

// ValidateFieldsUser validate all fields in user
func (u *UserModel) ValidateFieldsUser() []shared.ErrorsModel {
	err := []shared.ErrorsModel{}

	if e := u.ValidateEmail(); e != nil {
		err = append(err, shared.ErrorsModel{Field: "email", Message: e.Error()})
	}

	if u.Name == "" {
		err = append(err, shared.ErrorsModel{Field: "name", Message: ErrInvalidFieldNameErrorUser.Error()})
	}

	if len(err) >= 1 {
		return err
	}

	return nil
}
