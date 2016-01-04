var search = document.getElementById('search');
var searchResults = document.getElementById('search-results');
var lastRequest = undefined

window.searchAnime = function(e) {
	if(e !== 'force') {
		e = event || window.event;

		var keyCode = e.keyCode || e.which;

		if(keyCode !== 13)
			return;
	}

	if(lastRequest) {
		lastRequest.abort();
		lastRequest = undefined
	}

	lastRequest = kaze.get('/_/anime/search/' + search.value, function(error, body) {
		searchResults.innerHTML = body;
	});
};

if(search) {
	search.select();

	if(search.value)
		window.searchAnime('force');
}