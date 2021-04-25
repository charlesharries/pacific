/* eslint-disable react-hooks/exhaustive-deps */
import { useRef, useEffect, DependencyList, EffectCallback } from 'react';

export default function useSubsequentEffect(fn: EffectCallback, inputs: DependencyList): void {
  const mounted = useRef(false);

  useEffect(() => {
    if (mounted.current) {
      fn();
    } else {
      mounted.current = true;
    }
  }, inputs);
}
