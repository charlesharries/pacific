import { DateProvider } from '../lib/date';
import { StatusProvider } from '../lib/status';
import DateSelector from './DateSelector';
import Editor from './Editor';
import '../../sass/app.scss';
import Calendar from './Calendar';

export function App(): JSX.Element {
  return (
    <DateProvider>
      <StatusProvider>
        <div className="App">
          <Calendar />

          <DateSelector />

          <Editor />
        </div>
      </StatusProvider>
    </DateProvider>
  );
}
