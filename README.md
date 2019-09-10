# notify.moe

[![Godoc][godoc-image]][godoc-url]
[![Report][report-image]][report-url]
[![Tests][tests-image]][tests-url]
[![Coverage][coverage-image]][coverage-url]
[![Sponsor][sponsor-image]][sponsor-url]

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

## Style

Please take a look at the [style guidelines](https://github.com/akyoto/quality/blob/master/STYLE.md) if you'd like to make a pull request.

## Sponsors

| [![Cedric Fung](https://avatars3.githubusercontent.com/u/2269238?s=70&v=4)](https://github.com/cedricfung) | [![Scott Rayapoullé](https://avatars3.githubusercontent.com/u/11772084?s=70&v=4)](https://github.com/soulcramer) | [![Eduard Urbach](https://avatars3.githubusercontent.com/u/438936?s=70&v=4)](https://eduardurbach.com) |
| --- | --- | --- |
| [Cedric Fung](https://github.com/cedricfung) | [Scott Rayapoullé](https://github.com/soulcramer) | [Eduard Urbach](https://eduardurbach.com) |

Want to see [your own name here?](https://github.com/users/akyoto/sponsorship)

[godoc-image]: https://godoc.org/github.com/animenotifier/notify.moe?status.svg
[godoc-url]: https://godoc.org/github.com/animenotifier/notify.moe
[report-image]: https://goreportcard.com/badge/github.com/animenotifier/notify.moe
[report-url]: https://goreportcard.com/report/github.com/animenotifier/notify.moe
[tests-image]: https://cloud.drone.io/api/badges/animenotifier/notify.moe/status.svg
[tests-url]: https://cloud.drone.io/animenotifier/notify.moe
[coverage-image]: https://codecov.io/gh/animenotifier/notify.moe/graph/badge.svg
[coverage-url]: https://codecov.io/gh/animenotifier/notify.moe
[sponsor-image]: https://img.shields.io/badge/github-donate-green.svg
[sponsor-url]: https://github.com/users/akyoto/sponsorship
