"use strict";
var arn = require('./lib');
var app = require('aero')();
global.app = app;
global.arn = arn;
global.HTTP = require('http-status-codes');
app.on('database ready', function (db) {
    global.db = db;
});
// For POST requests
app.use(require('body-parser').json());
// Start the server
app.run();
