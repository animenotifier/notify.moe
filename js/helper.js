// Get URLs
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

// Set object (store objects in localStorage)
Storage.prototype.setObject = function(key, value) {
	this.setItem(key, JSON.stringify(value));
}

// Get object (retrieve objects from localStorage)
Storage.prototype.getObject = function(key) {
	var value = this.getItem(key);
	return value && JSON.parse(value);
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