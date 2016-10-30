import * as arn from 'arn'

export function on(eventName, func) {
	arn.events.on(eventName, func)
}