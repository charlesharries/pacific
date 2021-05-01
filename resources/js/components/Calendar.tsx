import ReactDOM from 'react-dom';
import { humanMonth, useDate, isSameDay } from '../lib/date';

interface MonthProps {
  days: Date[];
}

interface DayProps {
  day: Date | null;
}

type CalendarDay = Date | null;

function Day({ day }: DayProps): JSX.Element {
  const { current, setCurrent } = useDate();

  if (!day) {
    return <td className="Calendar__day" />;
  }

  const classes = [
    'Calendar__day',
    'align-center',
    'font-xs',
    isSameDay(day, current) ? 'is-current' : '',
    isSameDay(day, new Date()) ? 'is-today' : '',
  ]
    .filter((d) => !!d)
    .join(' ');

  return (
    <td className={classes}>
      <button type="button" onClick={() => setCurrent(day)}>
        {day.getDate()}
      </button>
    </td>
  );
}

function Month({ days }: MonthProps): JSX.Element {
  const { current } = useDate();
  let currentWeek = 0;

  function weeks() {
    const ds: CalendarDay[][] = [[]];
    const firstDayOfWeek = 1;

    // Pre-fill first week
    for (let i = 0; i < days[0].getDay() - firstDayOfWeek; i += 1) {
      ds[currentWeek].push(null);
    }

    // Fill the rest of the month
    days.forEach((day) => {
      if (ds[currentWeek].length >= 7) {
        currentWeek += 1;
      }

      ds[currentWeek] = ds[currentWeek] || [];
      ds[currentWeek].push(day);
    });

    return ds;
  }

  function weekHasDay(week: CalendarDay[], day: Date) {
    return week.some((d) => d && isSameDay(d, day));
  }

  return (
    <div className="mt-sm">
      <h4 className="align-right">
        {humanMonth(days[0])} {days[0].getFullYear()}
      </h4>
      <table className="Calendar mt-xs">
        {weeks().map((week) => (
          <tr className={`Calendar__week ${weekHasDay(week, current) ? 'is-current' : ''}`}>
            {week.map((day) => (
              <Day day={day} />
            ))}
          </tr>
        ))}
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
