"use strict";
const events_1 = require("events");
const getLikeController_1 = require("./getLikeController");
const getUnlikeController_1 = require("./getUnlikeController");
const fixListProviderUserName_1 = require("./fixListProviderUserName");
const isActiveUser_1 = require("./isActiveUser");
const aerospike = require('aero-aerospike');
class AnimeReleaseNotifier {
    constructor() {
        this.fixListProviderUserName = fixListProviderUserName_1.fixListProviderUserName;
        this.getLikeController = getLikeController_1.getLikeController;
        this.getUnlikeController = getUnlikeController_1.getUnlikeController;
        this.isActiveUser = isActiveUser_1.isActiveUser;
        this.events = new events_1.EventEmitter();
        this.api = require('../security/api-keys.json');
        this.db = aerospike.client(require('../config.json').database);
        this.db.connect().then(() => console.log('Successfully connected to database!'));
        this.listProviders = {
            AniList: require('./services/AniList.js')
        };
        this.animeProviders = {
            Nyaa: require('./services/Nyaa.js')
        };
    }
    auth(req, res, role) {
        if (!req.user) {
            res.end('Not logged in!');
            return false;
        }
        if (req.user.role !== 'admin' && req.user.role !== role) {
            res.end('Not authorized to view this page!');
            return false;
        }
        return true;
    }
    getLocation(user) {
        let locationAPI = `http://api.ipinfodb.com/v3/ip-city/?key=${this.api.ipInfoDB.clientID}&ip=${user.ip}&format=json`;
        return request(locationAPI).then(JSON.parse);
    }
    changeNick(user, newNick) {
        const userNameTakenMessage = 'Username is already taken.';
        let oldNick = user.nick;
        if (oldNick === newNick)
            return Promise.resolve();
        return this.db.get('NickToUser', newNick).then(record => {
            return Promise.reject(userNameTakenMessage);
        }).catch(error => {
            if (error === userNameTakenMessage)
                return;
            user.nick = newNick;
            return Promise.all([
                this.db.remove('NickToUser', oldNick),
                this.db.set('NickToUser', newNick, { userId: user.id }),
                this.db.set('Users', user.id, user)
            ]);
        });
    }
    fixGenre(genre) {
        return genre.replace(/ /g, '').replace(/-/g, '').toLowerCase();
    }
}
exports.AnimeReleaseNotifier = AnimeReleaseNotifier;
module.exports = new AnimeReleaseNotifier();
