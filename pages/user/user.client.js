var weekDays = ['Sunday', 'Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday'];
var userName = $('nick').textContent.trim();

const loadingMessages = [
	'We\'re testing your patience...',
	'The bits are flowing slowly today...',
	'Calculating who is best girl...',
	'Connecting to another world line...',
	'Server-chan where are you?...',
	'Summoning the loli queen...',
	'20 years later...'
]

window.getWeekDay = function(timeStamp) {
	var date = new Date(timeStamp * 1000);
	return weekDays[date.getDay()];
};

window.loadAnimeList = function(clearCache) {
	var animeList = $('anime-list-container');

	// Loading animation
	animeList.innerHTML =
		'<div class="anime-list appear"><div class="sk-folding-cube">' +
			'<div class="sk-cube1 sk-cube"></div>' +
			'<div class="sk-cube2 sk-cube"></div>' +
			'<div class="sk-cube4 sk-cube"></div>' +
			'<div class="sk-cube3 sk-cube"></div>' +
		'</div>' +
		`<div class="loading-message">${loadingMessages[Math.floor(Math.random() * loadingMessages.length)]}</div></div>`;

	$.getJSON('/api/animelist/' + userName + (clearCache ? '/clearCache' : '')).then(response => {
		if(response && response.error) {
			animeList.textContent = 'Error loading your anime list: ' + response.error;
			return;
		}

		if(!response.watching) {
			animeList.textContent = 'There are no anime your watching list.';
			return;
		}

		animeList.innerHTML = '';
		var listProviderLink = $('list-provider-link');
		if(listProviderLink) {
			listProviderLink.href = response.listUrl;
		}

		var list = document.createElement('ul');
		list.className = 'anime-list';

		var loggedIn = animeList.dataset.logged === 'true';
		var newAnimeCount = 0;

		response.watching.forEach(function(anime) {
			var item = document.createElement('li');
			item.className = 'anime';

			// Image
			var image = document.createElement('img');
			image.src = anime.image;
			image.alt = anime.preferredTitle;
			image.className = 'anime-list-image';
			item.appendChild(image);

			// Link
			var link = document.createElement('a');
			link.appendChild(document.createTextNode(anime.title[response.titleLanguage]));

			if(anime.id) {
				link.href = '/anime/' + anime.id;
				link.className = 'anime-title ajax';
			} else {
				link.href = anime.url;
				link.target = '_blank';
				link.className = 'anime-title';
			}

			item.appendChild(link);

			// Spacer
			var spacer = document.createElement('span');
			spacer.className = 'spacer';
			item.appendChild(spacer);

			var addIconLink = function(iconName, url, linkClass, tooltip) {
				var link = document.createElement(url ? 'a' : 'div');

				if(url) {
					link.href = url;
					link.target = '_blank';
				}

				link.className = linkClass;
				link.title = tooltip;

				var icon = document.createElement('i');
				icon.className = 'fa fa-' + iconName + ' anime-status-icon fa-fw';
				link.appendChild(icon);

				item.appendChild(link);
			};

			var addAiringDate = function() {
				if(anime.airingDate.remaining === null)
					return;

				var airingDate = document.createElement('span');
				airingDate.className = 'airing-date';

				var airingDatePrefix = document.createElement('span');
				airingDatePrefix.className = 'airing-date-prefix';
				airingDatePrefix.textContent = anime.airingDate.remaining > 0 ? 'Airing in ' : 'Aired ';
				airingDate.appendChild(airingDatePrefix);
				airingDate.appendChild(document.createTextNode(anime.airingDate.remainingString));
				airingDate.title = window.getWeekDay(anime.airingDate.timeStamp);
				item.appendChild(airingDate);
			};

			if(loggedIn) {
				if(anime.episodes.watched > 0 && anime.episodes.watched === anime.episodes.max) {
					addIconLink(
						'check-circle',
						response.listUrl,
						'anime-completed',
						'You completed this anime.'
					);
				} else if(anime.episodes.available >= anime.episodes.next) {
					var behind = (anime.episodes.available - anime.episodes.watched);
					var episodes = document.createElement('span');
					episodes.className = 'episodes-behind';
					episodes.appendChild(document.createTextNode(behind));
					episodes.title = behind + ' new episode' + (behind === 1 ? '' : 's');
					item.appendChild(episodes);

					var isDownload = (anime.animeProvider.type === undefined || anime.animeProvider.type === 'download');
					var isBatch = anime.animeProvider.isBatch

					addIconLink(
						isDownload ? (isBatch ? 'archive' : 'cloud-download') : 'eye',
						anime.animeProvider.nextEpisode ? anime.animeProvider.nextEpisode.url : anime.animeProvider.url,
						'anime-download-link',
						isBatch ? ('You watched ' + anime.episodes.watched + ' out of ' + anime.episodes.available + ' available.') : ((isDownload ? 'Download' : 'Watch') + ' episode ' + anime.episodes.next)
					);

					newAnimeCount++;
				} else if(anime.episodes.available > 0 && anime.episodes.available === anime.episodes.watched) {
					addAiringDate();

					addIconLink(
						'check',
						anime.animeProvider.nextEpisode ? anime.animeProvider.nextEpisode.url : anime.animeProvider.url,
						'anime-up-to-date',
						'You watched ' + anime.episodes.watched + ' out of ' + anime.episodes.available + ' available.'
					);
				} else if(anime.episodes.available === 0) {
					addAiringDate();

					addIconLink(
						'exclamation-circle',
						anime.animeProvider.url,
						'anime-warning',
						'Could not find your anime title on the anime provider.'
					);
				}
			}

			// If this is embedded in an iframe, send additional info to the top site
			if(window.top) {
				window.top.postMessage(JSON.stringify({
					newAnimeCount: newAnimeCount
				}), '*')
			}

			list.appendChild(item);
		});

		animeList.appendChild(list);
		$.ajaxifyLinks();
	}).catch(error => {
		error = JSON.parse(error)
		animeList.innerHTML = '' // Remove loading animation
		let pre = document.createElement('pre')
		pre.appendChild(document.createTextNode(JSON.stringify(error, null, '\t')))
		animeList.appendChild(pre)

		if(error.error && error.error.endsWith('username has not been specified'))
			return

		let p = document.createElement('p')
		p.className = 'report-bug-hint'
		p.innerHTML = 'Seems like you encountered a problem. Want to <a href="/threads/new/bug" class="ajax">report the bug</a> on the forums?'
		animeList.appendChild(p)
		$.ajaxifyLinks(p)
	});
};

window.editMessage = function() {
	alert('work in progress')
}

window.deleteMessage = messageId => $.post('/api/messages/delete/' + messageId).then(window.loadMessages)

window.sendMessage = function() {
	let userName = $('nick').textContent
	let postInput = $('post-input')

	if(!postInput.value)
		return

	$.post('/api/messages', {
		recipient: userName,
		text: postInput.value
	}).then(response => {
		postInput.value = ''
		window.loadMessages()
	})
}

window.loadMessages = function() {
	let posts = $('posts')

	if(!posts)
		return

	$.get('/_/messages/user/' + userName).then(response => {
		posts.innerHTML = response
		updateAvatars()
		$.ajaxifyLinks()
		// $.executeScripts(posts)
		// $.emit('DOMContentLoaded')
	})
}

window.follow = function(userName) {
	$.get('/api/users/follow/' + userName).then(response => {
		$('unfollow').style.display = 'block'
		$('follow').style.display = 'none'
	})
}

window.unfollow = function(userName) {
	$.get('/api/users/unfollow/' + userName).then(response => {
		$('follow').style.display = 'block'
		$('unfollow').style.display = 'none'
	})
}

window.loadAnimeList()
window.loadMessages()
