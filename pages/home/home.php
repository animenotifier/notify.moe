<style scoped>
	section {
		position: relative;
		left: 50%;
		transform: translateX(-50%);
		float: left;
		display: block;
		clear: both;
		transform-style: preserve-3d;
		/*border: 1px solid red;*/
	}

	section header {
		width: 100%;
		clear: both;
	}

	section h2 {
		text-align: center;
		font-size: 2em;
	}

	.button-list {
		float: left;
		display: block;
		margin: 0;
		/*border: 1px dashed yellow;*/

		position: relative;
		left: 50%;
		transform: translateX(-50%);
	}

	.button-list li {
		float: left;
		display: block;
	}

	.button-link {
		float: left;
		display: inline-block;

		padding: 1em;
		margin: 0.5em;
		/*background-color: rgba(8, 8, 8, 0.02);*/
		border-radius: 5px;
		width: 256px;
		height: 256px;
		text-align: center;
		transform: scale(0.7);
	}

	.button-link:hover {
		width: 256px;
		height: 256px;
		text-decoration: none;
		/*box-shadow: 0 0 8px rgba(0, 0, 0, 0.5);*/
		transform: scale(1);
	}

	.button-content {
		/*position: relative;
		top: 50%;
		transform: translateY(-50%);*/
	}

	.button-content span[itemprop="name"] {
		display: none;
	}

	.button-content h3 {
		font-size: 1em;
		font-weight: normal;
		margin-bottom: 1em !important;
		text-align: center;
	}

	.no-scale {
		transform: scale(1.0);
	}
</style>

<section>
	<header>
		<h2>Pick a platform</h2>
	</header>

	<div class="button-list">
		<a href="javascript:chrome.webstore.install();" class="button-link" title="Chrome" itemscope itemtype='http://schema.org/Product'>
			<div class="button-content">
				<span itemprop="name">Anime Release Notifier - Chrome Extension</span>
				<img itemprop="image" src="/images/platforms/chrome.png" alt="Chrome" width="256" height="256">
			</div>
		</a>
		<a href="https://addons.mozilla.org/en-US/firefox/addon/anime-release-notifier/" target="_blank" class="button-link" title="Firefox" itemscope itemtype='http://schema.org/Product'>
			<div class="button-content">
				<span itemprop="name">Anime Release Notifier - Firefox Extension</span>
				<img itemprop="image" src="/images/platforms/firefox.png" alt="Firefox" width="256" height="256">
			</div>
		</a>
		<a href="https://play.google.com/store/apps/details?id=com.freezingwind.animereleasenotifier" target="_blank" class="button-link" title="Android" itemscope itemtype='http://schema.org/Product'>
			<div class="button-content">
				<span itemprop="name">Anime Release Notifier - Android App</span>
				<img itemprop="image" src="/images/platforms/android.png" alt="Android" width="256" height="256">
			</div>
		</a>
		<a href="/pc" class="ajax button-link" title="PC" width="256" height="256" itemscope itemtype='http://schema.org/Product'>
			<div class="button-content">
				<span itemprop="name">Anime Release Notifier - Web Interface</span>
				<img itemprop="image" src="/images/platforms/pc.png" alt="Web">
			</div>
		</a>
	</div>
</section>

<section>
	<header>
		<h2>How it works</h2>
	</header>

	<!--<div class="button-link no-scale">
		<div class="button-content">
			Sends your anime list through a banana-powered microwave to watch anime.
		</div>
	</div>-->

	<div class="button-link no-scale">
		<div class="button-content">
			<h3>Get your watching list:</h3>

			<ul>
				<li><a href="http://anilist.co" target="_blank">anilist.co</a></li>
				<li><a href="http://anime-planet.com" target="_blank">anime-planet.com</a></li>
				<li><a href="http://hummingbird.me" target="_blank">hummingbird.me</a></li>
				<li><a href="http://myanimelist.net" target="_blank">myanimelist.net</a></li>
			</ul>
		</div>
	</div>

	<!--<div class="button-link no-scale">
		<div class="button-content">
			<h3>Let ARN build a unified format:</h3>

			<pre style="font-size: 0.7em; word-wrap: break-word; line-height: 1.3em; tab-size: 0;">{
"name": "Test User",
"listUrl": "http://hummingbird.me/
users/Tester/library",
"watching": [...]
}</pre>
		</div>
	</div>-->

	<div class="button-link no-scale">
		<div class="button-content">
			<h3>Connect it with:</h3>

			<ul>
				<li style="text-decoration: none"><a href="http://nyaa.se" target="_blank" rel="nofollow">nyaa.se</a></li>
				<li style="text-decoration: line-through"><a href="http://kissanime.com" target="_blank" rel="nofollow">kissanime.com</a></li>
				<li style="text-decoration: line-through"><a href="http://animeshow.tv" target="_blank" rel="nofollow">animeshow.tv</a></li>
				<li style="text-decoration: line-through"><a href="http://twist.moe" target="_blank" rel="nofollow">twist.moe</a></li>
			</ul>
		</div>
	</div>
</section>

<!--<section>
	<header>
		<h2 title="for those hundreds of hours of unpaid work :(">Buy me a drink</h2>
	</header>

	<div class="button-list">
		<a href="http://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=DADU374FK8X2J&lc=US" target="_blank" class="button-link" title="PayPal" itemscope itemtype='http://schema.org/PaymentMethod'>
			<div class="button-content">
				<img itemprop="image" src="/images/platforms/paypal.png" alt="PayPal" width="256" height="256">
			</div>
		</a>

		<a href="bitcoin:1NSDrUSUZYWki4XRXN8gZZXahNx6HZbHWK" class="button-link" title="BitCoin" itemscope itemtype='http://schema.org/PaymentMethod'>
			<div class="button-content">
				<img itemprop="image" src="/images/platforms/bitcoin.png" alt="BitCoin" width="256" height="256">
			</div>
		</a>
	</div>
</section>-->