import * as arn from 'lib'

export function on(eventName, func) {
	arn.events.on(eventName, func)
}