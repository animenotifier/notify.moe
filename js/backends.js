var animeBackends = {
	// Nyaa
	"nyaa.se": {
		host: "www.nyaa.se",
		episodeRegEx: /[ _]-[ _]([0-9]{1,3})v?\d?[ _][^a-zA-Z0-9]/,

		process: function(anime) {
			// Nyaa filter settings
			var key = "store.settings." + anime.originalTitle;
			var quality = localStorage[key + ":quality"];
			var subs = localStorage[key + ":subs"];
			var customSearchTitle = localStorage[key + ":search"];

			if(customSearchTitle)
				customSearchTitle = customSearchTitle.replace(/"/g, "");

			var urlObject = {};
			this.getURLs(customSearchTitle ? customSearchTitle : anime.originalTitle, quality ? quality.replace(/"/g, "") : "", subs ? subs.replace(/"/g, "") : "", urlObject);
			
			// Create link
			anime.element.href = urlObject.url;

			// Post-process
			var req = new XMLHttpRequest();
			req.overrideMimeType('text/xml');
			req.open("GET", urlObject.rssUrl, true);

			var episodeRegEx = this.episodeRegEx;
			var getImageAndNotify = this.getImageAndNotify;

			req.onload = function(e) {
				anime.latestEpisodeNumber = -1;
				//console.log(anime.title);

				// Find latest episode
				var itemList = e.target.responseXML.querySelectorAll("item");
				[].forEach.call(
					itemList, 
					function(item) {
						var title = item.getElementsByTagName("title")[0].innerHTML;
						var link = item.getElementsByTagName("link")[0].innerHTML;
						var pubDate = item.getElementsByTagName("pubDate")[0].innerHTML;

						var match = episodeRegEx.exec(title);
						if(match != null) {
							var episodeNumber = parseInt(match[1]);

							if(!isNaN(episodeNumber) && episodeNumber > anime.latestEpisodeNumber) {
								anime.latestEpisodeNumber = episodeNumber;
								anime.latestEpisodeLink = link.replace("&amp;", "&");
							}
						}
					}
				);

				// Add episode number
				var episodesHTML;
				if(anime.latestEpisodeNumber == -1) {
					// Use watched episodes as a fallback
					//anime.latestEpisodeNumber = anime.watchedEpisodes;
					episodesHTML = "<span class='episodes latest-episode-fail'>";
				} else {
					episodesHTML = "<span class='episodes'>";
				}
				episodesHTML += "<span class='watched-episode-number'>" + anime.watchedEpisodes + "</span> "
					+ "<span class='latest-episode-number'>/ " + (anime.latestEpisodeNumber != -1 ? anime.latestEpisodeNumber : "?") + "</span> "
					+ "<span class='max-episode-part'>[" + anime.maxEpisodes + "]</span></span>";
				anime.element.innerHTML += episodesHTML;

				// Make it green
				if(anime.latestEpisodeNumber > anime.watchedEpisodes && anime.watchedEpisodes != "-") {
					anime.element.className += " new-episodes";
				} else if(anime.maxEpisodes > 0 && anime.watchedEpisodes == anime.maxEpisodes) {
					anime.element.className += " completed";
				}

				anime.element.title = "You watched " + anime.watchedEpisodes + " episodes out of " + anime.latestEpisodeNumber + " available (maximum: " + anime.maxEpisodes + ")";

				getImageAndNotify(anime);
			};

			req.send(null);
		},

		getURLs: function(animeTitle, quality, subs, obj) {
			var nyaaSearchTitle = makeAnimeSearchTitle(animeTitle)
									.replace(/_/g, "+")
									.replace(/ /g, "+")
									.replace(/\+\+/g, "+");

			var nyaaSuffix = ("&cats=1_37&filter=0&sort=2&term=" + nyaaSearchTitle + "+" + quality + "+" + subs).replace(/\++$/, "");
			
			obj.url = "http://www.nyaa.se/?page=search" + nyaaSuffix;
			obj.rssUrl = "http://www.nyaa.se/?page=rss" + nyaaSuffix;
		},

		getImageAndNotify: function(anime) {
			var maxEpisodeDifference = parseInt(animeUpdater.settings["maxEpisodeDifference"]);
			var showImages = true;

			// Notification options
			var notificationOptions = {
				type: "basic",
				title: anime.title + " [Ep. " + anime.latestEpisodeNumber + "]",
				message: "New episode available\n",
				buttons: [{
					title: "Download"
				}]
			};

			// Fetch image
			animeUpdater.animeListProvider.queryImage(anime, function(coverUrl) {
				notificationOptions.iconUrl = coverUrl;
				anime.imageUrl = coverUrl;

				if(showImages) {
					/*nyaa.style.background = "url('" + anime.imageUrl + "') no-repeat center center fixed";
					nyaa.style.backgroundSize = "cover";*/
					//nyaa.innerHTML = "<img src='" + anime.imageUrl + "' alt='" + anime.title + "'/> " + nyaa.innerHTML;

					var animeImg = document.createElement("img");
					animeImg.src = anime.imageUrl;
					animeImg.alt = anime.title;

					anime.element.appendChild(animeImg);
				}

				// Do we have latest episode info?
				if(anime.latestEpisodeNumber != -1) {
					//console.log(anime.title + "...");
					var key = "anime." + anime.title;
					var cached = localStorage.getObject(key);
					//console.log(cached);

					if(typeof(cached) != "undefined" && cached != null) {
						var latestEpisodeCached = parseInt(cached.latestEpisodeNumber);
						//console.log(latestEpisodeCached);

						// Just released?
						if(!isNaN(latestEpisodeCached) && anime.latestEpisodeNumber > latestEpisodeCached && anime.latestEpisodeNumber - anime.watchedEpisodes <= maxEpisodeDifference) {
							// Display notification
							chrome.notifications.create("", notificationOptions, function(notificationId) {
								animeUpdater.notificationIdToLink[notificationId] = anime.latestEpisodeLink;
							});
						}
					}

					// Save in cache
					localStorage.setObject(key, {
						latestEpisodeNumber: anime.latestEpisodeNumber
					});
				}
			});
		}
	},

	// Shana project
	"shanaproject.com": {
		host: "shanaproject.com",

		process: function(anime) {
			var searchTitle = makeAnimeSearchTitle(anime.title)
									.replace(/_/g, "+")
									.replace(/ /g, "+")
									.replace(/\+\+/g, "+");

			var shanaURL = "http://www.shanaproject.com/search/?title=" + searchTitle;

			anime.element.href = shanaURL;

			// Check new
			var req = new XMLHttpRequest();
			req.open("GET", shanaURL, true);
			req.onload = function(e) {
				var html = e.target.responseText;

				if(html.indexOf('<div class="release_episode">' + anime.nextEpisodeToWatch + '</div>') != -1)
					anime.element.className += " new-episodes";
			};
			req.send(null);

			// TODO: Add eps number available
			anime.latestEpisodeNumber = -1;

			// Add episode number
			var episodesHTML;
			if(anime.latestEpisodeNumber == -1) {
				// Use watched episodes as a fallback
				//anime.latestEpisodeNumber = anime.watchedEpisodes;
				episodesHTML = "<span class='episodes latest-episode-fail'>";
			} else {
				episodesHTML = "<span class='episodes'>";
			}
			episodesHTML += "<span class='watched-episode-number'>" + anime.watchedEpisodes + "</span> "
				+ "<span class='latest-episode-number'>/ " + (anime.latestEpisodeNumber != -1 ? anime.latestEpisodeNumber : "?") + "</span> "
				+ "<span class='max-episode-part'>[" + anime.maxEpisodes + "]</span></span>";
			anime.element.innerHTML += episodesHTML;

			// Fetch image
			animeUpdater.animeListProvider.queryImage(anime, function(coverUrl) {
				anime.imageUrl = coverUrl;

				var animeImg = document.createElement("img");
				animeImg.src = anime.imageUrl;
				animeImg.alt = anime.title;

				anime.element.appendChild(animeImg);
			});
		}
	},

	// Kissanime
	"kissanime.com": {
		host: "kissanime.com",

		process: function(anime) {
			//var watchAnimeUrl = "http://www.watch-anime.net/" + anime.searchTitle.toLowerCase().replace(/ /g, "-") + "/" + anime.nextEpisodeToWatch;
			var kissAnimeURL = "http://kissanime.com/Anime/"
									+ anime.title.replace(/ /g, "-").replace(/!/g, "").replace(/:/g, "")
									+ "/Episode-" + ("000" + anime.nextEpisodeToWatch).slice(-3);

			anime.element.href = kissAnimeURL;

			// Check new
			var req = new XMLHttpRequest();
			req.open("GET", kissAnimeURL, true);
			req.onload = function(e) {
				var html = e.target.responseText;

				if(html.indexOf("ContentVideo") != -1)
					anime.element.className += " new-episodes";
			};
			req.send(null);

			// Fetch image
			animeUpdater.animeListProvider.queryImage(anime, function(coverUrl) {
				anime.imageUrl = coverUrl;

				var animeImg = document.createElement("img");
				animeImg.src = anime.imageUrl;
				animeImg.alt = anime.title;

				anime.element.appendChild(animeImg);
			});
		}
	}
};