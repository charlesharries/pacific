import 'preact/debug';
import ReactDOM, { createElement } from 'preact/compat';
import { App } from './components/App';

const $root = document.getElementById('root');

if ($root) {
  ReactDOM.render(createElement(App, {}), $root);
}
