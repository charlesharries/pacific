import { useEffect } from 'react';
import { dateString, useDate } from '../lib/date';

export default function useHistory(): void {
  const { current } = useDate();

  useEffect(() => {
    const date = dateString(current);
    const url = new URL(window.location.toString());

    url.pathname = date;

    window.history.pushState({}, '', url.toString());
  }, [current]);
}
