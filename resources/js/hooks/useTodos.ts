import { useEffect, useState } from 'react';
import { dateString, useDate } from '../lib/date';

export interface Todo {
  id: number;
  completed: boolean;
  content: string;
  date: Date;
}

export default function useTodos(): Todo[] {
  const [todos, setTodos] = useState<Todo[]>([]);
  const { current } = useDate();

  useEffect(() => {
    fetch(`/todos/${dateString(current)}`, {
      credentials: 'include',
    })
      .then((r) => r.json())
      .then((r: Todo[]) => {
        setTodos(r);
      });
  }, [current]);

  return todos;
}
