window.loadAnimeList = function() {
	var animeList = document.getElementById('animeList');
	animeList.innerText = 'Loading your anime list...';

	var userName = window.location.pathname.substring(2);

	kaze.getJSON('/api/animelist/' + userName, function(error, response) {
		if(error) {
			animeList.innerText = 'Error loading your anime list: ' + error.toString();
			return;
		}

		if(!response.watching) {
			animeList.innerText = 'There are no anime your watching list.';
			return;
		}

		animeList.innerHTML = '';

		var list = document.createElement('ul');
		list.className = 'anime-list';

		var loggedIn = animeList.dataset.logged === 'true';

		response.watching.forEach(function(anime) {
			var item = document.createElement('li');
			item.className = 'anime';

			var link = document.createElement('a');
			link.appendChild(document.createTextNode(anime.title));
			link.href = anime.url;
			link.target = '_blank';
			link.className = 'anime-link';

			item.appendChild(link);

			if(loggedIn && anime.episodes.available >= anime.episodes.next) {
				var view = document.createElement('a');
				view.href = anime.animeProvider.nextEpisodeUrl;
				view.target = '_blank';
				view.className = 'anime-view-link';

				if(anime.episodes.next && anime.episodes.next !== 0)
					view.title = 'Watch episode ' + anime.episodes.next;

				var icon = document.createElement('div');
				icon.className = 'glyphicon glyphicon-eye-open';
				view.appendChild(icon);
				item.appendChild(view);
			}

			list.appendChild(item);
		});

		animeList.appendChild(list);
	});
};

window.loadAnimeList();