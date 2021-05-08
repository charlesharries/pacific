# Pacific
Because it's dark and you can throw all of your garbage into it.

## Prerequisites
**Node is required to run this project.** If you don't have Node installed already, I recommend [`nvm`](https://github.com/nvm-sh/nvm) for managing Node and its versions on your machine.

**Go is required to run this project.** You'll have to build the server from source, which requires Go on your machine. You can find installation instructions on [the Go website](https://golang.org/doc/install).

I'm also using `yarn` as my package manager, but you should be able to use `npm` just as easily, if you don't want to install another dependency.

## Installation
The application comes in 2 parts:
1. A Go server speaking to a sqlite database, exposing endpoints for user authentication and entry saving, and
2. A React application for actually running the frontend.

### Go server
You can install dependencies and build the Go server using the following command:

```sh
$ yarn build:server # or alternately: npm run build:server
```

This will install the server's dependencies and build the executable at `./bin/web`.

Once it's built, you can start up the server by running:

```sh
$ yarn start:server # or alternately: npm run start:server
```

### React application
The React application uses `esbuild` to build assets to the `public/` folder. Install dependencies and build JS with the following commands:

```sh
$ yarn install # or: npm install
$ yarn build # or: npm run build
```

