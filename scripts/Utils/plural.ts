const specialized = {
	"new activity": "new activities"
}

export default function plural(count: number, singular: string): string {
	if(count === 1 || count === -1) {
		return count + " " + singular
	}

	if(specialized[singular]) {
		return count + " " + specialized[singular]
	}

	return count + " " + singular + "s"
}
