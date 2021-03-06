import { createElement } from 'react';
import ReactDOM from 'react-dom';
import { App } from './components/App';

const $root = document.getElementById('root');

if ($root) {
  ReactDOM.render(createElement(App, {}), $root);
}
