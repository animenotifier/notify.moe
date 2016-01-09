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
		if(response && response.error) {
			animeList.innerText = 'Error loading your anime list: ' + response.error;
			return;
		}

		if(error) {
			animeList.innerText = 'Error: ' + error;
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

			if(loggedIn) {
				if(anime.episodes.watched > 0 && anime.episodes.watched === anime.episodes.max) {
					var completed = document.createElement('a');
					completed.href = response.listUrl;
					completed.target = '_blank';
					completed.className = 'anime-completed';
					completed.title = 'You completed this anime.';

					var icon = document.createElement('div');
					icon.className = 'glyphicon glyphicon-ok-sign';
					completed.appendChild(icon);

					item.appendChild(completed);
				} else if(anime.episodes.available >= anime.episodes.next) {
					var download = document.createElement('a');
					download.href = anime.animeProvider.nextEpisodeUrl;
					download.target = '_blank';
					download.className = 'anime-download-link';

					if(anime.episodes.next && anime.episodes.next !== 0)
						download.title = 'Download episode ' + anime.episodes.next;

					var icon = document.createElement('div');
					icon.className = 'glyphicon glyphicon-cloud-download';
					download.appendChild(icon);
					item.appendChild(download);

					var behind = (anime.episodes.available - anime.episodes.watched);
					var episodes = document.createElement('span');
					episodes.className = 'episodes-behind';
					episodes.appendChild(document.createTextNode(behind + (behind === 1 ? ' episode' : ' episodes') + ' behind'));
					item.appendChild(episodes);
				} else if(anime.episodes.available > 0 && anime.episodes.available === anime.episodes.watched) {
					var ok = document.createElement('a');
					ok.href = anime.animeProvider.nextEpisodeUrl;
					ok.target = '_blank';
					ok.className = 'anime-up-to-date';
					ok.title = 'You watched ' + anime.episodes.watched + ' out of ' + anime.episodes.available + ' available.';

					var icon = document.createElement('div');
					icon.className = 'glyphicon glyphicon-ok';
					ok.appendChild(icon);

					item.appendChild(ok);
				} else if(anime.episodes.available === 0) {
					var warning = document.createElement('a');
					warning.href = anime.animeProvider.url;
					warning.target = '_blank';
					warning.className = 'anime-warning';
					warning.title = 'Could not find your anime title on the anime provider.';

					var icon = document.createElement('div');
					icon.className = 'glyphicon glyphicon-alert';
					warning.appendChild(icon);

					item.appendChild(warning);
				}
			}

			list.appendChild(item);
		});

		animeList.appendChild(list);
	});
};

window.loadAnimeList();