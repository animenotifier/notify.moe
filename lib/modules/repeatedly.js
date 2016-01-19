'use strict'

arn.repeatedly = (updateTimeInSeconds, func) => {
	func()
	setInterval(func, updateTimeInSeconds * 1000)
}