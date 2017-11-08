# Installation

## Prerequisites

* Install [Ubuntu](https://www.ubuntu.com/) or any of its derivates
* Install [Go](https://golang.org/dl/) (1.9 or higher)
* Install [TypeScript](https://www.typescriptlang.org/) (2.5 or higher)

## Download the repository and its dependencies

* `go get github.com/animenotifier/notify.moe`

## Build all

* Navigate to the project directory `notify.moe`
* Run `make tools` to install [pack](https://github.com/aerogo/pack) & [run](https://github.com/aerogo/run)
* Run `make ports` to set up local port forwarding *(80 to 4000, 443 to 4001)*
* Run `make all`

## Hosts

* Add `127.0.0.1 beta.notify.moe` to `/etc/hosts`

## Browser

* Start Chrome via `google-chrome --ignore-certificate-errors`

## Database

* `go get github.com/animenotifier/database`
* `ln -s $GOPATH/src/github.com/animenotifier/database ~/.aero/db/arn`

## Run

* Start the web server in notify.moe directory: `run`
* Open `https://beta.notify.moe` which should now resolve to localhost