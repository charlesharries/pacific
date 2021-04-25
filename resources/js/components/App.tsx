import { JSX } from 'preact/compat';
import { DateProvider } from '../lib/date';
import { StatusProvider } from '../lib/status';
import DateSelector from './DateSelector';
import Editor from './Editor';
import StatusIndicator from './StatusIndicator';

export function App(): JSX.Element {
  return (
    <DateProvider>
      <StatusProvider>
        <div className="App">
          <DateSelector />

          <StatusIndicator />

          <Editor />
        </div>
      </StatusProvider>
    </DateProvider>
  );
}
