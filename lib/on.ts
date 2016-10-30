import * as arn from '.'

export function on(eventName, func) {
	arn.events.on(eventName, func)
}