var search = document.getElementById('search');
var searchResults = document.getElementById('search-results');
var lastRequest = undefined

window.searchAnime = function() {
	/*if(force !== 'force') {
		ev = ev || event || window.event;

		if(ev.keyCode !== 13)
			return;
	}*/

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

	window.searchAnime();
}