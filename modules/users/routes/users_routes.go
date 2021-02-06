package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheus0214/projects/learn_gin/modules/users/models"
	"github.com/matheus0214/projects/learn_gin/modules/users/services"
)

// UsersRoutes interface represent user routes
type UsersRoutes interface {
	Show()
	Create()
	Options()
}

type usersRoutesImpl struct{}

func (ur *usersRoutesImpl) Show(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"user": "Matheus",
	})
}

func (ur *usersRoutesImpl) Create(c *gin.Context) {
	datas, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Erro to get datas",
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
		c.JSON(http.StatusOK, gin.H{
			"error": models.ErrInvalidFieldEmailErrorUser.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, newUser)
	return
}

func (ur *usersRoutesImpl) Options(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"METHODS ALLOWED": "GET, POST",
		"Authorization":   "Bearer Token",
	})
}

// NewUserRoutes return methods to routes user
func NewUserRoutes() *usersRoutesImpl {
	return &usersRoutesImpl{}
}
