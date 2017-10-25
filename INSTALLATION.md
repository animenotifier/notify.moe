# Anime Notifier

## Installation

### Prerequisites

* Install a Debian based operating system
* Install [Go](https://golang.org/dl/) (1.9 or higher)
* Install [TypeScript](https://www.typescriptlang.org/) (2.5 or higher)

### Download the repository and its dependencies

* `go get github.com/animenotifier/notify.moe`

### Build all

* Run `make tools` to install [pack](https://github.com/aerogo/pack) & [run](https://github.com/aerogo/run)
* Run `make all`
* Run `make ports` to set up local port forwarding *(80 to 4000, 443 to 4001)*

### Hosts

* Add the following lines to `/etc/hosts`:

```
45.32.159.101 arn-db
127.0.0.1     beta.notify.moe
```

### HTTPS

* [Create the certificate](https://stackoverflow.com/questions/10175812/how-to-create-a-self-signed-certificate-with-openssl) `notify.moe/security/fullchain.pem` (domain: `beta.notify.moe`)
* Create the private key `notify.moe/security/privkey.pem` (make sure it's decoded)

### API keys

* Get a Google OAuth 2.0 client key & secret from [console.developers.google.com](https://console.developers.google.com)
* Create the file `notify.moe/security/api-keys.json`:

```json
{
	"google": {
		"id": "YOUR_KEY",
		"secret": "YOUR_SECRET"
	}
}
```

### Run

* Start the web server in notify.moe directory: `run`
* Open `https://beta.notify.moe` which should now resolve to localhost