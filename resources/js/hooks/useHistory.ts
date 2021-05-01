import { useCallback, useEffect } from 'react';
import { dateString, useDate, dateFromUrl } from '../lib/date';

export default function useHistory(): void {
  const { current, setCurrent, setViewing } = useDate();

  const handlePopState = useCallback(() => {
    setCurrent(dateFromUrl());
    setViewing(dateFromUrl());
  }, [setCurrent, setViewing]);

  useEffect(() => {
    window.addEventListener('popstate', handlePopState);

    return () => {
      window.removeEventListener('popstate', handlePopState);
    };
  }, [handlePopState]);

  useEffect(() => {
    const date = dateString(current);
    const url = new URL(window.location.toString());

    // If we've just hit the back button, current will be the same as
    // the URL for some reason--so don't push that change back into
    // window.history.
    if (`/${date}` === url.pathname) {
      return;
    }

    url.pathname = date;
    window.history.pushState({ url: date }, '', url.toString());
  }, [current]);
}
