export type CalendarDay = Date | null;

export default function useWeeks(days: Date[]): CalendarDay[][] {
  let currentWeek = 0;
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
