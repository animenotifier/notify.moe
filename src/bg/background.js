// if you checked "fancy-settings" in extensionizr.com, uncomment this lines

var settings = new Store("settings", {
	"userName": "",
	"quality": "",
	"subs": "",
	"otherSearch": "",
	"updateInterval": "5",
	"maxEpisodeDifference": "1",
	"animeProvider": "nyaa.se",
	"animeListProvider": "anilist.co"
});

//example of using a message handler from the inject scripts
chrome.runtime.onMessage.addListener(function(request, sender, sendResponse) {
	//chrome.pageAction.show(sender.tab.id);
	sendResponse(settings.toObject());
});

chrome.notifications.onButtonClicked.addListener(function(notificationId, btnIdx) {
	if(btnIdx === 0) {
		var link = animeUpdater.notificationIdToLink[notificationId];
		window.open(link);
	}
});

var backgroundUpdate = function() {
	//localStorage["anime.Black Bullet"] = {latestEpisodeNumber: 0};
	animeUpdater.requestAnimeList(settings.toObject(), function() {
		
	});
};

document.addEventListener('DOMContentLoaded', function() {
	backgroundUpdate();

	// Get update interval
	var updateInterval = parseInt(settings.toObject()["updateInterval"]);
	if(isNaN(updateInterval) || updateInterval < 1)
		updateInterval = 5;

	// Update consistently
	console.log("Interval: " + updateInterval);
	window.setInterval(function() {
		console.log("Updating anime list in the background...");
		backgroundUpdate();
	}, 60 * 1000 * updateInterval);
});

//;

/*animeUpdater.animeList.forEach(function(entry) {
			console.log(entry.title + " > " + entry.latestEpisodeNumber);
		});*/

		/*chrome.notifications.create("", {
			type: "list",
			title: "New anime episodes",
			message: "",
			iconUrl: "/icons/icon128.png",
			items: [{ title: entry.title, message: entry.latestEpisodeNumber}]
		}, function(notificationId) {
			// ...
		});*/