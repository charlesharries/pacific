import { useEffect } from 'react';
import { useStatus } from '../lib/status';
import { dateString, StateUpdater, useDate } from '../lib/date';

const emptyEditor: string = '';

type ApiNote = {
  id: number;
  content: string;
  date: string;
};

export default function useLoadEditor(setValue: StateUpdater<string>): void {
  const { setStatus } = useStatus();
  const { current } = useDate();

  useEffect(() => {
    setStatus('loading');

    setValue(emptyEditor);

    fetch(`/notes/${dateString(current)}`, {
      credentials: 'include',
    })
      .then((r) => r.json())
      .then((r: ApiNote) => {
        if (r.content) {
          setValue(JSON.parse(r.content) as string);
        }

        setStatus('success');
      });
  }, [current, setStatus, setValue]);
}
