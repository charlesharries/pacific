import { createContext, useContext, useState } from 'react';

export type StateUpdater<T> = React.Dispatch<React.SetStateAction<T>>;

type DateContext = {
  current: Date;
  viewing: Date;
  viewNext: () => void;
  viewPrev: () => void;
  viewToday: () => void;
  setCurrent: StateUpdater<Date>;
};

/**
 * Turn the current URL into a date. If no date is in the URL, just
 * return the current date.
 *
 * @returns {Date}
 */
export function dateFromUrl(): Date {
  const url = new URL(window.location.toString());
  const date = url.pathname;

  if (Number.isNaN(Date.parse(date))) {
    return new Date();
  }

  return new Date(date);
}

type DateProviderProps = {
  children: JSX.Element;
};

function useProvideDate(): DateContext {
  const [current, setCurrent] = useState<Date>(dateFromUrl());
  const [viewing, setViewing] = useState<Date>(dateFromUrl());

  function viewNext() {
    const to = new Date(viewing.valueOf());
    to.setDate(to.getDate() + 7);
    setViewing(to);
  }

  function viewPrev() {
    const to = new Date(viewing.valueOf());
    to.setDate(to.getDate() - 7);
    setViewing(to);
  }

  function viewToday() {
    const today = new Date();
    today.setHours(0, 0, 0, 0);

    setCurrent(today);
    setViewing(today);
  }

  return { current, viewing, viewNext, viewPrev, setCurrent, viewToday };
}

export function dateString(date: Date): string {
  const year = date.getFullYear();
  const month = date.getMonth() + 1;
  const day = date.getDate();
  const prefix = (num: number): string => (num < 10 ? `0${num}` : `${num}`);

  return `${year}-${prefix(month)}-${prefix(day)}`;
}

export function humanDate(date: Date): string {
  const months = [
    'January',
    'February',
    'March',
    'April',
    'May',
    'June',
    'July',
    'August',
    'September',
    'October',
    'November',
    'December',
  ];
  const year = date.getFullYear();
  const month = date.getMonth();
  const day = date.getDate();

  return `${day} ${months[month]} ${year}`;
}

export function humanDayOfWeek(date: Date): string {
  const days = ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'];

  return days[date.getDay()];
}

export function isSameDay(day1: Date, day2: Date): boolean {
  return (
    day1.getFullYear() === day2.getFullYear() &&
    day1.getMonth() === day2.getMonth() &&
    day1.getDate() === day2.getDate()
  );
}

const dateContext = createContext<DateContext>(null!);

export function DateProvider({ children }: DateProviderProps): JSX.Element {
  const dateProvider = useProvideDate();

  return <dateContext.Provider value={dateProvider}>{children}</dateContext.Provider>;
}

export const useDate = (): DateContext => useContext(dateContext);
