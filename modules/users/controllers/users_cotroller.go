package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"matheus0214/learn_gin/modules/users/models"
	"matheus0214/learn_gin/modules/users/services"

	"github.com/gin-gonic/gin"
)

// UsersController implemets routes to user
type UsersController struct{}

// Show route to get specif user
func (ur *UsersController) Show(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"user": "Matheus",
	})
}

// Create create one user
func (ur *UsersController) Create(c *gin.Context) {
	datas, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error to get datas",
		})
		return
	}

	newUser := &models.UserModel{}
	err = json.Unmarshal(datas, newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Erro to process user datas",
		})
		return
	}

	errorsUser := newUser.ValidateFieldsUser()
	if errorsUser != nil {
		c.JSON(http.StatusBadRequest, errorsUser)
		return
	}

	_, err = services.CreateUser(newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, newUser)
	return
}

// Options show options from user routes
func (ur *UsersController) Options(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"METHODS ALLOWED": "GET, POST",
		"Authorization":   "Bearer Token",
	})
}
