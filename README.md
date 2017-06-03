# Anime Notifier

Fetches your anime "watching" list and notifies you when a new anime episode is available. It also displays the time until a new episode is released.


Supports anime lists from:
- [anilist.co](https://anilist.co)
- [anime-planet.com](http://anime-planet.com)
- [myanimelist.net](http://myanimelist.net)
- [hummingbird.me](https://hummingbird.me)

Forum threads:
- [AniList](http://anilist.co/forum/thread/64)
- [HummingBird](https://forums.hummingbird.me/t/16787)
- [MyAnimeList](http://myanimelist.net/forum/?topicid=1175519)

Powered by:
- [Aero](https://github.com/aerojs/aero) | Web framework
- [Aerospike](https://github.com/aerospike) | Database

## Installation for developers

If you want to run this site locally on your own computer make sure to do the following:

* Get a Linux desktop or server (e.g. Ubuntu)
* Install node.js
* Install Aerospike
* Configure a namespace called `arn` in the Aerospike config, 4 GB space
* Clone the notify.moe repository
* Create a self-signed SSL certificate and put it inside a new `notify.moe/security/` directory
* You'll need to prepare a lot of API keys. Save them under `notify.moe/security/api-keys.json`
* Once database and API keys are setup, run `npm i -g pm2` to install pm2
* Install the dependencies by running `npm install`
* Start the website with `pm2 start ecosystem.json`
* Visit `https://localhost:5001/` in your browser

[![By Eduard Urbach](http://forthebadge.com/images/badges/built-with-love.svg)](https://github.com/blitzprog)
