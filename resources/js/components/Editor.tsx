import { useEffect, useMemo, useState } from 'preact/hooks';
import { JSX } from 'preact/jsx-runtime';
import { createEditor, BaseEditor, Descendant } from 'slate';
import { Slate, Editable, withReact, ReactEditor } from 'slate-react';
import { useDate } from '../lib/date';
import { useStatus } from '../lib/status';
import useDebounce from '../lib/useDebounce';
import useSubsequentEffect from '../lib/useSubsequentEffect';

type CustomText = { text: string };
type CustomElement = { type: 'paragraph'; children: CustomText[] };

declare module 'slate' {
  interface CustomTypes {
    Editor: BaseEditor & ReactEditor;
    Element: CustomElement;
    Text: CustomText;
  }
}

function dateString(date: Date): string {
  const year = date.getFullYear();
  const month = date.getMonth() + 1;
  const day = date.getDate();
  const prefix = (num: number): string => (num < 10 ? `0${num}` : `${num}`);

  return `${year}-${prefix(month)}-${prefix(day)}`;
}

type ApiResponse = {
  error: boolean;
  message: string;
};

type ApiNote = {
  id: number;
  content: string;
  date: string;
};

function csrfToken(): string {
  const $meta: HTMLMetaElement | null = document.querySelector('meta[name="csrf_token"]');

  return $meta ? $meta.content : '';
}

export default function Editor(): JSX.Element {
  const { status, setStatus } = useStatus();
  const { current } = useDate();
  const editor = useMemo(() => withReact(createEditor()), []);
  const [value, setValue] = useState<Descendant[]>([]);
  const debouncedValue = useDebounce<Descendant[]>(value, 1000);

  // Handle first load
  useEffect(() => {
    setStatus('loading');

    fetch(`/notes/${dateString(current)}`, {
      credentials: 'include',
    })
      .then((r) => r.json())
      .then((r: ApiNote) => {
        setValue(JSON.parse(r.content) as Descendant[]);
      });
  }, []);

  // Handle saves
  useSubsequentEffect(() => {
    if (debouncedValue) {
      setStatus('loading');

      // Save the debouncedValue to the database
      try {
        (async () => {
          const formBody = new URLSearchParams();
          formBody.append('content', JSON.stringify(debouncedValue));
          formBody.append('csrf_token', csrfToken());

          // eslint-disable-next-line @typescript-eslint/no-unsafe-assignment
          const data: ApiResponse = await fetch(`/notes/${dateString(current)}`, {
            method: 'POST',
            headers: {
              'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8',
            },
            credentials: 'include',
            body: formBody,
          }).then((r) => r.json());

          if (data.error) {
            throw new Error('api error');
          }
          setStatus('success');
        })();
      } catch (err) {
        console.error(err);
        setStatus('error');
      }
    }
  }, [current, debouncedValue, setStatus]);

  return (
    <>
      <p>Status: {status}</p>
      <Slate editor={editor} value={value} onChange={(val) => setValue(val)}>
        <Editable />
      </Slate>
    </>
  );
}
