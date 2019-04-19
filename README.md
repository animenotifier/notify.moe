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

```bash
git clone https://github.com/animenotifier/notify.moe.git && cd notify.moe
```

Download the developer tools:

```bash
docker pull animenotifier/notify.moe
```

Run the developer tools:

```bash
docker-compose run notify.moe
```

### Usage

Your home directory is mounted as `/my` inside Docker.

Usually you'd want to clone all repositories you use into a `projects` directory. This directory can be accessed by both your favourite editor on the host machine and also inside Docker.

* Fork the notify.moe repository on GitHub
* Enter the notify.moe directory: `cd notify.moe`
* Compile TypeScript files using: `tsc`
* Compile template/style files using: `pack` (optional)
* Start the web server using: `run`

The `run` binary is a file watcher that will restart the web server when it detects code changes.

### Networking

* Add `beta.notify.moe 127.0.0.1` to your `hosts` file
* Forward TCP port 4001 to 443 (Linux / MacOS users can run `make ports`)

### In your browser

* Open the settings, search for certificates
* Import the file `security/default/root.crt` as a trusted Root authority
* Open `https://beta.notify.moe`

## Find Us

* [Discord](https://discord.gg/0kimAmMCeXGXuzNF)
* [Facebook](https://www.facebook.com/animenotifier)
* [Twitter](https://twitter.com/animenotifier)
* [Google+](https://plus.google.com/+AnimeReleaseNotifierOfficial)
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
