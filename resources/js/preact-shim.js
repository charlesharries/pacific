// We need this to inject into every module, so that we don't have to
// go about importing all sorts of stuff we're not actually using w/i
// our component files.
import React from 'preact/compat';

export { React };
