global.app = require('aero')();
global.arn = require('./lib');
global.HTTP = require('http-status-codes');
app.on('database ready', db => {
    global.db = db;
});
app.use(require('body-parser').json());
app.run();
