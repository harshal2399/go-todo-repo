package main

import (
	"todo_list/database"
	"todo_list/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()
	routes.RegisterRoutes(r)
	r.Run(":8080")

}
