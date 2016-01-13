var animeContainer = document.querySelector('.anime-container');

if(animeContainer && animeContainer.dataset.id) {
	console.log(animeContainer.dataset.id);
	makeSaveable('/api/anime/' + animeContainer.dataset.id);
	getGravatarImages();
} else {
	var search = document.getElementById('search');
	var searchResults = document.getElementById('search-results');
	var allAnimeObject = document.getElementById('all-anime');
	var allAnime = JSON.parse(allAnimeObject.text);
	var animeTitles = Object.keys(allAnime);
	var lastRequest = undefined;
	var maxSearchResults = 14;

	window.similar = function(a, b) {
	    var lengthA = a.length;
	    var lengthB = b.length;

	    var equivalency = 0;

	    var minLength = (a.length > b.length) ? b.length : a.length;
	    var maxLength = (a.length < b.length) ? b.length : a.length;

	    for(var i = 0; i < minLength; i++) {
	        if(a[i] === b[i]) {
	            equivalency++;
	        }
	    }

	    return equivalency / maxLength;
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

			if(title.toLowerCase().indexOf(term) === -1)
				continue;

			results.push(title);

			/*if(results.length >= maxSearchResults)
				break;*/
		}

		results.sort(function(a, b) {
			var similarityA = window.similar(a.toLowerCase(), term)
			var similarityB = window.similar(b.toLowerCase(), term)

			if(similarityA === similarityB)
				return a > b

			return similarityA < similarityB
		});

		if(results.length >= maxSearchResults)
			results.length = maxSearchResults;

		for(i = 0; i < results.length; i++) {
			var title = results[i];

			var element = document.createElement('a');
			element.className = 'search-result ajax';
			element.href = '/anime/' + allAnime[title];
			element.appendChild(document.createTextNode(title));

			searchResults.appendChild(element);
		}
	};

	if(search) {
		search.select();

		window.searchAnime();
	}
}