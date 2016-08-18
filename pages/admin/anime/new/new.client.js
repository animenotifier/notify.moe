let animeTextArea = $('anime-json')
let status = $('status')

animeTextArea.focus()

window.addAnime = () => {
	let anime = null
	
	try {
		anime = JSON.parse(animeTextArea.value)
		$.post('/api/anime/add', anime)
		.then(response => {
			response = JSON.parse(response)
			status.innerHTML = `Added anime: <a href="/anime/${response.id}" target="_blank">${response.title.romaji}</a>`
		})
		.catch(error => status.textContent = error)
	} catch(error) {
		console.error(error)
		status.textContent = error
	}
}