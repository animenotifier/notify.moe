"use strict";
const arn = require("./lib");
let app = require('aero')();
global.app = app;
global.arn = arn;
global.HTTP = require('http-status-codes');
app.on('database ready', db => {
    global.db = db;
});
app.use(require('body-parser').json());
app.run();
