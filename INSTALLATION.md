# Installation

## Youtube guide

[![notify.moe Source Code Guide](https://i1.ytimg.com/vi/c6e-F51e_8w/maxresdefault.jpg)](https://www.youtube.com/watch?v=c6e-F51e_8w&amp=&t=3m42s)

## Prerequisites

* Install [Ubuntu](https://www.ubuntu.com/) or [MacOS](https://en.wikipedia.org/wiki/MacOS)
* Install [Go](https://golang.org/dl/) (1.9 or higher)
* Install [TypeScript](https://www.typescriptlang.org/) (2.6 or higher)
* Optional: [Git LFS](https://git-lfs.github.com/) if you want the video files

## Download the repository

* `go get github.com/animenotifier/notify.moe/...`

## Build all

* `cd $GOPATH/src/github.com/animenotifier/notify.moe`
* `make all`

## Browser

* `make ports` to set up local port forwarding *(80 to 4000, 443 to 4001)*
* `make browser` to start Google Chrome

## Database

* `git clone https://github.com/animenotifier/database ~/.aero/db/arn`

## Hosts

* Add `127.0.0.1 beta.notify.moe` to `/etc/hosts`

## Run

* Start the web server in notify.moe directory: `run`
* Open `https://beta.notify.moe` which should now resolve to localhost