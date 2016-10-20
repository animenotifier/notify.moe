export function plural(count: number, singular: string): string {
	return (count === 1 || count === -1) ? (count + ' ' + singular) : (count + ' ' + singular + 's')
}