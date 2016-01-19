'use strict'

class AutoUpdateCache {
	constructor(func, updateTimeInSeconds, defaultValue) {
		this.cache = defaultValue
		func().then(newValue => this.cache = newValue)

		setInterval(() => {
			func().then(newValue => this.cache = newValue)
		}, updateTimeInSeconds * 1000)
	}
}

module.exports = AutoUpdateCache