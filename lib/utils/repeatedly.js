'use strict'

module.exports = (updateTimeInSeconds, func) => {
	func()
	setInterval(func, updateTimeInSeconds * 1000)
}