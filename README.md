# Anime Notifier

[![Godoc reference][godoc-image]][godoc-url]
[![Go report card][goreportcard-image]][goreportcard-url]
[![Build status][travis-image]][travis-url]

## How to

### Prerequisites

* Install [Docker](https://www.docker.com/get-started)
* Install [Docker Compose](https://docs.docker.com/compose/install/)

### Installation

Download the source code:

```shell
git clone https://github.com/animenotifier/notify.moe.git && cd notify.moe
```

Download the developer tools:

```shell
docker pull animenotifier/notify.moe
```

Start the developer tools:

```shell
docker-compose up -d
```

Attach to a terminal:

```shell
docker attach notify.moe
```

### Start the server

* Enter the notify.moe directory: `cd notify.moe`
* Download dependencies: `go mod download`
* Compile TypeScript files using: `tsc`
* Start the web server using: `run`

### Networking

* Add `beta.notify.moe 127.0.0.1` to your `hosts` file

### In your browser

* Open the settings, search for certificates
* Import the file `security/default/root.crt` as a trusted Root authority
* Open `https://beta.notify.moe`

### Tips

* You can detach from the terminal using `Ctrl P -> Ctrl Q`.
* If you need to shutdown everything, use `docker-compose down`.
* Your home directory is mounted as `/my` inside Docker.
* Fork the notify.moe repository and upload your changes to the fork.
* Clone all the repositories you use into a `projects` directory inside your home files.
* The `run` binary is a file watcher that will restart the web server when it detects code changes.
* File modification events [don't work](https://github.com/docker/for-win/issues/56) on Docker for Windows.
* Use an editor like [Visual Studio Code](http://code.visualstudio.com) to access the source code on the host.
* To automatically compile TypeScript files in VS Code, press `Ctrl Shift B` and select `tsc: watch`.
* Use a Linux system for maximum performance.

### Bookmark

Create a bookmark in your browser and set this code as the URL:

```js
javascript:(() => {
	location = location.href.indexOf('://beta.') === -1 ?
	location.href.replace('://', '://beta.') : location.href.replace('://beta.', '://');
})();
```

Clicking this bookmark will let you switch between `notify.moe` (live) and `beta.notify.moe` (development).

## Find us

* [Discord](https://discord.gg/0kimAmMCeXGXuzNF)
* [Facebook](https://www.facebook.com/animenotifier)
* [Twitter](https://twitter.com/animenotifier)
* [Docker](https://hub.docker.com/r/animenotifier/notify.moe)
* [GitHub](https://github.com/animenotifier/notify.moe)

## Contributing

Please read [CONTRIBUTING.md](https://github.com/animenotifier/notify.moe/blob/go/CONTRIBUTING.md) for details on how to contribute to this project.

## License

This project is licensed under the [MIT License](https://github.com/animenotifier/notify.moe/blob/go/LICENSE).

## Author

| [![Eduard Urbach on Twitter](https://gravatar.com/avatar/16ed4d41a5f244d1b10de1b791657989?s=70)](https://twitter.com/eduardurbach "Follow @eduardurbach on Twitter") |
|---|
| [Eduard Urbach](https://eduardurbach.com) |

[godoc-image]: https://godoc.org/github.com/animenotifier/notify.moe?status.svg
[godoc-url]: https://godoc.org/github.com/animenotifier/notify.moe
[goreportcard-image]: https://goreportcard.com/badge/github.com/animenotifier/notify.moe
[goreportcard-url]: https://goreportcard.com/report/github.com/animenotifier/notify.moe
[travis-image]: https://travis-ci.org/animenotifier/notify.moe.svg?branch=go
[travis-url]: https://travis-ci.org/animenotifier/notify.moe
