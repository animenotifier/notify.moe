# Installation

## Youtube guide

[![notify.moe Source Code Guide](https://i1.ytimg.com/vi/c6e-F51e_8w/maxresdefault.jpg)](https://www.youtube.com/watch?v=c6e-F51e_8w&amp=&t=3m42s)

## Prerequisites

* Install [Ubuntu](https://www.ubuntu.com/) or [MacOS](https://en.wikipedia.org/wiki/MacOS)
* Install [Go](https://golang.org/dl/) (1.9 or higher)
* Install [TypeScript](https://www.typescriptlang.org/) (2.6 or higher)
* Install [Git LFS](https://git-lfs.github.com/)

## Confirm that prerequisites are installed

```bash
go version
tsc --version
git lfs version
```

## Start the installation

```bash
curl -s https://raw.githubusercontent.com/animenotifier/notify.moe/go/install.sh | sudo bash
```

## Run the server

* Start the web server in notify.moe directory: `run`
* Open `https://beta.notify.moe` which should now resolve to localhost
