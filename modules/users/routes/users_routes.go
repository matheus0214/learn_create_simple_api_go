package routes

import "matheus0214/learn_gin/modules/users/controllers"

// UsersRoutes interface represent user routes
type UsersRoutes interface {
	Show()
	Create()
	Options()
}

// NewUserRoutes return methods to routes user
func NewUserRoutes() *controllers.UsersController {
	return &controllers.UsersController{}
}
