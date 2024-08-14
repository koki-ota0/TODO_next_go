package services

import (
	"errors"
	"todo-server/internal/db"
	"todo-server/internal/models"
)

// CreateTodo は新しい Todo を作成し、データベースに保存します
func CreateTodo(todo *models.Todo) error {
	_, err := db.GetDB().Model(todo).Insert()
	return err
}

// GetTodos はすべての Todo をデータベースから取得します
func GetTodos() ([]models.Todo, error) {
	var todos []models.Todo
	err := db.GetDB().Model(&todos).Select()
	return todos, err
}

// UpdateTodo は指定された ID の Todo を更新します
func UpdateTodo(id int, updatedTodo *models.Todo) (*models.Todo, error) {
	todo := &models.Todo{ID: id}
	err := db.GetDB().Model(todo).WherePK().Select()
	if err != nil {
		return nil, errors.New("Todo not found")
	}

	todo.Title = updatedTodo.Title
	todo.Status = updatedTodo.Status

	_, err = db.GetDB().Model(todo).WherePK().Update()
	if err != nil {
		return nil, err
	}

	return todo, nil
}

// DeleteTodo は指定された ID の Todo を削除します
func DeleteTodo(id int) error {
	todo := &models.Todo{ID: id}
	_, err := db.GetDB().Model(todo).WherePK().Delete()
	if err != nil {
		return errors.New("Todo not found")
	}
	return nil
}
