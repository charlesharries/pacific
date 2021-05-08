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

Once it's built, duplicate the `.env.example` file and rename it `.env`. The default should work here, but make sure to change your `APP_SECRET` to some random string; the longer, the better.

Then, you can start up the server by running:

```sh
$ yarn start:server # or alternately: npm run start:server
```

This starts up a server at `http://localhost:9001`.

### React application
The React application uses `esbuild` to build assets to the `public/` folder. Install dependencies and build JS with the following commands:

```sh
$ yarn install # or: npm install
$ yarn build # or: npm run build
```

Because we're using `esbuild` here, the build command should go pretty fast, especially if you're used to Webpack.

You should now have the fully-build assets in your `public/` folder. Visiting `http://localhost:9001` should show the application in a logged-out state.

#### Development
You can start up the development server by running the following:

```sh
$ yarn dev # or: npm run dev
```

The dev bundle is basically the same as the production bundle, but with sourcemaps and watch mode enabled. See `build.js` and `build.production.js` for more details.

## Getting started
Once you've got your assets built and the server up and running, it's time to register. Visit `http://localhost:9001` (or click 'Register' in the sidebar). Input your details here--the password has to be 10 characters or longer, and the email needs to be, well, an email.

Once you've registered, you won't be able to log in--I've added an `active` field to the database to prevent arbitrary signups in production. You'll have to open your database (located at `database/db.sqlite` if you kept the defaults in `.env`) and change your new user's `active` field to `1` before you can log in.

Once you've done that, you should be able to input your username and password that you registered with, and it should log you in to the application!
