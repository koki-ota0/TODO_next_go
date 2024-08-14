import { ChangeEvent, useState } from 'react';
import { Todo } from '../../types/todo';

interface TodoItemProps {
  todo: Todo;
  onDelete: (id: number) => void;
  onUpdate: (id: number, updatedTodo: Todo) => void;
}

const TodoItem: React.FC<TodoItemProps> = ({ todo, onDelete, onUpdate }) => {
  const [isEditing, setIsEditing] = useState(false);
  const [title, setTitle] = useState(todo.title);
  const [deadline, setDeadline] = useState(todo.deadline || '');

  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    setTitle(e.target.value);
  };

  const handleDeadlineChange = (e: ChangeEvent<HTMLInputElement>) => {
    setDeadline(e.target.value);
  };

  const handleSave = () => {
    onUpdate(todo.id, { ...todo, title, deadline });
    setIsEditing(false);
  };

  return (
    <li className="todo-item">
      {isEditing ? (
        <>
          <input type="text" value={title} onChange={handleChange} />
          <input type="date" value={deadline} onChange={handleDeadlineChange} />
          <button className="edit" onClick={handleSave}>Save</button>
        </>
      ) : (
        <>
          <span>{todo.title}</span>
          <span>{todo.deadline}</span>
          <button className="edit" onClick={() => setIsEditing(true)}>Edit</button>
          <button className="delete" onClick={() => onDelete(todo.id)}>Delete</button>
        </>
      )}
    </li>
  );
};

export default TodoItem;
