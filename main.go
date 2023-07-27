package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Done    sql.NullBool `gorm:"default:false"`
	Content string
}

var db *gorm.DB
var err error

func main() {

	db, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("连接数据库失败", err)
	}

	// 迁移 schema，相当于让GORM帮你创建表
	db.AutoMigrate(&Todo{})

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Gin TODO List",
		})
	})

	r.POST("/create_todo", func(c *gin.Context) {
		content := c.PostForm("content")

		todo := Todo{Content: content}

		result := db.Create(&todo)

		if result.Error != nil {
			log.Fatal("创建todo失败：", result.Error)
		}

		str := fmt.Sprintf(`
    <tr>
      <td>%v</td>
      <td>%v</td>
      <td>%v</td>
      <td>
      <button type="button" hx-delete="/todo/%v" hx-target="table>tbody"
    hx-swap="innerHTML">删除</button>
<button type="button">编辑</button>
      </td>
    </tr>
  `, todo.ID, todo.Content, todo.CreatedAt.Format("2006-01-02 15:04:05"), todo.ID)

		c.String(http.StatusOK, str)
	})

	// 根据id删除TODO
	r.DELETE("/todo/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		todo := Todo{}
		_ = db.First(&todo, id)

		_ = db.Delete(&todo, id)

		str := ""
		todos := []Todo{}

		db.Find(&todos)

		for _, todo := range todos {
			str += fmt.Sprintf(`
      <tr>
        <td>%v</td>
        <td>%v</td>
        <td>%v</td>
        <td>
        <button type="button" hx-delete="/todo/%v" hx-target="table>tbody"
    hx-swap="innerHTML">删除</button>
<button type="button">编辑</button>
        </td>
      </tr>
    `, todo.ID, todo.Content, todo.CreatedAt.Format("2006-01-02 15:04:05"), todo.ID)
		}

		c.String(http.StatusOK, str)
	})

	r.GET("/get_todos", func(c *gin.Context) {
		str := ""
		todos := []Todo{}

		db.Find(&todos)

		for _, todo := range todos {
			str += fmt.Sprintf(`
      <tr>
        <td>%v</td>
        <td>%v</td>
        <td>%v</td>
        <td>
        <button type="button" hx-delete="/todo/%v" hx-target="table>tbody"
    hx-swap="innerHTML">删除</button>
<button type="button">编辑</button>
        </td>
      </tr>
    `, todo.ID, todo.Content, todo.CreatedAt.Format("2006-01-02 15:04:05"), todo.ID)
		}

		c.String(http.StatusOK, str)
	})

	// 获取所有TODO
	r.GET("/todos", GetTodos)

	// 根据id获取TODO
	r.GET("/todos/:id", GetTodoById)

	// 新增TODO
	r.POST("/todos", CreateTodo)

	// 根据id删除TODO
	r.DELETE("/todos/:id", DeleteTodoById)

	// 根据id更新TODO
	r.PUT("/todos/:id", UpdateTodoById)

	r.Run("localhost:9000")
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "air installed",
	})
}

func GetTodos(c *gin.Context) {
	todos := []Todo{}

	result := db.Find(&todos)

	if result.Error != nil {
		log.Fatal("查询数据库失败", result.Error)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": todos,
	})
}

func GetTodoById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "非法的id",
		})

		return
	}

	todo := Todo{}
	result := db.First(&todo, id)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": nil,
			"msg":  fmt.Sprintf("未找到id为%v的todo", id),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": todo,
	})
}

func CreateTodo(c *gin.Context) {
	todo := Todo{}

	err := c.ShouldBindJSON(&todo)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": nil,
			"msg":  fmt.Sprintf("解析请求出错: %v", err.Error()),
		})
	}

	result := db.Create(&todo)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": nil,
			"msg":  fmt.Sprintf("创建todo失败: %v", err.Error()),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": todo,
	})
}

func DeleteTodoById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "非法的id",
		})

		return
	}

	todo := Todo{}
	result := db.First(&todo, id)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": nil,
			"msg":  fmt.Sprintf("未找到id为%v的todo", id),
		})
		return
	}

	result = db.Delete(&todo, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": nil,
			"msg":  "删除todo失败",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": todo,
	})
}

func UpdateTodoById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "非法的id",
		})

		return
	}

	todo := Todo{}
	result := db.First(&todo, id)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": nil,
			"msg":  fmt.Sprintf("未找到id为%v的todo", id),
		})
		return
	}

	updateTodo := Todo{}
	err = c.ShouldBindJSON(&updateTodo)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": nil,
			"msg":  fmt.Sprintf("解析请求出错: %v", err.Error()),
		})
		return
	}

	todo.Content = updateTodo.Content
	todo.Done = updateTodo.Done

	result = db.Save(&todo)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": nil,
			"msg":  "更新todo失败",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": todo,
	})
}
