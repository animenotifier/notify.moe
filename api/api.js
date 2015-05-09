"use strict";

let express = require("aero").express;
let app = express();

app.get("/", function(request, response) {
	response.send("Anime Release Notifier API");
});

app.get("/users/:name", function(request, response) {
	let userName = request.params.name;
	
	response.json({
		userName: userName
	});
});

module.exports = app;