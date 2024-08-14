import TodoItem from './TodoItem';
import { Todo } from '../../types/todo';

interface TodoListProps {
  todos: Todo[];
  onDelete: (id: number) => void;
  onUpdate: (id: number, updatedTodo: Todo) => void;
}

const TodoList: React.FC<TodoListProps> = ({ todos, onDelete, onUpdate }) => {
  return (
    <ul>
      {todos.map(todo => (
        <TodoItem
          key={todo.id}
          todo={todo}
          onDelete={onDelete}
          onUpdate={onUpdate}
        />
      ))}
    </ul>
  );
};

export default TodoList;
