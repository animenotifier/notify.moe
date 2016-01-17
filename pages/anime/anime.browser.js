var animeContainer = document.querySelector('.anime-container');

if(animeContainer && animeContainer.dataset.id) {
	console.log(animeContainer.dataset.id);
	makeSaveable('/api/anime/' + animeContainer.dataset.id);
	getGravatarImages();
} else {
	var search = document.getElementById('search');
	var searchResults = document.getElementById('search-results');
	var allAnimeObject = document.getElementById('all-anime');
	var lastRequest = undefined;
	var maxSearchResults = 14;
	var allAnime = localStorage.getItem('allAnimeTitles');
	var animeTitles = null;
	var redownload = true;

	// Copyright (c) 2015 Jordan Thomas
	window.similar = function(s1, s2) {
		var m = 0;
		var i;
		var j;

		// Exit early if either are empty.
		if(s1.length === 0 || s2.length === 0) {
			return 0;
		}

		// Exit early if they're an exact match.
		if(s1 === s2) {
			return 1;
		}

		var range = (Math.floor(Math.max(s1.length, s2.length) / 2)) - 1;
		var s1Matches = new Array(s1.length);
		var s2Matches = new Array(s2.length);

		for (i = 0; i < s1.length; i++) {
			var low  = (i >= range) ? i - range : 0;
			var high = (i + range <= s2.length) ? (i + range) : (s2.length - 1);

			for (j = low; j <= high; j++) {
				if(s1Matches[i] !== true && s2Matches[j] !== true && s1[i] === s2[j]) {
					++m;
					s1Matches[i] = s2Matches[j] = true;
					break;
				}
			}
		}

		// Exit early if no matches were found.
		if(m === 0) {
			return 0;
		}

		// Count the transpositions.
		var k = 0;
		var numTrans = 0;

		for(i = 0; i < s1.length; i++) {
			if(s1Matches[i] === true) {
				for (j = k; j < s2.length; j++) {
					if(s2Matches[j] === true) {
						k = j + 1;
						break;
					}
				}

				if(s1[i] !== s2[j]) {
					++numTrans;
				}
			}
		}

		var weight = (m / s1.length + m / s2.length + (m - (numTrans / 2)) / m) / 3;
		var l = 0;
		var p = 0.1;

		if(weight > 0.7) {
			while(s1[l] === s2[l] && l < 4) {
				++l;
			}

			weight = weight + l * p * (1 - weight);
		}

		return weight;
	}

	window.searchAnime = function() {
		var term = search.value.trim().toLowerCase();

		if(!term) {
			searchResults.innerHTML = animeTitles.length + ' anime in the database. Powered by Anilist.';
			searchResults.className = 'anime-count';
			return;
		}

		searchResults.className = '';
		searchResults.innerHTML = '';

		var i = 0;
		var results = [];

		for(i = 0; i < animeTitles.length; i++) {
			var title = animeTitles[i];
			var titleLower = title.toLowerCase();

			if(titleLower === term) {
				results.push({
					title: title,
					similarity: 1
				});
			} else if(term.length >= 2 && titleLower.startsWith(term)) {
				results.push({
					title: title,
					similarity: 0.999
				});
			} else if(term.length >= 3 && titleLower.indexOf(term) !== -1) {
				results.push({
					title: title,
					similarity: 0.989
				});
			} else {
				var similarity = window.similar(titleLower, term);

				if(similarity >= 0.87) {
					results.push({
						title: title,
						similarity: similarity
					});
				}
			}

			/*if(results.length >= maxSearchResults)
				break;*/
		}

		results.sort(function(a, b) {
			if(a.similarity === b.similarity)
				return a.title.localeCompare(b.title)

			return a.similarity > b.similarity ? -1 : 1
		});

		if(results.length >= maxSearchResults)
			results.length = maxSearchResults;

		for(i = 0; i < results.length; i++) {
			var result = results[i];

			var element = document.createElement('a');
			element.className = 'search-result ajax';
			element.href = '/anime/' + allAnime[result.title];
			element.style.opacity = (result.similarity - 0.8) * 5;
			element.appendChild(document.createTextNode(result.title));

			searchResults.appendChild(element);
		}
	};

	window.activateSearch = function() {
		search.disabled = false;
		search.select();
		window.searchAnime();
	};

	window.downloadSearchList = function() {
		kaze.getJSON('/api/searchlist').then(function(json) {
			allAnime = json;
			animeTitles = Object.keys(allAnime);
			console.log(animeTitles.length);
			localStorage.setItem('allAnimeTitles', JSON.stringify(allAnime));
			window.activateSearch();
		});
	};

	if(allAnime && allAnime !== null) {
		allAnime = JSON.parse(allAnime);
		animeTitles = Object.keys(allAnime);

		if(animeTitles.length === parseInt(search.dataset.count))
			window.activateSearch();
		else
			window.downloadSearchList();
	} else {
		window.downloadSearchList();
	}
}