/* eslint-disable */
const esbuild = require('esbuild');
const { sassPlugin } = require('esbuild-sass-plugin');

esbuild
  .build({
    entryPoints: ['./resources/js/index.js'],
    bundle: true,
    sourcemap: true,
    inject: ['./resources/js/process-shim.js'],
    external: ['*.woff', '*.woff2'],
    define: {
      'process.env.NODE_ENV': '"development"',
      'process.env.NODE_DEBUG': false,
      global: 'window',
    },
    watch: {
      onRebuild(error, result) {
        if (error) console.error('build failed:', error);
        else console.log('\u0007');
      },
    },
    outfile: 'public/index.js',
    plugins: [sassPlugin()],
  })
  .then(() => console.log('built.'))
  .catch((e) => {
    console.error(e.message);
    process.exit(1);
  });
