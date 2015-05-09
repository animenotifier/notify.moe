"use strict";

let aero = require("aero");
let api = require("./api/api");

// Website
aero.start();

// API
aero.app.use("/api", api);