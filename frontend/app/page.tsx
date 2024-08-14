'use client';

import { useState, useEffect } from 'react';
import axios from 'axios';
import TodoList from './components/TodoList';
import TodoForm from './components/TodoForm';
import { Todo } from '../types/todo';

export default function Home() {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [newTodo, setNewTodo] = useState<string>("");
  const [newDeadline, setNewDeadline] = useState<string>("");

  useEffect(() => {
    fetchTodos();
  }, []);

  const fetchTodos = async () => {
    try {
      const response = await axios.get('http://localhost:8080/todos');
      if (Array.isArray(response.data)) {
        setTodos(response.data);
      } else {
        console.error("Unexpected response format:", response.data);
      }
    } catch (error) {
      console.error("Error fetching todos:", error);
    }
  };

  const addTodo = async () => {
    if (!newTodo.trim()) return;
    try {
      await axios.post('http://localhost:8080/todos', { title: newTodo, deadline: newDeadline, status: "pending" });
      setNewTodo("");
      setNewDeadline("");
      fetchTodos();
    } catch (error) {
      console.error("Error adding todo:", error);
    }
  };

  const deleteTodo = async (id: number) => {
    try {
      await axios.delete(`http://localhost:8080/todos/${id}`);
      fetchTodos();
    } catch (error) {
      console.error("Error deleting todo:", error);
    }
  };

  const updateTodo = async (id: number, updatedTodo: Todo) => {
    try {
      await axios.put(`http://localhost:8080/todos/${id}`, updatedTodo);
      fetchTodos();
    } catch (error) {
      console.error("Error updating todo:", error);
    }
  };

  const sortTodosByDeadline = () => {
    const sortedTodos = [...todos].sort((a, b) => {
      if (a.deadline && b.deadline) {
        return new Date(a.deadline).getTime() - new Date(b.deadline).getTime();
      }
      return 0;
    });
    setTodos(sortedTodos);
  };

  return (
    <div className="container">
      <h1>TODO List</h1>
      <TodoForm
        newTodo={newTodo}
        newDeadline={newDeadline}
        onNewTodoChange={setNewTodo}
        onNewDeadlineChange={setNewDeadline}
        onAddTodo={addTodo}
      />
      <button onClick={sortTodosByDeadline}>Sort by Deadline</button>
      <TodoList
        todos={todos}
        onDelete={deleteTodo}
        onUpdate={updateTodo}
      />
    </div>
  );
}
