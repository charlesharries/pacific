import { useEffect, useMemo, useState } from 'preact/hooks';
import { JSX } from 'preact/jsx-runtime';
import { createEditor, BaseEditor, Descendant } from 'slate';
import { Slate, Editable, withReact, ReactEditor } from 'slate-react';
import { useStatus } from '../lib/status';
import useDebounce from '../lib/useDebounce';

type CustomText = { text: string };
type CustomElement = { type: 'paragraph'; children: CustomText[] };

declare module 'slate' {
  interface CustomTypes {
    Editor: BaseEditor & ReactEditor;
    Element: CustomElement;
    Text: CustomText;
  }
}

export default function Editor(): JSX.Element {
  const { status, setStatus } = useStatus();
  const editor = useMemo(() => withReact(createEditor()), []);
  const [value, setValue] = useState<Descendant[]>([
    {
      type: 'paragraph',
      children: [{ text: 'A line of text in a paragraph' }],
    },
  ]);
  const debouncedValue = useDebounce<Descendant[]>(value, 1000);

  useEffect(() => {
    if (debouncedValue) {
      setStatus('loading');

      // Save the debouncedValue to the database

      setTimeout(() => {
        setStatus('success');
      }, 1000);
    }
  }, [debouncedValue, setStatus]);

  return (
    <>
      <p>Status: {status}</p>
      <Slate editor={editor} value={value} onChange={(val) => setValue(val)}>
        <Editable />
      </Slate>
    </>
  );
}
