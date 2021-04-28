import { dateString, useDate } from '../lib/date';
import { useStatus } from '../lib/status';
import useCsrf from './useCsrf';
import useDebounce from './useDebounce';
import useSubsequentEffect from './useSubsequentEffect';

type ApiResponse = {
  error: boolean;
  message: string;
};

/**
 * How many milliseconds after stopping typing should we wait before
 * we send the autosave request?
 */
const autosaveDelay: number = 750

/**
 * Watch a value, and when it changes, POST that value to the backend
 * for saving.
 *
 * @param   {string} debouncedValue Value to watch and autosave.
 * @returns {void}
 */
export default function useAutosave(value: string): void {
  const { setStatus } = useStatus();
  const { current } = useDate();
  const csrfToken = useCsrf();
  const debouncedValue = useDebounce<string>(value, autosaveDelay);

  useSubsequentEffect(() => {
    if (debouncedValue && debouncedValue.length) {
      setStatus('loading');

      try {
        (async () => {
          const formBody = new URLSearchParams();
          formBody.append('content', JSON.stringify(debouncedValue));
          formBody.append('csrf_token', csrfToken);

          // eslint-disable-next-line @typescript-eslint/no-unsafe-assignment
          const data: ApiResponse = await fetch(`/notes/${dateString(current)}`, {
            method: 'POST',
            headers: {
              'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8',
            },
            credentials: 'include',
            body: formBody,
          }).then((r) => r.json());

          if (data.error) {
            throw new Error('api error');
          }
          setStatus('success');
        })();
      } catch (err) {
        console.error(err);
        setStatus('error');
      }
    }
  }, [debouncedValue, setStatus]);
}
