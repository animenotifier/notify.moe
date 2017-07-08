export function* findAll(className: string): IterableIterator<HTMLElement> {
	let elements = document.getElementsByClassName(className)
	
	for(let i = 0; i < elements.length; ++i) {
		yield elements[i] as HTMLElement
	}
}

export function delay<T>(millis: number, value?: T): Promise<T> {
	return new Promise(resolve => setTimeout(() => resolve(value), millis))
}

export function plural(count: number, singular: string): string {
	return (count === 1 || count === -1) ? (count + " " + singular) : (count + " " + singular + "s")
}

export function canUseWebP(): boolean {
    let canvas = document.createElement("canvas")

    if(!!(canvas.getContext && canvas.getContext("2d"))) {
        // WebP representation possible
        return canvas.toDataURL("image/webp").indexOf("data:image/webp") === 0
    } else {
        // In very old browsers (IE 8) canvas is not supported
        return false
    }
}