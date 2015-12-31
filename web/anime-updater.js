var animeUpdater = {
	// Request anime list
	requestAnimeList: function() {
		var $animeList = $("#anime-list");
		var $settings = $("#settings");

		var userName = localStorage.userName;

		if(!userName || $.type(userName) !== "string") {
			$animeList.hide();
			$settings.show();
			return;
		}

		console.log(userName);

		$settings.hide();
		$animeList.show();

		$.getJSON("https://animereleasenotifier.com/api/animelist/" + userName, function(json) {
			var animeList = new AnimeList(json, $animeList, 1, function(anime) {
				anime.sendNotification();
			});

			// Footer
			this.buildFooter(this.getProfileUrl(userName), animeList.listUrl);
		}.bind(this));
	},

	// Build footer
	buildFooter: function(profileUrl, listUrl) {
		var userName = localStorage.userName;

		// Create footer
		var footer = document.createElement("div");
		footer.className = "footer";

		$(footer).html(
			"<a href='" + profileUrl + "' target='_blank' title='Profile'>" + userName + "</a> | " +
			"<a href='" + listUrl + "' target='_blank' title='Anime List'>Edit List</a>" +
			" <a href='http://anichart.net/airing' target='_blank' title='Chart'><img src='https://animereleasenotifier.com/images/icons/chart.png' alt='Chart'/></a>" +
			" <a href='javascript:toggleSettings();' title='Options'><img src='https://animereleasenotifier.com/images/icons/settings.png' alt='Options'/></a>"
		);

		$("#anime-list").append(footer);
	},

	getProfileUrl: function(userName) {
		return 'https://animereleasenotifier.com/user/' + userName;
	}
};