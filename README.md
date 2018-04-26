# Installation

[![Godoc reference][godoc-image]][godoc-url]
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
* In your browser, import the file `security/default/rootCA.pem` as a trusted Root authority
* Open `https://beta.notify.moe`

[godoc-image]: https://godoc.org/github.com/animenotifier/notify.moe?status.svg
[godoc-url]: https://godoc.org/github.com/animenotifier/notify.moe
[travis-image]: https://travis-ci.org/animenotifier/notify.moe.svg?branch=go
[travis-url]: https://travis-ci.org/animenotifier/notify.moe
