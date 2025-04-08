package routes

import (
	"todo_list/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	todos := r.Group("/todos")

	{
		todos.GET("/", controllers.GetTodos)
		todos.POST("/", controllers.CreateTodo)
		todos.PUT("/:id", controllers.UpdateTodo)
		todos.DELETE("/:id", controllers.DeleteTodo)
	}
}
