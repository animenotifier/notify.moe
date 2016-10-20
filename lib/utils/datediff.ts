
export function inSeconds(a: Date, b: Date): number {
	return b.valueOf() - a.valueOf()
}

export function inMinutes(a: Date, b: Date): number {
	let val = (b.valueOf() - a.valueOf()) / 60.0

	if(Math.abs(val) < 1)
		return 0
	else
		return Math.ceil(val)
}

export function inHours(a: Date, b: Date): number {
	let val = (b.valueOf() - a.valueOf()) / (60 * 60.0)

	if(Math.abs(val) < 1)
		return 0
	else
		return Math.ceil(val)
}

export function inDays(a: Date, b: Date): number {
	let val = (b.valueOf() - a.valueOf()) / (24 * 60 * 60.0)

	if(Math.abs(val) < 1)
		return 0
	else
		return Math.ceil(val)
}