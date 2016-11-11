"use strict";
const arn = require("arn");
global.arn = arn;
global.app = require('aero')();
app.on('database ready', db => {
    global.db = db;
});
app.use(require('body-parser').json());
app.run();
