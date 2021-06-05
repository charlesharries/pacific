import ReactDOM from 'react-dom';
import useWeeks, { CalendarDay } from '../hooks/useWeeks';
import { humanMonth, useDate, isSameDay } from '../lib/date';

interface MonthProps {
  days: Date[];
}

interface DayProps {
  day: Date | null;
  isWeekend: boolean;
}

function Day({ day, isWeekend }: DayProps): JSX.Element {
  const { current, setCurrent, setViewing } = useDate();

  if (!day) {
    return <td className="Calendar__day" />;
  }

  function goTo(d: Date): void {
    setCurrent(d);
    setViewing(d);
  }

  const classes = [
    'Calendar__day',
    'align-center',
    'font-xs',
    isSameDay(day, current) ? 'is-current' : '',
    isSameDay(day, new Date()) ? 'is-today' : '',
    isWeekend ? 'is-weekend' : '',
  ]
    .filter((d) => !!d)
    .join(' ');

  return (
    <td className={classes}>
      <button type="button" onClick={() => goTo(day)}>
        {day.getDate()}
      </button>
    </td>
  );
}

function Month({ days }: MonthProps): JSX.Element {
  const { current, viewing } = useDate();
  const weeks = useWeeks(days);

  function weekHasDay(week: CalendarDay[], day: Date) {
    return week.some((d) => d && isSameDay(d, day));
  }

  return (
    <div className="mt-sm">
      <h5 className="align-right">
        {humanMonth(days[0])} {days[0].getFullYear()}
      </h5>
      <table className="Calendar mt-xs">
        <tbody>
          {weeks.map((week, i) => (
            <tr
              key={`${humanMonth(days[0])}-${i}`}
              className={`Calendar__week ${weekHasDay(week, viewing) ? 'is-current' : ''}`}
            >
              {week.map((day, di) => (
                <Day key={`${i}-${di}`} day={day} isWeekend={[5, 6].includes(di)} />
              ))}
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

function Calendar(): JSX.Element {
  const { current } = useDate();

  function offsetMonth(offset: number): Date {
    const d = new Date(current.getTime());
    d.setMonth(current.getMonth() + offset);

    return d;
  }

  function getMonth(date: Date): Date[] {
    const lastDayOfMonth = new Date(date.getFullYear(), date.getMonth() + 1, 0);
    const days = [];

    for (let i = 1; i <= lastDayOfMonth.getDate(); i += 1) {
      days.push(new Date(date.getFullYear(), date.getMonth(), i));
    }

    return days;
  }

  return (
    <div className="p-md">
      <Month days={getMonth(offsetMonth(-1))} />
      <Month days={getMonth(current)} />
      <Month days={getMonth(offsetMonth(1))} />
    </div>
  );
}

export default function CalendarWrapper(): JSX.Element {
  const calendarRoot = document.getElementById('calendar');

  return ReactDOM.createPortal(<Calendar />, calendarRoot as Element);
}
