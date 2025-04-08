package controllers

import (
	"fmt"
	"net/http"
	"time"
	"todo_list/database"
	"todo_list/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetTodos(c *gin.Context) {
	var todos []models.Todo

	result := database.DB.Find(&todos)
	if result.Error != nil {
		fmt.Println("DB Error:", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	fmt.Println("Todos fetched:", todos)
	c.JSON(http.StatusOK, todos)
}
func CreateTodo(c *gin.Context) {
	var todo models.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	todo.Id = uuid.New()
	todo.CreatedAt = time.Now()

	result := database.DB.Create(&todo)
	if result.Error != nil {
		fmt.Println(result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}

	c.JSON(http.StatusCreated, todo)
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	if err := database.DB.Where("id = ?", id).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		fmt.Println("Todo not found:", err)
		return
	}

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	result := database.DB.Save(&todo)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update todo"})
		fmt.Println("DB Error:", result.Error)
		return
	}
	c.JSON(http.StatusOK, todo)

}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	if err := database.DB.Where("id = ?", id).First(&todo).Error; err != nil {
		c.JSON((http.StatusNotFound), gin.H{"error": "Todo not found"})
		fmt.Println("Todo not found:", err)
		return
	}

	result := database.DB.Delete(&todo)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete todo"})
		fmt.Println("DB Error:", result.Error)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})

}
