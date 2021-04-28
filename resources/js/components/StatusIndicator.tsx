import { useStatus } from '../lib/status';
import Check from './icons/Check';
import Batsu from './icons/Batsu';
import Loading from './icons/Loading';

export default function StatusIndicator(): JSX.Element {
  const { status } = useStatus();

  switch (status) {
    case 'success':
      return <Check />;
    case 'loading':
      return <Loading />;
    case 'error':
      return <Batsu />;
    default:
      return <p>?</p>;
  }
}
