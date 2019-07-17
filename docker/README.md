# Docker

## Prerequisites

* Install [Docker](https://www.docker.com/get-started)
* Install [Docker Compose](https://docs.docker.com/compose/install/)

## Installation

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

## Start the server

* Enter the notify.moe directory: `cd notify.moe`
* Download dependencies: `go mod download`
* Compile TypeScript files using: `tsc`
* Start the web server using: `run`

## Networking

* Add `beta.notify.moe 127.0.0.1` to your `hosts` file

## In your browser

* Open the settings, search for certificates
* Import the file `security/default/root.crt` as a trusted Root authority
* Open `https://beta.notify.moe`

## Tips

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