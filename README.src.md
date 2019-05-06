# {name}

{go:header}

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

### What now?

* Try the [example task for newcomers](docs/new-contributor-task.md).
* Install Pug/Jade and Stylus extensions for your editor.
* If you're **not** using VS Code, map Pug/Jade to `.pixy` and Stylus to `.scarlet` files (they're similar).
* Make some changes and upload them to your fork.
* Create a pull request on this repository (with the diffs of your fork).

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

## Statistics

![Uptime (30 days)](https://img.shields.io/uptimerobot/ratio/m777678498-177cb6b2990056768877bc2a.svg)
![Mozilla Observatory](https://img.shields.io/mozilla-observatory/grade/notify.moe.svg?publish)
![Chrome Extension](https://img.shields.io/chrome-web-store/users/hajchfikckiofgilinkpifobdbiajfch.svg?label=chrome%20users)
![Firefox Extension](https://img.shields.io/amo/users/anime-notifier.svg?label=firefox%20users)

{go:footer}
