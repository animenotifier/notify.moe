export function fixGenre(genre: string) {
	return genre.replace(/ /g, '').replace(/-/g, '').toLowerCase()
}