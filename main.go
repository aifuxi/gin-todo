package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateTodoReq struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Done    bool   `json:"done"`
}

type UpdateTodoReq struct {
	Content string `json:"content"`
	Done    bool   `json:"done"`
}

type Todo struct {
	CreateTodoReq
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var todos = []Todo{}

func main() {
	r := gin.Default()

	r.GET("/ping", Ping)

	// 获取所有TODO
	r.GET("/todo", GetTodo)

	// 根据id获取TODO
	r.GET("/todo/:id", GetTodoById)

	// 新增TODO
	r.POST("/todo", CreateTodo)

	// 根据id删除TODO
	r.DELETE("/todo/:id", DeleteTodoById)

	// 根据id更新TODO
	r.PUT("/todo/:id", UpdateTodoById)

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "air installed",
	})
}

func GetTodo(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": todos,
	})
}

func GetTodoById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{
			"msg": "非法的id",
		})

		return
	}

	todoIndex := -1

	for i, v := range todos {
		if v.Id == id {
			todoIndex = i
			break
		}
	}

	if todoIndex == -1 {
		c.JSON(400, gin.H{
			"msg": fmt.Sprintf("未找到id为%v的todo", id),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": todos[todoIndex],
	})
}

func CreateTodo(c *gin.Context) {
	req := CreateTodoReq{}

	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(500, gin.H{
			"data": nil,
			"msg":  "解析请求出错",
		})
	}

	todo := Todo{
		CreateTodoReq: req,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	todos = append(todos, todo)

	c.JSON(200, gin.H{
		"data": todo,
	})
}

func DeleteTodoById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{
			"msg": "非法的id",
		})

		return
	}

	deleteIndex := -1

	for i, v := range todos {
		if v.Id == id {
			deleteIndex = i
			break
		}
	}

	if deleteIndex == -1 {
		c.JSON(400, gin.H{
			"msg": fmt.Sprintf("未找到id为%v的todo", id),
		})
		return
	}
	// [a, b, c, d]  =>  如果要删除c这个元素，那么先找到c所在的索引，2
	// [:2] => 拿到 [a, b]; [(2+1);] => [3:]拿到[d]，然后 [d]展开，append到[1,2]这个切片
	// 最后把得到的值重新赋值给todos就删除了
	deletedTodo := todos[deleteIndex]
	todos = append(todos[:deleteIndex], todos[deleteIndex+1:]...)

	c.JSON(200, gin.H{
		"data": deletedTodo,
	})
}

func UpdateTodoById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{
			"msg": "非法的id",
		})

		return
	}

	updateIndex := -1

	for i, v := range todos {
		if v.Id == id {
			updateIndex = i
			break
		}
	}

	if updateIndex == -1 {
		c.JSON(400, gin.H{
			"msg": fmt.Sprintf("未找到id为%v的todo", id),
		})
		return
	}

	updateReq := UpdateTodoReq{}

	err = c.ShouldBindJSON(&updateReq)

	if err != nil {
		c.JSON(500, gin.H{
			"data": nil,
			"msg":  "解析请求出错",
		})
	}

	// 更新对应索引上的todo
	todos[updateIndex].Content = updateReq.Content
	todos[updateIndex].Done = updateReq.Done
	todos[updateIndex].UpdatedAt = time.Now()

	c.JSON(200, gin.H{
		"data": todos[updateIndex],
	})
}
