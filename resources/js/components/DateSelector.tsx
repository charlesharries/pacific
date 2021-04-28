import { useDate, humanDate, humanDayOfWeek } from '../lib/date';

export default function DateSelector(): JSX.Element {
  const { current, viewing, viewNext, viewPrev, setCurrent, viewToday } = useDate();

  /**
   * Get this week as an array of dates.
   *
   * We have to use (viewing.getDay() || 7) because Sunday (day 0)
   * will try to roll over onto next week.
   *
   * @returns {Date[]} This week
   */
  function thisWeek(): Date[] {
    const week: Date[] = [];

    for (let i = 1; i <= 7; i += 1) {
      const copy = new Date(viewing.valueOf());
      const first = viewing.getDate() - (viewing.getDay() || 7) + i;
      week.push(new Date(copy.setDate(first)));
    }

    return week;
  }

  /**
   * Is the given date the same as the current date?
   *
   * @param   {Date} date
   * @returns {boolean}
   */
  function isCurrent(day: Date): boolean {
    return day.getTime() === current.getTime();
  }

  /**
   * Is the given day in the past?
   *
   * @param   {Date} day
   * @returns {boolean}
   */
  function isPast(day: Date): boolean {
    const today = new Date();
    today.setHours(0);

    return day.getTime() <= today.getTime();
  }

  /**
   * Get the classnames for the given day.
   *
   * @param   {Date} date
   * @returns {string}
   */
  function classNames(day: Date): string {
    const cn = ['DateSelector__date'];

    if (isCurrent(day)) cn.push('is-current');
    if (isPast(day)) cn.push('is-past');

    return cn.filter((n) => !!n).join(' ');
  }

  return (
    <div className="DateSelector p-md py-sm">
      <p>
        <span className="font-sm">Today</span>
        <br />
        <span className="font-bold">{humanDate(current)}</span>
      </p>

      <ul className="DateSelector__dates m-0">
        {thisWeek().map((day) => (
          <li key={day.valueOf()} className={classNames(day)}>
            <button type="button" onClick={() => setCurrent(day)}>
              <span className="uppercase font-sm font-bold leading-loose">
                {humanDayOfWeek(day)}
              </span>
              <br />
              {day.getDate()}
            </button>
          </li>
        ))}
      </ul>

      <div className="DateSelector__navigate">
        <button type="button" className="button" onClick={viewPrev}>
          Prev
        </button>

        <button type="button" className="button" onClick={viewNext}>
          Next
        </button>
      </div>

      <div className="DateSelector__today">
        <button type="button" className="button" onClick={viewToday}>
          Today
        </button>
      </div>
    </div>
  );
}
