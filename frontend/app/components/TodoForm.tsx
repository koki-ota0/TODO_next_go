import { ChangeEvent } from 'react';

interface TodoFormProps {
  newTodo: string;
  newDeadline: string;
  onNewTodoChange: (value: string) => void;
  onNewDeadlineChange: (value: string) => void;
  onAddTodo: () => void;
}

const TodoForm: React.FC<TodoFormProps> = ({ newTodo, newDeadline, onNewTodoChange, onNewDeadlineChange, onAddTodo }) => {
  return (
    <div>
      <input
        type="text"
        value={newTodo}
        onChange={(e: ChangeEvent<HTMLInputElement>) => onNewTodoChange(e.target.value)}
        placeholder="New Todo"
      />
      <input
        type="date"
        value={newDeadline}
        onChange={(e: ChangeEvent<HTMLInputElement>) => onNewDeadlineChange(e.target.value)}
      />
      <button onClick={onAddTodo}>Add Todo</button>
    </div>
  );
};

export default TodoForm;
