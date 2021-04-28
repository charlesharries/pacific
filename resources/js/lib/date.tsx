import { createContext, ReactChildren, useContext, useState } from 'react';

export type StateUpdater<T> = React.Dispatch<React.SetStateAction<T>>;

type DateContext = {
  current: Date;
  viewing: Date;
  viewNext: () => void;
  viewPrev: () => void;
  setCurrent: StateUpdater<Date>;
};

function useProvideDate(): DateContext {
  const [current, setCurrent] = useState<Date>(new Date());
  const [viewing, setViewing] = useState<Date>(new Date());

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

  return { current, viewing, viewNext, viewPrev, setCurrent };
}

export function dateString(date: Date): string {
  const year = date.getFullYear();
  const month = date.getMonth() + 1;
  const day = date.getDate();
  const prefix = (num: number): string => (num < 10 ? `0${num}` : `${num}`);

  return `${year}-${prefix(month)}-${prefix(day)}`;
}

type DateProviderProps = {
  children: JSX.Element;
};

const dateContext = createContext<DateContext>(null!);

export function DateProvider({ children }: DateProviderProps): JSX.Element {
  const dateProvider = useProvideDate();

  return <dateContext.Provider value={dateProvider}>{children}</dateContext.Provider>;
}

export const useDate = (): DateContext => useContext(dateContext);
