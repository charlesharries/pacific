import { JSX } from 'preact/compat';
import { DateProvider } from '../lib/date';
import DateSelector from './DateSelector';

export function App(): JSX.Element {
  return (
    <DateProvider>
      <div className="App">
        <DateSelector />
      </div>
    </DateProvider>
  );
}
