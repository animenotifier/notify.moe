'use strict'

module.exports = {
	inSeconds: function(a, b) {
        return parseInt(b - a)
    },

	inMinutes: function(a, b) {
		let val = (b - a) / 60.0

		if(Math.abs(val) < 1)
			return 0
		else
        	return Math.ceil(val)
    },

	inHours: function(a, b) {
        let val = (b - a) / (60 * 60.0)

		if(Math.abs(val) < 1)
			return 0
		else
        	return Math.ceil(val)
    },

    inDays: function(a, b) {
		let val = (b - a) / (24 * 60 * 60.0)

		if(Math.abs(val) < 1)
			return 0
		else
        	return Math.ceil(val)
    }
}