import { JSX } from 'preact/jsx-runtime';
import { useStatus } from '../lib/status';

export default function StatusIndicator(): JSX.Element {
  const { status } = useStatus();

  return <p>{status}</p>;
}
