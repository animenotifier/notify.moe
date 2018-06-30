# Installation

[![Godoc reference][godoc-image]][godoc-url]
[![Go report card][goreportcard-image]][goreportcard-url]
[![Build status][travis-image]][travis-url]

## Prerequisites

* Install [Ubuntu](https://www.ubuntu.com/) or [MacOS](https://en.wikipedia.org/wiki/MacOS)
* Install [Go](https://golang.org/dl/)
* Install [TypeScript](https://www.typescriptlang.org/)
* Install [Git LFS](https://git-lfs.github.com/)

## Start the installation

```bash
curl -s https://raw.githubusercontent.com/animenotifier/notify.moe/go/install.sh | bash
```

## Run the server

* Start the web server in notify.moe directory using: `run`
* In your browser, import the file `security/default/root.crt` as a trusted Root authority
* Open `https://beta.notify.moe`

## OS restarts

* If you restart your operating system, run `make ports` to update your port bindings

## Author

| [![Eduard Urbach on Twitter](https://gravatar.com/avatar/16ed4d41a5f244d1b10de1b791657989?s=70)](http://twitter.com/eduardurbach "Follow @eduardurbach on Twitter") |
|---|
| [Eduard Urbach](https://eduardurbach.com) |

[godoc-image]: https://godoc.org/github.com/animenotifier/notify.moe?status.svg
[godoc-url]: https://godoc.org/github.com/animenotifier/notify.moe
[goreportcard-image]: https://goreportcard.com/badge/github.com/animenotifier/notify.moe
[goreportcard-url]: https://goreportcard.com/report/github.com/animenotifier/notify.moe
[travis-image]: https://travis-ci.org/animenotifier/notify.moe.svg?branch=go
[travis-url]: https://travis-ci.org/animenotifier/notify.moe
