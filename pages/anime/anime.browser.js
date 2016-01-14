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

	window.activateSearch = function() {
		search.disabled = false;
		search.select();
		window.searchAnime();
	};

	window.downloadSearchList = function() {
		kaze.getJSON('/api/searchlist', function(error, json) {
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