var animeListProviders = {
	// Myanimelist
	"myanimelist.net": {
		host: "myanimelist.net",
		url: "http://myanimelist.net/animelist/",
		urlSuffix: "&status=1",
		//ignoreAfter: 'class="header_completed"',
		animeRegEx: /<a href=["']\/anime\/([^\/]+)\/([^"']+)/g,
		progressRegEx: />([0-9-]+)[^0-9-]{1,10}([0-9-]+)</,
		animeImgRegEx: /<img src="(http:\/\/cdn.myanimelist.net\/images\/anime\/[^"]+)/,
		queryImage: function(entry, callBack) {
			var provider = animeListProviders["myanimelist.net"];
			var animeInfoRequest = new XMLHttpRequest();
			animeInfoRequest.open("GET", "http://myanimelist.net/anime/" + entry.id);
			animeInfoRequest.onload = function() {
				var html = this.responseText;

				// Get image
				var match = provider.animeImgRegEx.exec(html);
				if(match != null) {
					var coverUrl = match[1].replace("cdn.myanimelist.net", "myanimelist.net");
					callBack(coverUrl);
				}
			};
			animeInfoRequest.send(null);
		}
	},

	// Anilist
	"anilist.co": {
		host: "anilist.co",
		url: "http://anilist.co/animelist/",
		urlSuffix: "",
		ignoreAfter: "</table>",
		animeRegEx: /<a href=["']\/anime\/([^\/]+)\/[^"']+[^>]*>([^<]+)</g,
		progressRegEx: /class='cep[0-9]+[^0-9]+([0-9-]+)[^0-9-]{1,10}([0-9-]*)/,
		queryImage: function(entry, callBack) {
			callBack("http://anilist.co/img/dir/anime/reg/" + entry.id + ".jpg");
		}
	},

	// Hummingbird
	"hummingbird.me": {
		host: "hummingbird.me",
		url: "http://hummingbird.me/users/",
		apiUrl: "https://hummingbirdv1.p.mashape.com/users/",
		urlSuffix: "/library",
		apiUrlSuffix: "/library?status=currently-watching",
		beforeSend: function(req) {
			req.responseType = "json";
			req.setRequestHeader("X-Mashape-Key", "nr5IdgBU8pmshScE5qxAH92MmFwWp1oqx4mjsnA5igw5vcKlXu");
		},
		overrideParse: function(data) {
			return data.map(function(entry) {
				return {
					image: entry.anime.cover_image, // added here for efficiency
					id: entry.anime.id,
					title: entry.anime.title,
					originalTitle: entry.anime.title,
					watchedEpisodes: entry.episodes_watched,
					nextEpisodeToWatch: entry.episodes_watched + 1,
					maxEpisodes: entry.anime.episode_count,
					hasNewEpisodes: false,
					latestEpisodeNumber: -1,
					days: 0,
					hours: 0,
					minutes: 0
				};
			});
		},
		queryImage: function(entry, callBack) {
			callBack(entry.image);
		}
	}
};
