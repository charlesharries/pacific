import useTodos, { Todo } from '../hooks/useTodos';

function TodoItem({ id, completed, content, date }: Todo) {
  const name = `todo-${id}`;

  return (
    <li className="Todos__item">
      <label className="Todos__label ml-sm" htmlFor={name}>
        <input type="checkbox" id={name} name={name} />
        <input
          className="Todos__label__content"
          type="text"
          defaultValue={content}
          name={`${name}-content`}
        />
      </label>
    </li>
  );
}

export default function Todos(): JSX.Element {
  const todos = useTodos();

  console.log(todos);

  return (
    <ul className="Todos">
      {todos.map((t) => (
        <TodoItem {...t} />
      ))}
    </ul>
  );
}
