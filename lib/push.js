'use strict'

let apiKeys = require('../security/api-keys.json')
let gcm = require('node-gcm')

var sender = new gcm.Sender(apiKeys.gcm.serverKey);

/*var message = new gcm.Message({
    notification: {
        title: 'Hello, World',
        icon: 'https://notify.moe/images/characters/arn-waifu.png',
        body: 'This is a notification that will be displayed ASAP.'
    }
})

console.log('Sending test notification')
sender.sendNoRetry(message, {
	registrationTokens
}, function(err, response) {
	if(err) console.error(err);
	else    console.log(response);
})*/

module.exports = {

}