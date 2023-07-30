package controller

import (
	"fmt"
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

	c.HTML(http.StatusOK, "todo_content.html", gin.H{
		"todos": []model.Todo{todo},
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

func GetEditTodoById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := model.GetTodoById(id)

	fmt.Println(todo)
	if err != nil {
		log.Fatalln(err)
	}

	c.HTML(http.StatusOK, "todo_update_modal.html", gin.H{
		"todo": todo,
	})
}

func CancelEditTodo(c *gin.Context) {

	c.String(http.StatusOK, "")
}

func UpdateTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	todo, err := model.GetTodoById(id)
	content := c.PostForm("content")
	fmt.Printf("content: %v\n", content)
	todo.Content = content

	if err != nil {
		log.Fatalln(err)
	}

	err = model.UpdateTodoById(id, &todo)

	if err != nil {
		log.Fatalln(err)
	}

	c.HTML(http.StatusOK, "todo_content.html", gin.H{
		"todos": []model.Todo{todo},
	})
}

func DoneTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := model.GetTodoById(id)

	if err != nil {
		log.Fatalln(err)
	}

	todInDB, err2 := model.DoneTodoById(id, !todo.Done)

	if err2 != nil {
		log.Fatalln(err2)
	}

	c.HTML(http.StatusOK, "todo_content.html", gin.H{
		"todos": []model.Todo{todInDB},
	})
}
