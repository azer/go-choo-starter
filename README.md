## go-choo-starter

Starter for [choo](https://github.com/yoshuawuyts/choo) projects with [Go](http://golang.org) backend. No configuration needed.

### How It Works?

* Renders [choo](https://github.com/yoshuawuyts/choo) views on server-side with [go-duktape](https://github.com/olebedev/go-duktape).
* Provides a simple Makefile to manage your project;
  * `make start` starts everything
  * `make stop` stops all
  * `make develop` starts the server and watches everything (Go, JS, CSS) for changes.
  * `make setup` installs all dependencies (Go and JS)

### Install

Clone the repo and install the dependencies:

```bash
git clone git@github.com:azer/go-choo-starter.git hello-world
cd hello-world
make setup # Runs `go get` and `npm install` for you.
```

### Usage

Here is how you start the server:

```bash
make develop
```

`develop` watches your code (Go, JS and CSS) and applies changes immediately. If you don't need that, you can run `make start` and `make stop` commands. You should use these two commands when you're not actively changing your code.
