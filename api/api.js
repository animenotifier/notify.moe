"use strict";

var express = require("aero").express;
var app = express();

app.get("/", function(request, response) {
	response.send("Anime Release Notifier API");
});

app.get("/users/:name", function(request, response) {
	var userName = request.params.name;
	
	response.json({
		userName: userName
	});
});

module.exports = app;