import { JSX } from 'preact/jsx-runtime';
import { useDate } from '../lib/date';

export default function DateSelector(): JSX.Element {
  const { current, viewing, viewNext, viewPrev, setCurrent } = useDate();

  function thisWeek(): Date[] {
    const week: Date[] = [];

    for (let i = 1; i <= 7; i += 1) {
      const copy = new Date(viewing.valueOf());
      const first = viewing.getDate() - viewing.getDay() + i;
      week.push(new Date(copy.setDate(first)));
    }

    return week;
  }

  return (
    <div className="DateSelector">
      <ul className="DateSelector__dates">
        {thisWeek().map((day) => (
          <li>
            <button type="button" onClick={() => setCurrent(day)}>
              {day.toDateString()}
            </button>

            <p>{current.getTime() === day.getTime() ? 'Current!' : ''}</p>
          </li>
        ))}
      </ul>

      <button type="button" onClick={viewPrev}>
        Previous week
      </button>

      <button type="button" onClick={viewNext}>
        Next week
      </button>
    </div>
  );
}
