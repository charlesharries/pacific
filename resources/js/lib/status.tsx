import { createContext, ReactChildren, useContext, useState } from 'react';

type Status = 'idle' | 'loading' | 'success' | 'error';

type StatusContext = {
  status: Status;
  setStatus: (s: Status) => void;
};

function useProvideStatus(): StatusContext {
  const [status, setStatus] = useState<Status>('idle');

  return { status, setStatus };
}

type StatusProviderProps = {
  children: JSX.Element;
};

const statusContext = createContext<StatusContext>(null!);

export function StatusProvider({ children }: StatusProviderProps): JSX.Element {
  const statusProvider = useProvideStatus();

  return <statusContext.Provider value={statusProvider}>{children}</statusContext.Provider>;
}

export const useStatus = (): StatusContext => useContext(statusContext);
