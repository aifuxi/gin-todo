package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/aifuxi/gin-todo/internal/model"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "TEST",
	})
}

func GetTodos(c *gin.Context) {
	todos, err := model.GetTodos()

	if err != nil {
		log.Fatalln(err)
	}

	c.HTML(http.StatusOK, "todo_content.html", gin.H{
		"todos": todos,
	})
}

func CreateTodo(c *gin.Context) {
	content := c.PostForm("content")
	todo := model.Todo{Content: content, Done: false}

	err := model.CreateTodo(&todo)

	if err != nil {
		log.Fatalln(err)
	}

	todos, err2 := model.GetTodos()

	if err2 != nil {
		log.Fatalln(err2)
	}

	c.HTML(http.StatusOK, "todo_content.html", gin.H{
		"todos": todos,
	})
}

func DeleteTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	_, err := model.DeleteTodoById(id)

	if err != nil {
		log.Fatalln(err)
	}

	c.String(http.StatusOK, "")
}
