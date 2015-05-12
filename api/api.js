"use strict";

let express = require("aero").express;
let app = express();
let riak = require("nodiak").getClient("http", "104.236.134.105", 8098);

riak.ping(function(err, response) {
	console.log(response);
	
	riak.bucket("Accounts").objects.all(function(err, r_objs) {
	    console.log("Fetched.");
	});
});

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