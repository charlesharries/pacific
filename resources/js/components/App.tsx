import { JSX } from 'preact/compat';
import { DateProvider } from '../lib/date';
import { StatusProvider } from '../lib/status';
import DateSelector from './DateSelector';
import Editor from './Editor';

export function App(): JSX.Element {
  return (
    <DateProvider>
      <StatusProvider>
        <div className="App">
          <DateSelector />

          <Editor />
        </div>
      </StatusProvider>
    </DateProvider>
  );
}
