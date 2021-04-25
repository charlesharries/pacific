import { useRef, useEffect } from 'preact/hooks';

export default function useSubsequentEffect(fn: () => void, inputs: string[]): void {
  const mounted = useRef(false);

  useEffect(() => {
    if (mounted.current) {
      fn();
    } else {
      mounted.current = true;
    }
  }, inputs);
}
