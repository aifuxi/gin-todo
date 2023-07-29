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
	r.POST("/create_todo", controller.CreateTodo)
	r.DELETE("/delete_todo/:id", controller.DeleteTodo)

	// g := r.Group("/todo")
	// {
	// 	g.GET("/:id", controller.GetTodoById)
	// 	g.PUT("/:id", controller.UpdateTodoById)
	// 	g.DELETE("/:id", controller.DeleteTodoById)
	// 	g.GET("/", controller.GetTodo)
	// 	g.POST("/", controller.CreateTodo)
	// }

	return r
}
