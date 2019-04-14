# Installation

[![Godoc reference][godoc-image]][godoc-url]
[![Go report card][goreportcard-image]][goreportcard-url]
[![Build status][travis-image]][travis-url]

## Prerequisites

* Install [Docker](https://www.docker.com/get-started) and [Compose](https://docs.docker.com/compose/install/)

## Clone the repository

```bash
git clone https://github.com/animenotifier/notify.moe.git
```

## Download the dev image

```bash
docker pull animenotifier/notify.moe
```

## Run the dev image

```bash
docker-compose run notify.moe
```

## Run the server

* Compile TypeScript files using: `tsc`
* Start the web server in notify.moe directory using: `run`
* In your browser, import the file `security/default/root.crt` as a trusted Root authority
* Open `https://beta.notify.moe`

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
