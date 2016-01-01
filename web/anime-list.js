function AnimeList(json, $animeList, maxEpisodeDifference, notificationCallBack) {
	this.json = json;
	this.element = $animeList;

	this.failCount = 0;
	this.successCount = 0;
	this.newCount = 0;
	this.listUrl = json.listUrl;

	$animeList.html("");

	json.watching.forEach(function(anime) {
		var cssClass = "anime";

		// Action URL
		anime.actionUrl = anime.url;

		if(anime.animeProvider.url)
			anime.actionUrl = anime.animeProvider.url;

		if(anime.animeProvider.nextEpisodeUrl && anime.episodes.available >= anime.episodes.next)
			anime.actionUrl = anime.animeProvider.nextEpisodeUrl;

		//if(anime.animeProvider.videoUrl && typeof anime.animeProvider.videoHash != "undefined" && anime.animeProvider.videoHash != "" && typeof asp.wrap != "undefined")
		//	anime.animeProvider.videoUrl = asp.wrap(anime.animeProvider.videoHash);

		// New episodes
		if(anime.episodes.watched < anime.episodes.available - anime.episodes.offset) {
			cssClass += " new-episodes";
			this.newCount += 1;
		} else if(anime.episodes.max > 0 && anime.episodes.watched == anime.episodes.max) {
			cssClass += " completed";
		}

		if(anime.episodes.available == -1)
			this.failCount += 1;
		else
			this.successCount += 1;

		// Available
		var available = '?';

		if(anime.episodes.available !== -1) {
			available = anime.episodes.available - anime.episodes.offset;
		}

		var max = (anime.episodes.max != -1 ? anime.episodes.max : '?');
		var tooltip = "You watched " + anime.episodes.watched + " episodes out of " + available + " available (maximum: " + max + ")";

		$animeList.append("<a href='" + anime.actionUrl.replace(/'/g, "%27") + "' target='_blank' class='" + cssClass + "' title='" + tooltip + "' itemscope itemtype='http://schema.org/ViewAction'>" +
			'<span class="title">' + anime.title + '</span>' +
			(anime.airingDate.timeStamp != -1 ? ('<span class="release-time">' + anime.airingDate.remainingString + '</span>') : '') +
			'<span class="episodes"><span class="watched-episode-number">' + (anime.episodes.watched != -1 ? anime.episodes.watched : '?') +
			'</span> <span class="latest-episode-number">/ ' + available +
			'</span> <span class="max-episode-part">[' + max +
			']</span></span>' +
			'<img src="' + anime.image + '" alt="Cover image">' +
			'</a>' +
			''//((anime.animeProvider.videoUrl != "") ? ('<a href="' + anime.animeProvider.videoUrl + '" class="direct-video-link" target="_blank">V</a>') : '')
		);

		// Notifications
		anime.sendNotification = function() {
			var displayNotification = function() {
				var notification = new Notification(anime.title, {
					body: "Episode " + anime.episodes.available + " released!",
					icon: anime.image
				});
			};

			if(!("Notification" in window)) {
				console.log("Browser doesn't support notifications");
				return;
			}

			if(Notification.permission === "granted") {
				displayNotification();
			} else {
				Notification.requestPermission(function(permission) {
					if(permission === "granted") {
						displayNotification();
					}
				});
			}
		};

		// Notification callback
		if(notificationCallBack) {
			var key = anime.title + ":episodes-available";
			var availableCached = parseInt(localStorage.getItem(key));

			if(availableCached && anime.episodes.available > availableCached && availableCached != -1 && anime.episodes.available > anime.episodes.watched && anime.episodes.available <= anime.episodes.watched + maxEpisodeDifference) {
				notificationCallBack(anime);
			}

			localStorage.setItem(key, anime.episodes.available);
		}
	}.bind(this));

	this.length = this.successCount + this.failCount;

	if(this.length != 0)
		this.successRate = parseFloat(this.successCount) / this.length;
	else
		this.successRate = 0;

	// Cache anime list
	localStorage.animeListHTMLCache = $animeList.html();
}