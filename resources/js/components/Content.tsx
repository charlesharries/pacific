import Editor from './Editor';
import Todos from './Todos';

export default function Content(): JSX.Element {
  return (
    <section className="Content">
      <div className="Content__inner">
        <Todos />

        <Editor />
      </div>
    </section>
  );
}
