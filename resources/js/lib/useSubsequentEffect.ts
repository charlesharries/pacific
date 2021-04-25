/* eslint-disable react-hooks/exhaustive-deps */
import { useRef, useEffect, Inputs, EffectCallback } from 'preact/hooks';

export default function useSubsequentEffect(fn: EffectCallback, inputs: Inputs): void {
  const mounted = useRef(false);

  useEffect(() => {
    if (mounted.current) {
      fn();
    } else {
      mounted.current = true;
    }
  }, inputs);
}
