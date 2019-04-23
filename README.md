# Anime Notifier

[![Godoc reference][godoc-image]][godoc-url]
[![Go report card][goreportcard-image]][goreportcard-url]
[![Build status][travis-image]][travis-url]

## How to

### Prerequisites

* Install [Docker](https://www.docker.com/get-started) :whale:
* Install [Docker Compose](https://docs.docker.com/compose/install/) :whale:

### Installation

Download the source code:

```shell
git clone https://github.com/animenotifier/notify.moe.git && cd notify.moe
```

Download the developer tools:

```shell
docker pull animenotifier/notify.moe
```

Run the developer tools:

```shell
docker-compose run notify.moe
```

### Usage

Your home directory is mounted as `/my` inside Docker.

Usually you'd want to clone all repositories you use into a `projects` directory inside your home files. This directory can be accessed by both your favourite editor on the host machine and also inside Docker.

On your host:

* Fork the notify.moe repository on GitHub :new:
* Download the fork to your home directory :arrow_down:
* Enter the notify.moe directory: `cd notify.moe` :open_file_folder:
* Start the development tools `docker-compose run notify.moe` :whale:

Inside the docker container:

* Enter the notify.moe directory again `cd notify.moe` :open_file_folder:
* Compile TypeScript files using: `tsc` :shaved_ice:
* Start the web server using: `run` :pray:

The `run` binary is a file watcher that will restart the web server when it detects code changes.

### Networking

* Add `beta.notify.moe 127.0.0.1` to your `hosts` file :page_facing_up:
* If you're a Linux or Mac user, run `make ports` to forward ports :penguin:
* Otherwise, forward TCP port 443 to 4001 manually :thought_balloon:

### In your browser

* Open the settings, search for certificates :key:
* Import the file `security/default/root.crt` as a trusted Root authority :closed_lock_with_key:
* Open `https://beta.notify.moe` :house_with_garden:

## Find us

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
