let animeContainer = document.querySelector('.anime-container')

$.clear = element => {
	while(element.lastChild) {
		element.removeChild(element.lastChild)
	}
}

if(animeContainer && animeContainer.dataset.id) {
	makeSaveable('/api/anime/' + animeContainer.dataset.id)
} else {
	let search = $('search')
	let searchResults = $('search-results')
	let allAnimeObject = $('all-anime')
	let lastRequest = undefined
	let maxSearchResults = 14
	let allAnime = localStorage.getItem('allAnimeTitles')
	let animeTitles = null
	let redownload = true

	// Copyright (c) 2011 Andrei Mackenzie
	let levenshtein = function(a, b) {
		if(a.length === 0)
			return b.length

		if(b.length === 0)
			return a.length

		let matrix = []

		// increment along the first column of each row
		for(let i = 0; i <= b.length; i++) {
			matrix[i] = [i]
		}

		// increment each column in the first row
		for(let j = 0; j <= a.length; j++) {
			matrix[0][j] = j
		}

		// Fill in the rest of the matrix
		for(let i = 1; i <= b.length; i++) {
			for(j = 1; j <= a.length; j++) {
				if(b.charAt(i-1) === a.charAt(j-1)) {
					matrix[i][j] = matrix[i-1][j-1]
				} else {
					matrix[i][j] =
						Math.min(matrix[i-1][j-1] + 1, // substitution
						Math.min(matrix[i][j-1] + 1, // insertion
						matrix[i-1][j] + 1)) // deletion
				}
			}
		}

		return matrix[b.length][a.length]
	}

	window.getSearchResults = function(resolve, reject) {
		let term = search.value.trim().toLowerCase()

		if(!term) {
			searchResults.innerHTML = animeTitles.length + ' titles in the database. Powered by Anilist.'
			searchResults.className = 'anime-count'
			return reject(new Error('no search term'))
		}

		let results = []

		for(let title of animeTitles) {
			let titleLower = title.toLowerCase()

			if(titleLower === term) {
				results.push({
					title: title,
					distance: 0
				})
			} else if(term.length >= 2 && titleLower.startsWith(term)) {
				results.push({
					title: title,
					distance: 0.1
				})
			} else if(term.length >= 3 && titleLower.indexOf(term) !== -1) {
				results.push({
					title: title,
					distance: 0.2
				})
			}
		}

		// Nothing found? Do a precise search:
		/*if(results.length === 0 && term.length >= 4) {
			for(i = 0 i < animeTitles.length i++) {
				let title = animeTitles[i]
				let titleLower = title.toLowerCase()
				let distance = levenshtein(titleLower, term)

				if(distance <= title.length / 2) {
					results.push({
						title: title,
						distance: distance
					})
				}
			}
		}*/

		results.sort(function(a, b) {
			if(a.distance === b.distance)
				return a.title.localeCompare(b.title)

			return a.distance > b.distance ? 1 : -1
		})

		if(results.length >= maxSearchResults)
			results.length = maxSearchResults

		resolve(results)
	}

	window.displaySearchResults = function(results) {
		searchResults.className = ''
		$.clear(searchResults)

		for(let result of results) {
			let element = document.createElement('a')
			element.className = 'search-result ajax'
			element.href = '/anime/' + allAnime[result.title]
			//element.style.opacity = (result.similarity - 0.8) * 5
			element.appendChild(document.createTextNode(result.title))

			searchResults.appendChild(element)
		}

		$.ajaxifyLinks()
	}

	window.searchAnime = function() {
		return new Promise(window.getSearchResults).then(window.displaySearchResults).catch(function(error) {
			if(error.message !== 'no search term')
				console.error(error)
		})
	}

	window.activateSearch = function() {
		search.disabled = false
		search.select()
		window.searchAnime()
	}

	window.downloadSearchList = function() {
		$.getJSON('/api/searchlist').then(function(json) {
			allAnime = json
			animeTitles = Object.keys(allAnime)
			console.log(animeTitles.length)
			localStorage.setItem('allAnimeTitles', JSON.stringify(allAnime))
			window.activateSearch()
		})
	}

	if(allAnime && allAnime !== null) {
		allAnime = JSON.parse(allAnime)
		animeTitles = Object.keys(allAnime)

		if(animeTitles.length === parseInt(search.dataset.count))
			window.activateSearch()
		else
			window.downloadSearchList()
	} else {
		window.downloadSearchList()
	}
}