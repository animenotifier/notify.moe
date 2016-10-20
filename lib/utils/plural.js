"use strict";
function plural(count, singular) {
    return (count === 1 || count === -1) ? (count + ' ' + singular) : (count + ' ' + singular + 's');
}
exports.plural = plural;
