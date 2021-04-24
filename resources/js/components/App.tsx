import { JSX } from 'preact/compat';
import { DateProvider } from '../lib/date';
import DateSelector from './DateSelector';
import Editor from './Editor';

export function App(): JSX.Element {
  return (
    <DateProvider>
      <div className="App">
        <DateSelector />

        <Editor />
      </div>
    </DateProvider>
  );
}
