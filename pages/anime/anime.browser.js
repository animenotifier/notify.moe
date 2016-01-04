var search = document.getElementById('search');
var searchResults = document.getElementById('search-results');
var lastRequest = undefined

window.searchAnime = function() {
	if(lastRequest) {
		lastRequest.abort();
		lastRequest = undefined
	}

	lastRequest = kaze.get('/_/anime/search/' + search.value, function(error, body) {
		searchResults.innerHTML = body;
	});
};

search.select();
window.searchAnime();