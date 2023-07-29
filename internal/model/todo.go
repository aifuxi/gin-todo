package model

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Done    bool
	Content string
}

func GetTodos() ([]Todo, error) {
	todos := []Todo{}

	result := db.Find(&todos)

	if result.Error != nil {
		return nil, result.Error
	}

	return todos, nil
}

func GetTodoById(id int) (Todo, error) {
	todo := Todo{}
	result := db.First(&todo, id)

	if result.Error != nil {

		return todo, result.Error
	}

	return todo, nil
}

func CreateTodo(todo *Todo) error {

	result := db.Create(&todo)

	if result.Error != nil {

		return result.Error
	}

	return nil
}

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

func UpdateTodoById(id int, todo *Todo) error {

	todoInDB := Todo{}
	result := db.First(&todo, id)

	if result.Error != nil {

		return result.Error
	}

	todoInDB.Content = todo.Content
	todoInDB.Done = todo.Done

	result = db.Save(&todo)

	if result.Error != nil {

		return result.Error
	}

	return nil
}
