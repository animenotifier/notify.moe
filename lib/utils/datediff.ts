export function inSeconds(a: number, b: number): number {
	return b - a
}

export function inMinutes(a: number, b: number): number {
	let val = (b - a) / 60.0

	if(Math.abs(val) < 1)
		return 0
	else
		return Math.ceil(val)
}

export function inHours(a: number, b: number): number {
	let val = (b - a) / (60 * 60.0)

	if(Math.abs(val) < 1)
		return 0
	else
		return Math.ceil(val)
}

export function inDays(a: number, b: number): number {
	let val = (b - a) / (24 * 60 * 60.0)

	if(Math.abs(val) < 1)
		return 0
	else
		return Math.ceil(val)
}