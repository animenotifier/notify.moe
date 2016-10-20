"use strict";
const events_1 = require("events");
const getLikeController_1 = require("./getLikeController");
exports.getLikeController = getLikeController_1.getLikeController;
const getUnlikeController_1 = require("./getUnlikeController");
exports.getUnlikeController = getUnlikeController_1.getUnlikeController;
const fixListProviderUserName_1 = require("./fixListProviderUserName");
exports.fixListProviderUserName = fixListProviderUserName_1.fixListProviderUserName;
const isActiveUser_1 = require("./isActiveUser");
exports.isActiveUser = isActiveUser_1.isActiveUser;
const getAnimeIdBySimilarTitle_1 = require("./getAnimeIdBySimilarTitle");
exports.getAnimeIdBySimilarTitle = getAnimeIdBySimilarTitle_1.getAnimeIdBySimilarTitle;
const getIdByTitle_1 = require("./getIdByTitle");
exports.getIdByTitle = getIdByTitle_1.getIdByTitle;
const aerospike = require('aero-aerospike');
const events = new events_1.EventEmitter();
exports.events = events;
const api = require('../security/api-keys.json');
exports.api = api;
const db = aerospike.client(require('../config.json').database);
exports.db = db;
const listProviders = {
    AniList: require('./services/AniList.js')
};
exports.listProviders = listProviders;
const animeProviders = {
    Nyaa: require('./services/Nyaa.js')
};
exports.animeProviders = animeProviders;
function auth(req, res, role) {
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
exports.auth = auth;
function getLocation(user) {
    let locationAPI = `http://api.ipinfodb.com/v3/ip-city/?key=${this.api.ipInfoDB.clientID}&ip=${user.ip}&format=json`;
    return request(locationAPI).then(JSON.parse);
}
exports.getLocation = getLocation;
function changeNick(user, newNick) {
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
exports.changeNick = changeNick;
function fixGenre(genre) {
    return genre.replace(/ /g, '').replace(/-/g, '').toLowerCase();
}
exports.fixGenre = fixGenre;
db.connect().then(() => console.log('Successfully connected to database!'));
