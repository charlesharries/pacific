import { useState } from 'react';
import RichEditor from 'rich-markdown-editor';
import { useStatus } from '../lib/status';
import useAutosave from '../hooks/useAutosave';
import useLoadEditor from '../hooks/useLoadEditor';
import theme from '../lib/editor-theme';

export default function Editor(): JSX.Element | null {
  const [value, setValue] = useState<string>('');
  const { status } = useStatus();

  useLoadEditor(setValue);
  useAutosave(value);

  if (status === 'loading' && !value) {
    return null;
  }

  return <RichEditor defaultValue={value} onChange={(v) => setValue(v)} theme={theme} />;
}
