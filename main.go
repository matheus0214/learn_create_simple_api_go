package main

import (
	"matheus0214/learn_gin/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.Routes(r)

	r.Run(":3000")
}
