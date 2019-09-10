# {name}

{go:header}

## Prerequisites

* Install [Linux](https://en.wikipedia.org/wiki/Linux), [MacOS](https://en.wikipedia.org/wiki/MacOS) or [WSL](https://en.wikipedia.org/wiki/Windows_Subsystem_for_Linux)
* Install [Go](https://golang.org/)
* Install [TypeScript](https://www.typescriptlang.org/)

## Installation

```shell
git clone https://github.com/animenotifier/notify.moe.git
cd notify.moe
go mod download
make tools
make assets
make server
make db
./notify.moe
```

## Networking

* Add `beta.notify.moe 127.0.0.1` to your `hosts` file
* Run `make ports`

## In your browser

* Open the settings, search for certificates
* Import the file `security/default/root.crt` as a trusted Root authority
* Open `https://beta.notify.moe`

## What now?

* Try the [example task for newcomers](docs/new-contributor-task.md).
* Make some changes and upload them to a new branch on your fork.
* Create a pull request on this repository.

## Find us

* [Discord](https://discord.gg/0kimAmMCeXGXuzNF)
* [Facebook](https://www.facebook.com/animenotifier)
* [Twitter](https://twitter.com/animenotifier)
* [GitHub](https://github.com/animenotifier/notify.moe)

## Contributing

Please read [CONTRIBUTING.md](docs/contributing.md) for details on how to contribute to this project.

## Statistics

![Uptime (30 days)](https://img.shields.io/uptimerobot/ratio/m777678498-177cb6b2990056768877bc2a.svg)
![Mozilla Observatory](https://img.shields.io/mozilla-observatory/grade/notify.moe.svg?publish)
![Chrome Extension](https://img.shields.io/chrome-web-store/users/hajchfikckiofgilinkpifobdbiajfch.svg?label=chrome%20users)
![Firefox Extension](https://img.shields.io/amo/users/anime-notifier.svg?label=firefox%20users)

{go:footer}
