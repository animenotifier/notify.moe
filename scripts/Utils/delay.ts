export function delay<T>(millis: number, value?: T): Promise<T> {
	return new Promise(resolve => setTimeout(() => resolve(value), millis))
}