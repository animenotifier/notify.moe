var aero = require("aero");
var api = require("./api/api");

// Website
aero.start();

// API
aero.app.use("/api", api);