"use strict";
function inSeconds(a, b) {
    return b - a;
}
exports.inSeconds = inSeconds;
function inMinutes(a, b) {
    let val = (b - a) / 60.0;
    if (Math.abs(val) < 1)
        return 0;
    else
        return Math.ceil(val);
}
exports.inMinutes = inMinutes;
function inHours(a, b) {
    let val = (b - a) / (60 * 60.0);
    if (Math.abs(val) < 1)
        return 0;
    else
        return Math.ceil(val);
}
exports.inHours = inHours;
function inDays(a, b) {
    let val = (b - a) / (24 * 60 * 60.0);
    if (Math.abs(val) < 1)
        return 0;
    else
        return Math.ceil(val);
}
exports.inDays = inDays;
