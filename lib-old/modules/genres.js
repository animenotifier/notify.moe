arn.fixGenre = genre => {
	return genre.replace(/ /g, '').replace(/-/g, '').toLowerCase()
}