import { DateProvider } from '../lib/date';
import { StatusProvider } from '../lib/status';
import DateSelector from './DateSelector';
import '../../sass/app.scss';
import Calendar from './Calendar';
import Content from './Content';

export function App(): JSX.Element {
  return (
    <DateProvider>
      <StatusProvider>
        <div className="App">
          <Calendar />

          <DateSelector />

          <Content />
        </div>
      </StatusProvider>
    </DateProvider>
  );
}
