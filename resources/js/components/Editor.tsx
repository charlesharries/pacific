import { useMemo, useState } from 'preact/hooks';
import { JSX } from 'preact/jsx-runtime';
import { createEditor, BaseEditor, Descendant } from 'slate';
import { Slate, Editable, withReact, ReactEditor } from 'slate-react';

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
  const editor = useMemo(() => withReact(createEditor()), []);
  const [value, setValue] = useState<Descendant[]>([
    {
      type: 'paragraph',
      children: [{ text: 'A line of text in a paragraph' }],
    },
  ]);

  function handleChange(newValue: Descendant[]): void {
    setValue(newValue);
  }

  return (
    <Slate editor={editor} value={value} onChange={handleChange}>
      <Editable />
    </Slate>
  );
}
