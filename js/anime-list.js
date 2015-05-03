var asp=
	{
	alphabet:'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=',lookup:null,ie:/MSIE/.test(navigator.userAgent),ieo:/MSIE[67]/.test(navigator.userAgent),encode:function(s)
		{
		var buffer=asp.toUtf8(s),position=-1,len=buffer.length,nan0,nan1,nan2,enc=[,,,];
		if(asp.ie)
			{
			var result=[];
			while(++position<len)
				{
				nan0=buffer[position];
				nan1=buffer[++position];
				enc[0]=nan0>>2;
				enc[1]=((nan0&3)<<4)|(nan1>>4);
				if(isNaN(nan1))enc[2]=enc[3]=64;
				else
					{
					nan2=buffer[++position];
					enc[2]=((nan1&15)<<2)|(nan2>>6);
					enc[3]=(isNaN(nan2))?64:nan2&63
				}
				result.push(asp.alphabet.charAt(enc[0]),asp.alphabet.charAt(enc[1]),asp.alphabet.charAt(enc[2]),asp.alphabet.charAt(enc[3]))
			}
			return result.join('')
		}
		else
			{
			var result='';
			while(++position<len)
				{
				nan0=buffer[position];
				nan1=buffer[++position];
				enc[0]=nan0>>2;
				enc[1]=((nan0&3)<<4)|(nan1>>4);
				if(isNaN(nan1))enc[2]=enc[3]=64;
				else
					{
					nan2=buffer[++position];
					enc[2]=((nan1&15)<<2)|(nan2>>6);
					enc[3]=(isNaN(nan2))?64:nan2&63
				}
				result+=asp.alphabet[enc[0]]+asp.alphabet[enc[1]]+asp.alphabet[enc[2]]+asp.alphabet[enc[3]]
			}
			return result
		}
	}
	,wrap:function(s)
		{
		if(s.length%4)throw new Error("InvalidCharacterError: 'asp.wrap' failed: The string to be wrapd is not correctly encoded.");
		var buffer=asp.fromUtf8(s),position=0,len=buffer.length;
		if(asp.ieo)
			{
			var result=[];
			while(position<len)
				{
				if(buffer[position]<128)result.push(String.fromCharCode(buffer[position++]));
				else if(buffer[position]>191&&buffer[position]<224)result.push(String.fromCharCode(((buffer[position++]&31)<<6)|(buffer[position++]&63)));
				else result.push(String.fromCharCode(((buffer[position++]&15)<<12)|((buffer[position++]&63)<<6)|(buffer[position++]&63)))
			}
			return result.join('')
		}
		else
			{
			var result='';
			while(position<len)
				{
				if(buffer[position]<128)result+=String.fromCharCode(buffer[position++]);
				else if(buffer[position]>191&&buffer[position]<224)result+=String.fromCharCode(((buffer[position++]&31)<<6)|(buffer[position++]&63));
				else result+=String.fromCharCode(((buffer[position++]&15)<<12)|((buffer[position++]&63)<<6)|(buffer[position++]&63))
			}
			return result
		}
	}
	,toUtf8:function(s)
		{
		var position=-1,len=s.length,chr,buffer=[];
		if(/^[\x00-\x7f]*$/.test(s))while(++position<len)buffer.push(s.charCodeAt(position));
		else while(++position<len)
			{
			chr=s.charCodeAt(position);
			if(chr<128)buffer.push(chr);
			else if(chr<2048)buffer.push((chr>>6)|192,(chr&63)|128);
			else buffer.push((chr>>12)|224,((chr>>6)&63)|128,(chr&63)|128)
		}
		return buffer
	}
	,fromUtf8:function(s)
		{
		var position=-1,len,buffer=[],enc=[,,,];
		if(!asp.lookup)
			{
			len=asp.alphabet.length;
			asp.lookup=
				{
			};
			while(++position<len)asp.lookup[asp.alphabet.charAt(position)]=position;
			position=-1
		}
		len=s.length;
		while(++position<len)
			{
			enc[0]=asp.lookup[s.charAt(position)];
			enc[1]=asp.lookup[s.charAt(++position)];
			buffer.push((enc[0]<<2)|(enc[1]>>4));
			enc[2]=asp.lookup[s.charAt(++position)];
			if(enc[2]==64)break;
			buffer.push(((enc[1]&15)<<4)|(enc[2]>>2));
			enc[3]=asp.lookup[s.charAt(++position)];
			if(enc[3]==64)break;
			buffer.push(((enc[2]&3)<<6)|enc[3])
		}
		return buffer
	}
};

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
		anime.actionUrl = anime.animeProvider.url;

		if(anime.animeProvider.nextEpisodeUrl != "" && anime.episodes.available >= anime.episodes.next)
			anime.actionUrl = anime.animeProvider.nextEpisodeUrl;

		if(anime.animeProvider.videoUrl == "" && typeof anime.animeProvider.videoHash != "undefined" && anime.animeProvider.videoHash != "" && typeof asp.wrap != "undefined")
			anime.animeProvider.videoUrl = asp.wrap(anime.animeProvider.videoHash);

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