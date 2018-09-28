# GO + Vue.JS Project Template

This is a template for working with a Vue.js frontend and a GO backend.  It simplifies the development process by providing code recompiling and execution for the GO backend as well as HOT code reloading for the Vue.js frontend.

## Dependencies

#### Compilers / Runtimes

```
brew install node
brew install golang
```

#### Project Dependencies

```
npm install
go get
```

## Development Server

The development server runs __webpack-dev-server__ with HOT reloading turned on.  All requests that are not part of the javascript hot updates are proxied through to the GO backend.  The GO backend is restarted whenever a change happens to its source code.

```
make dev
```

## Deploying

When it comes time to deploy, the _Makefile_ will compile the javascript for production, build the go backend binary, and package them into a _.zip_ file which will be versioned with the current date + time.

```
make dist
```