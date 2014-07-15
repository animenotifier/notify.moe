var animeUpdater = {
	animeListProvider: null,
	settings: {
		"userName": "",
		"quality": "",
		"subs": "",
		"otherSearch": "",
		"updateInterval": "5",
		"maxEpisodeDifference": "1",
		"animeListProvider": "anilist.co"
	},
	animeList: [],
	userAnimeListURL: "",
	optionsURL: "",
	backgroundCallback: function() {},
	notificationIdToLink: {},
	qualityRegEx: /([0-9]{3,4})p[^a-zA-Z]/,
	subsRegEx: /^\[([^\]]*)\]/,
	aniChartHTML: null,
	animeListCreated : false,

	// Request anime list
	requestAnimeList: function(newSettings, callback) {
		if(typeof newSettings == 'undefined' || newSettings == null) {
			//console.log("undefined settings");
			var updater = this;
			chrome.runtime.sendMessage({}, function(response) {
				updater.settings = response;
				updater.sendRequest();
			});
		} else {
			//console.log("defined settings: " + newSettings);
			this.settings = newSettings;
			this.sendRequest();
		}
	},

	// Send request
	sendRequest: function() {
		this.animeListProvider = animeListProviders[this.settings["animeListProvider"]];
		this.optionsURL = chrome.extension.getURL("src/options_custom/index.html");

		// Username check
		var userName = this.settings["userName"];
		if(userName.length == 0) {
			document.body.innerHTML = "Please specify your username in the extension <a href='" + this.optionsURL + "' target='_blank'>options</a>.";
			return;
		}

		document.body.appendChild(document.createTextNode("Loading anime list for: " + userName));

		if(typeof callback != 'undefined')
			this.backgroundCallback = callback;

		this.userAnimeListURL = this.animeListProvider.url + userName + this.animeListProvider.urlSuffix;
		this.animeListCreated = false;

		var req = new XMLHttpRequest();
		req.open("GET", this.userAnimeListURL, true);
		req.onload = this.showAnimeList.bind(this);
		req.send(null);

		// Sort please
		this.requestAniChart();
	},

	// Parse anime list
	parseAnimeList: function(html) {
		var animeList = [];

		if(this.animeListProvider.ignoreAfter) {
			var ignoreAfterIndex = html.indexOf(this.animeListProvider.ignoreAfter);
			
			if(ignoreAfterIndex != -1)
				html = html.substr(0, ignoreAfterIndex);
		}

		var animeRegEx = this.animeListProvider.animeRegEx;
		var progressRegEx = this.animeListProvider.progressRegEx;

		var match = animeRegEx.exec(html);
		while(match != null) {
			var watchedEpisodes = "-";
			var maxEpisodes = "-";

			var progressMatch = progressRegEx.exec(html.substr(animeRegEx.lastIndex));
			if(progressMatch != null) {
				watchedEpisodes = parseInt(progressMatch[1]);
				maxEpisodes = parseInt(progressMatch[2]);

				if(!watchedEpisodes)
					watchedEpisodes = 0;

				if(!maxEpisodes)
					maxEpisodes = "-";
			}

			var animeTitle = match[2].trim().replace(/_/g, " ").replace(/-/g, " ");

			animeList.push({
				id: match[1],
				title: decodeHtmlEntities(animeTitle).replace(/&[a-zA-Z]{2,10};/g, " "),
				originalTitle: animeTitle,
				watchedEpisodes: watchedEpisodes,
				nextEpisodeToWatch: watchedEpisodes + 1,
				maxEpisodes: maxEpisodes,
				days: 0,
				hours: 0,
				minutes: 0
			});

			match = animeRegEx.exec(html);
		}

		return animeList;
	},

	// 
	queryPossibleAnimeOptions: function(animeTitle, subsProvider, callback) {
		var customSearchTitle = localStorage["store.settings." + animeTitle + ":search"];

		if(customSearchTitle)
			customSearchTitle = customSearchTitle.replace(/"/g, "");

		var urlObject = {};
		getURLs(customSearchTitle ? customSearchTitle : animeTitle, "", subsProvider, urlObject);

		var req = new XMLHttpRequest();
		req.overrideMimeType('text/xml');
		req.open("GET", urlObject.rssUrl, true);
		req.onload = function(e) {
			var qualities = [
				{
					"value": "",
					"text": "*"
				}
			];

			var subs = [
				{
					"value": "",
					"text": "*"
				}
			];

			var qualitiesFound = {};
			var subsFound = {};

			// Find quality and subs which are available
			var itemList = e.target.responseXML.querySelectorAll("item");
			[].forEach.call(
				itemList, 
				function(item) {
					var title = item.getElementsByTagName("title")[0].innerHTML;

					// Quality
					var match = animeUpdater.qualityRegEx.exec(title);
					if(match != null) {
						var quality = match[1];

						if(!(quality in qualitiesFound)) {
							qualities.push({
								"value": quality,
								"text" : quality + "p"
							});

							qualitiesFound[quality] = true;
						}
					}

					// Subs
					var match = animeUpdater.subsRegEx.exec(title);
					if(match != null) {
						var sub = match[1];

						if(!(sub in subsFound)) {
							subs.push({
								"value": sub,
								"text" : sub
							});

							subsFound[sub] = true;
						}
					}
				}
			);

			qualities.sort(function(a, b) {
				return parseInt(a["value"]) - parseInt(b["value"]);
			});

			subs.sort(function(a, b) {
				return a["text"].localeCompare(b["text"]);
			});

			callback(animeTitle, qualities, subs);
		};
		req.send(null);
	},

	// Show anime list
	showAnimeList: function(e) {
		// Status code
		if(e.target.status != 200) {
			console.log(e.target.statusText);
		}

		// Parse the list
		this.animeList = this.parseAnimeList(e.target.responseText);

		var userName = this.settings["userName"];
		var maxEpisodeDifference = parseInt(this.settings["maxEpisodeDifference"]);
		var showImages = true;

		var episodeRegEx = /[ _]-[ _]([0-9]{1,3})v?\d?[ _][^a-zA-Z0-9]/;
		var fetchImage = this.animeListProvider.queryImage;

		if(this.animeList.length == 0) {
			document.body.innerHTML = "No anime found in the watching list of " + 
				"<a href='" + this.userAnimeListURL + "' target='_blank'>" + userName + "</a>.<br/>" + 
				"Are you sure the <a href='" + this.optionsURL + "' target='_blank'>options</a> are correctly set up?";
		} else {
			document.body.innerHTML = "";
		}

		// Each anime
		this.animeList.forEach(function(entry) {
			// Nyaa links
			var key = "store.settings." + entry.originalTitle;
			var quality = localStorage[key + ":quality"];
			var subs = localStorage[key + ":subs"];
			var customSearchTitle = localStorage[key + ":search"];

			if(customSearchTitle)
				customSearchTitle = customSearchTitle.replace(/"/g, "");

			var urlObject = {};
			getURLs(customSearchTitle ? customSearchTitle : entry.originalTitle, quality ? quality.replace(/"/g, "") : "", subs ? subs.replace(/"/g, "") : "", urlObject);

			entry.element = document.createElement("a");
			entry.element.className = "anime";
			entry.element.href = urlObject.url;
			entry.element.target = "_blank";

			// TODO: http://cdn.myanimelist.net/images/anime/6/60479.jpg
			entry.element.appendChild(document.createTextNode(entry.title + " "));

			document.body.appendChild(entry.element);

			var req = new XMLHttpRequest();
			req.overrideMimeType('text/xml');
			req.open("GET", urlObject.rssUrl, true);
			req.onload = function(e) {
				entry.latestEpisodeNumber = -1;
				//console.log(entry.title);

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

							if(!isNaN(episodeNumber) && episodeNumber > entry.latestEpisodeNumber) {
								entry.latestEpisodeNumber = episodeNumber;
								entry.latestEpisodeLink = link.replace("&amp;", "&");
							}
						}
					}
				);

				// Add episode number
				var episodesHTML;
				if(entry.latestEpisodeNumber == -1) {
					// Use watched episodes as a fallback
					//entry.latestEpisodeNumber = entry.watchedEpisodes;
					episodesHTML = "<span class='episodes latest-episode-fail'>";
				} else {
					episodesHTML = "<span class='episodes'>";
				}
				episodesHTML += "<span class='watched-episode-number'>" + entry.watchedEpisodes + "</span> "
					+ "<span class='latest-episode-number'>/ " + (entry.latestEpisodeNumber != -1 ? entry.latestEpisodeNumber : "?") + "</span> "
					+ "<span class='max-episode-part'>[" + entry.maxEpisodes + "]</span></span>";
				entry.element.innerHTML += episodesHTML;

				// Make it green
				if(entry.latestEpisodeNumber > entry.watchedEpisodes && entry.watchedEpisodes != "-") {
					entry.element.className += " new-episodes";
				} else if(entry.maxEpisodes > 0 && entry.watchedEpisodes == entry.maxEpisodes) {
					entry.element.className += " completed";
				}

				entry.element.title = "You watched " + entry.watchedEpisodes + " episodes out of " + entry.latestEpisodeNumber + " available (maximum: " + entry.maxEpisodes + ")";

				// Notification options
				var notificationOptions = {
					type: "basic",
					title: entry.title + " [Ep. " + entry.latestEpisodeNumber + "]",
					message: "New episode available\n",
					buttons: [{
						title: "Download"
					}]
				};

				// Fetch image
				fetchImage(entry, function(coverUrl) {
					notificationOptions.iconUrl = coverUrl;
					entry.imageUrl = coverUrl;

					if(showImages) {
						/*nyaa.style.background = "url('" + entry.imageUrl + "') no-repeat center center fixed";
						nyaa.style.backgroundSize = "cover";*/
						//nyaa.innerHTML = "<img src='" + entry.imageUrl + "' alt='" + entry.title + "'/> " + nyaa.innerHTML;

						var animeImg = document.createElement("img");
						animeImg.src = entry.imageUrl;
						animeImg.alt = entry.title;

						entry.element.appendChild(animeImg);
					}

					// Do we have latest episode info?
					if(entry.latestEpisodeNumber != -1) {
						//console.log(entry.title + "...");
						var key = "anime." + entry.title;
						var cached = localStorage.getObject(key);
						//console.log(cached);

						if(typeof(cached) != "undefined" && cached != null) {
							var latestEpisodeCached = parseInt(cached.latestEpisodeNumber);
							//console.log(latestEpisodeCached);

							// Just released?
							if(!isNaN(latestEpisodeCached) && entry.latestEpisodeNumber > latestEpisodeCached && entry.latestEpisodeNumber - entry.watchedEpisodes <= maxEpisodeDifference) {
								// Display notification
								chrome.notifications.create("", notificationOptions, function(notificationId) {
									animeUpdater.notificationIdToLink[notificationId] = entry.latestEpisodeLink;
								});
							}
						}

						// Save in cache
						localStorage.setObject(key, {
							latestEpisodeNumber: entry.latestEpisodeNumber
						});
					}
				});
			};
			req.send(null);
		});
		
		// This is not perfectly correct in terms of real concurrency
		// but it doesn't even matter: Worst case scenario (<0.1%) is that the list
		// gets sorted twice.
		if(this.aniChartHTML && !this.animeListSorted) {
			this.animeListSorted = true;
			this.sortAnimeList(this.aniChartHTML);
		}
		
		this.animeListCreated = true;
		
		// Create footer
		var footer = document.createElement("div");
		footer.className = "footer";
		footer.innerHTML = "<a href='" + this.userAnimeListURL + "' target='_blank' title='Profile'>" + userName + "</a> " + 
							"<a href='" + this.optionsURL + "' target='_blank' title='Options'><img src='http://blitzprog.org/images/anime-release-notifier/settings.png' alt='Options'/></a> " + 
							"<a href='http://anichart.net/airing' target='_blank' title='Chart'><img src='http://blitzprog.org/images/anime-release-notifier/chart.png' alt='Chart'/></a>";
		document.body.appendChild(footer);

		this.backgroundCallback();
	},

	// Request AniChart
	requestAniChart: function() {
		var req = new XMLHttpRequest();
		req.open("GET", "http://anichart.net/airing", true);
		req.onload = this.receiveAniChart.bind(this);
		req.send(null);
	},

	receiveAniChart: function(e) {
		this.aniChartHTML = e.target.responseText;

		// This is not perfectly correct in terms of real concurrency
		// but it doesn't even matter: Worst case scenario (<0.1%) is that the list
		// gets sorted twice.
		if(this.animeListCreated && !this.animeListSorted) {
			this.animeListSorted = true;
			this.sortAnimeList(this.aniChartHTML);
		}
	},

	// Sort anime list
	sortAnimeList: function(html) {
		var anichartAnimeInfoRegEx = /<div class="anime_info_sml">/g;
		var anichartTitleRegEx = /class="title_sml[^"']*"><a href=["'][^"']*["'] target="_blank">([^<]+)<\/a>/;
		var daysRegEx = /<span class="cd_day">([0-9]{0,3})<\/span>/;
		var hoursRegEx = /<span class="cd_hr">([0-9]{0,2})<\/span>/;
		var minutesRegEx = /<span class="cd_min">([0-9]{0,2})<\/span>/;

		var infoMatch = anichartAnimeInfoRegEx.exec(html);
		while(infoMatch != null) {
			var lastIndex = anichartAnimeInfoRegEx.lastIndex;
			var animeInfo = html.substr(lastIndex, html.indexOf('<div class="title_studio_sml">', lastIndex) - lastIndex);

			var daysMatch = daysRegEx.exec(animeInfo);
			var hoursMatch = hoursRegEx.exec(animeInfo);
			var minutesMatch = minutesRegEx.exec(animeInfo);
			var match = anichartTitleRegEx.exec(animeInfo);

			if(match != null) {
				var title = match[1].replace("-", " ").toUpperCase();
				var anime;

				for(var i = 0, len = this.animeList.length; i < len; i++) {
					anime = this.animeList[i];
					if(title == anime.title.replace("-", " ").toUpperCase().replace(/\(TV\)/g, "").trim()) {
						anime.days = daysMatch ? parseInt(daysMatch[1]) : 0;
						anime.hours = hoursMatch ? parseInt(hoursMatch[1]) : 0;
						anime.minutes = minutesMatch ? parseInt(minutesMatch[1]) : 0;
						anime.daysRounded = Math.round(anime.days + (anime.hours / 24.0));

						//console.log(anime);

						// Display release time
						var releaseTime = "<span class='release-time'>";
						if(anime.days == 0) {
							if(anime.hours == 0) {
								if(anime.minutes != 0)
									releaseTime += plural(anime.minutes, "minute");
							} else {
								releaseTime += plural(anime.hours, "hour");
							}
						} else {
							releaseTime += plural(anime.daysRounded, "day");
						}
						releaseTime += "</span>";

						anime.element.innerHTML = releaseTime + anime.element.innerHTML;

						break;
					}
				}
			}

			infoMatch = anichartAnimeInfoRegEx.exec(html);
		}

		// The actual sorting
		this.animeList.sort(function(a, b) {
			var aUndefined = false, bUndefined = false;

			if(a.days == 0 && a.hours == 0 && a.minutes == 0)
				aUndefined = true;

			if(b.days == 0 && b.hours == 0 && b.minutes == 0)
				bUndefined = true;

			return (a.days - b.days) * 24 * 60 + (a.hours - b.hours) * 60 + (a.minutes - b.minutes) + aUndefined * 999999999 - bUndefined * 999999999;
		});

		// Sort DOM elements
		var lastElement = this.animeList[0];
		this.animeList.forEach(function(entry) {
			entry.element.parentNode.insertBefore(entry.element, lastElement);
			lastElement = entry.element.nextSibling;

			if(entry.days != 0 || entry.hours != 0 || entry.minutes != 0) {
				var factor = entry.daysRounded; // * 24 * 60 + entry.hours * 60 + entry.minutes;
				entry.element.style.opacity = Math.max(1.0 - (factor * factor) / 10.0, 0.2);
			}
		});
	}
};

Storage.prototype.setObject = function(key, value) {
	this.setItem(key, JSON.stringify(value));
}

Storage.prototype.getObject = function(key) {
	var value = this.getItem(key);
	return value && JSON.parse(value);
}

// Special titles
var specialAnimeSearchNames = {
	"Fairy Tail (2014)":
	"Fairy Tail S2",

	"Sidonia no Kishi":
	"Knights of Sidonia",

	"Gokukoku no Brynhildr":
	"Brynhildr",

	"Mahouka Koukou no Rettousei":
	"Mahouka",

	"JoJo no Kimyou na Bouken: Stardust Crusaders":
	"JoJo Stardust Crusaders",

	"Psycho Pass New Edit Version":
	"Psycho Pass Extended Edition"
}

// Helper functions
var replaceSpecialAnimeSearchNames = function(animeTitle) {
	if(animeTitle in specialAnimeSearchNames)
		return specialAnimeSearchNames[animeTitle];
	else
		return animeTitle;
};

var plural = function(count, noun) {
	return count + " " + (count == 1 ? noun : noun + "s");
};

var decodeHtmlEntities = function(str) {
	return str.replace(/&#(\d+);/g, function(match, dec) {
		return String.fromCharCode(dec);
	});
};

var removeHtmlEntities = function(str) {
	return str.replace(/&#\d+;/g, " ").replace(/&[a-zA-Z]{2,10};/g, " ");
};

var makeAnimeSearchTitle = function(animeTitle) {
	return removeHtmlEntities(replaceSpecialAnimeSearchNames(animeTitle))
			.replace(/:/g, "")
			.replace(/&/g, "")
			.replace(/\(TV\)/g, "")
			.replace(/[^a-zA-Z0-9!']+/g, " ");
};

var encodeHtmlEntities = function(str) {
	var buf = [];
	for (var i=str.length-1;i>=0;i--) {
		buf.unshift(['&#', str[i].charCodeAt(), ';'].join(''));
	}
	return buf.join('');
};

var getURLs = function(animeTitle, quality, subs, obj) {
	var nyaaSearchTitle = makeAnimeSearchTitle(animeTitle)
							.replace(/_/g, "+")
							.replace(/ /g, "+")
							.replace(/\+\+/g, "+");

	var nyaaSuffix = ("&cats=1_37&filter=0&sort=2&term=" + nyaaSearchTitle + "+" + quality + "+" + subs).replace(/\++$/, "");
	
	obj.url = "http://www.nyaa.se/?page=search" + nyaaSuffix;
	obj.rssUrl = "http://www.nyaa.se/?page=rss" + nyaaSuffix;

	//var watchAnimeUrl = "http://www.watch-anime.net/" + entry.searchTitle.toLowerCase().replace(/ /g, "-") + "/" + entry.nextEpisodeToWatch;
	//var kissAnimeUrl = "http://kissanime.com/Anime/" + entry.searchTitle.replace(/ /g, "-") + "/Episode-" + ("000" + entry.nextEpisodeToWatch).slice(-3);
}