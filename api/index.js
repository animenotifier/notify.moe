'use strict';

let aero = require('aero');
let api = {};
api.users = require('./users');

api.install = function() {
	aero.get(/^api\/users\/(.*)/, function(request, response) {
		let userName = request.params[0];
		let user = api.users.get(userName);

		response.setHeader('Content-Type', 'application/json');
		response.end(JSON.stringify(user));
	});
};

module.exports = api;