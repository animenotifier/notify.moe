<style scoped>
	.good {
		color: rgb(64, 255, 64);
	}

	.mal-problem {
		max-width: 1000px;
		margin: 0 auto;
	}
</style>

<section class="mal-problem">
	<h2>Explanation of the problem with MyAnimeList</h2>
	<p>
		myanimelist.net is currently blocking the access from the ARN servers because they are using a software called Incapsula.
		Incapsula falsely thinks the ARN server is trying to DoS them due to the high number of requests from all MAL users on ARN.
		I have contacted MAL in multiple ways now to request being whitelisted (that means ARN can access MAL again) but I still didn't receive any reply.
	</p>
	<p>
		Please note that this is not a mistake in the ARN code. I didn't change anything in the code of the ARN, it still works perfectly fine with MAL.
		However I can't influence the fact that they are blocking us right now. It is out my hands and all I can do is wait for one of my messages to be heard.
	</p>

	<h2>Moving</h2>

	<p>
		Due to MAL's blocking, some users decided to move to another list provider. If that is an option for you, please have a look at this objective comparison:
	</p>

	<table>
		<thead>
			<tr>
				<th>Provider</th>
				<th>MAL importer</th>
				<th>JSON API</th>
				<th>Image quality</th>
				<th>Average response time</th>
			</tr>
		</thead>

		<tbody>
			<tr>
				<td>myanimelist.net</td>
				<td>-</td>
				<td>No</td>
				<td>≈ 320px</td>
				<td>≈ 0.5 s</td>
			</tr>

			<tr>
				<td>anime-planet.com</td>
				<td>No</td>
				<td>No</td>
				<td>≈ 550px</td>
				<td>≈ 1.1 s</td>
			</tr>

			<tr>
				<td>anilist.co</td>
				<td class="good">Yes</td>
				<td class="good">Yes</td>
				<td>≈ 320px</td>
				<td class="good">≈ 0.3 s</td>
			</tr>

			<tr>
				<td>hummingbird.me</td>
				<td class="good">Yes</td>
				<td class="good">Yes</td>
				<td class="good">≈ 710px</td>
				<td class="good">≈ 0.3 s</td>
			</tr>
		</tbody>
	</table>

	<h2>Update 7th March, 2015</h2>

	<p>
		Since I didn't get any reply from the owner of MAL I will try to come up with a workaround.
		This will cost some more money for the servers but if MAL will work again it is well worth it (even though I <em>personally</em> dislike it).
		Sadly I haven't been getting any <a href="/contributors" class="ajax">donations</a> recently and I'm not rich either.
		If you want this project to continue being free of charge and also be included in the upcoming donator list with a link to your profile please consider making a donation.
		Apart from donations I am not making any money with this project, in fact it's actually too little to keep it up. I'm in the minus.
	</p>

	<h2>Update 9th March, 2015</h2>

	<p>
		I published version 1.0.1 and enforced an "undeletable" 30 min. cache on MAL requests.
		"Undeletable" cache means that even if you press the button for the chrome extension and it will not force a list update.
		However it seems Incapsula doesn't just block you based on number of requests.
		I bought a second server to access the MAL data from the fresh server with a new IP, however the new server got blocked the instant it made a 2nd data request.
		This is a serious problem with MAL + Incapsula. Data access is not possible from the server side.
		I can't believe that nobody noticed this yet, this problem needs to be addressed.
		However the only one who can do something about it (the owner of MAL) still didn't send me a reply after a week has passed.
	</p>

	<h2>Update 11th March, 2015</h2>

	<p>
		<img src="/images/screenshots/xinil-1.png" alt="Response">
	</p>

	<h2>Update 21th March, 2015</h2>

	<p>
		<img src="http://puu.sh/gJ9N4/83c3b93353.png" alt="Response">
	</p>

	<h2>Update 22th March, 2015</h2>

	<p>
		<a href="https://plus.google.com/106001822498547702835/posts/7bJ497QZxvQ"></a>
	</p>
</section>

<!--<h2>Update 8th March, 2015</h2>

<p>I bought a second server now. It might work again, please try it and let me know if it works again via the feedback page.</p>
<p>However I have set up a 30 minute cache to be 200% safe we don't get blocked again. This cache can't be cleared. That means MAL users will have to wait 30 minutes before their list can get updated. Other providers are not affected. Obviously this is not the best solution but if it works it's better than nothing. In the future the cache will be less strict.</p>-->