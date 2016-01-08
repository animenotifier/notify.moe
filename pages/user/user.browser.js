window.loadAnimeList = function() {
	var animeList = document.getElementById('animeList');

	// Loading animation
	animeList.innerHTML =
		'<div class="sk-folding-cube">' +
			'<div class="sk-cube1 sk-cube"></div>' +
			'<div class="sk-cube2 sk-cube"></div>' +
			'<div class="sk-cube4 sk-cube"></div>' +
			'<div class="sk-cube3 sk-cube"></div>' +
		'</div>';

	var userName = window.location.pathname.substring(2);

	kaze.getJSON('/api/animelist/' + userName, function(error, response) {
		if(error) {
			if(response && response.error)
				animeList.innerText = 'Error loading your anime list: ' + response.error;
			else
				animeList.innerText = 'Error loading your anime list: ' + error;
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
					view.title = 'Download episode ' + anime.episodes.next;

				var icon = document.createElement('div');
				icon.className = 'glyphicon glyphicon-cloud-download';
				view.appendChild(icon);
				item.appendChild(view);

				var behind = (anime.episodes.available - anime.episodes.watched);
				var episodes = document.createElement('span');
				episodes.className = 'episodes-behind';
				episodes.appendChild(document.createTextNode(behind + (behind === 1 ? ' episode' : ' episodes') + ' behind'));
				item.appendChild(episodes);
			}

			list.appendChild(item);
		});

		animeList.appendChild(list);
	});
};

window.loadAnimeList();