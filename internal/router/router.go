package router

import (
	"github.com/aifuxi/gin-todo/internal/controller"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/", controller.Index)
	r.GET("/get_todos", controller.GetTodos)
	r.GET("/cancel_edit_todo", controller.CancelEditTodo)
	r.GET("/edit_todo/:id", controller.GetEditTodoById)
	r.POST("/create_todo", controller.CreateTodo)
	r.DELETE("/delete_todo/:id", controller.DeleteTodo)
	r.PUT("/done_todo/:id", controller.DoneTodo)
	r.PUT("/update_todo/:id", controller.UpdateTodo)

	return r
}
