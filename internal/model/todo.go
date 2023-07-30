package model

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Done    bool
	Content string
}

// 获取所有TODO
func GetTodos() ([]Todo, error) {
	todos := []Todo{}

	result := db.Find(&todos)

	if result.Error != nil {
		return nil, result.Error
	}

	return todos, nil
}

// 根据id获取TODO
func GetTodoById(id int) (Todo, error) {
	todo := Todo{}
	result := db.First(&todo, id)

	if result.Error != nil {

		return todo, result.Error
	}

	return todo, nil
}

// 创建TODO
func CreateTodo(todo *Todo) error {

	result := db.Create(&todo)

	if result.Error != nil {

		return result.Error
	}

	return nil
}

// 根据id删除TODO
func DeleteTodoById(id int) (Todo, error) {
	todo := Todo{}
	result := db.First(&todo, id)

	if result.Error != nil {

		return todo, result.Error
	}

	result = db.Delete(&todo, id)

	if result.Error != nil {

		return Todo{}, result.Error
	}

	return todo, nil
}

// 根据id更新TODO完成状态
func DoneTodoById(id int, done bool) (Todo, error) {
	todoInDB := Todo{}
	result := db.First(&todoInDB, id)

	if result.Error != nil {

		return todoInDB, result.Error
	}

	db.Model(&todoInDB).Update("Done", done)

	if result.Error != nil {

		return todoInDB, result.Error
	}

	return todoInDB, nil
}

// 根据id更新TODO
func UpdateTodoById(id int, todo *Todo) error {

	result := db.First(&Todo{}, id)

	if result.Error != nil {

		return result.Error
	}

	result = db.Model(todo).Updates(todo)

	if result.Error != nil {

		return result.Error
	}

	return nil
}
