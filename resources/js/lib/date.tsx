import { ComponentChildren, createContext, JSX } from 'preact';
import { StateUpdater, useContext, useState } from 'preact/hooks';

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

type DateProviderProps = {
  children: ComponentChildren;
};

const dateContext = createContext<DateContext>(null!);

export function DateProvider({ children }: DateProviderProps): JSX.Element {
  const dateProvider = useProvideDate();

  return <dateContext.Provider value={dateProvider}>{children}</dateContext.Provider>;
}

export const useDate = (): DateContext => useContext(dateContext);
