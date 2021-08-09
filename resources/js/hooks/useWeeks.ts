export type CalendarDay = Date | null;

export default function useWeeks(days: Date[]): CalendarDay[][] {
  let currentWeek = 0;
  const ds: CalendarDay[][] = [[]];
  const firstDayOfWeek = 1;

  // If the first day of the week is 0 (Sunday), set it to 7 so that
  // we can arbitrarily set the first day of the week.
  const firstDayOfMonth = days[0].getDay() || 7

  // Pre-fill first week
  for (let i = 0; i < firstDayOfMonth - firstDayOfWeek; i += 1) {
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
