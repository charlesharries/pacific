const esbuild = require('esbuild')
const sassPlugin = require('esbuild-plugin-sass');

function onRebuild(err, result) {
  console.log({err, result});
}

esbuild.build({
  entryPoints: ['./resources/js/index.js'],
  bundle: true,
  sourcemap: true,
  inject: ['./resources/js/process-shim.js'],
  define: {
    'process.env.NODE_ENV': 'development',
    'process.env.NODE_DEBUG': false,
    'global': 'window',
  },
  watch: { onRebuild(error, result) {
    if (error) console.error('build failed:', error);
    else console.log('build succeeded:', result);
  } },
  outfile: 'public/index.js',
  plugins: [sassPlugin()],
}).catch(e => {
  console.error(e.message);
  process.exit(1);
})