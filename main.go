package main

import (
	"github.com/gin-gonic/gin"
	"github.com/matheus0214/projects/learn_gin/routes"
)

func main() {
	r := gin.Default()

	routes.Routes(r)

	r.Run(":3000")
}
