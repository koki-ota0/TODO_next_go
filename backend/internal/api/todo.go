package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo-server/internal/models"
	"todo-server/internal/services"

	"github.com/gorilla/mux"
)

// CreateTodoHandler handles the creation of a new Todo item
func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Call service to create the todo
	if err := services.CreateTodo(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

// GetTodosHandler handles retrieving all Todo items
func GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := services.GetTodos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// UpdateTodoHandler handles updating an existing Todo item
func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedTodo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&updatedTodo); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Call service to update the todo
	todo, err := services.UpdateTodo(id, &updatedTodo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

// DeleteTodoHandler handles deleting a Todo item
func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Call service to delete the todo
	if err := services.DeleteTodo(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
