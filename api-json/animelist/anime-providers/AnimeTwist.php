<?php
class AnimeTwist implements AnimeProvider {
	public function getAnimeInfo($anime) {
		$searchTitle = $anime["title"];

		$searchTitle = preg_replace("/[^[:alnum:]]/ui", '', $searchTitle);
		$searchTitle = strtolower($searchTitle);

		// Special
		$special = [
			"aldnoahzero2" =>
			"aldnoahzero",

			"tokyoghoula" =>
			"tokyoghouls2",
		];

		if(array_key_exists($searchTitle, $special))
			$searchTitle = $special[$searchTitle];

		$nextEpisodeToWatch = $anime['episodes']['next'];
		$baseURL = 'https://twist.moe/a/';
		$searchURL = "$baseURL$searchTitle";
		$url = "$searchURL/1";
		$nextEpisodeURL = "$searchURL/$nextEpisodeToWatch";

		$animeProviderInfo = array(
			"url" => $url,
			"nextEpisodeUrl" => $nextEpisodeURL,
			"videoUrl" => ""
		);

		// Cached last episode available
		$cacheTime = 5 * 60;
		$key = $anime["title"] . ":animeTwist-episodes-available";
		$anime["episodes"]["available"] = apc_fetch($key, $found);

		if(!$found) {
			$html = getHTML($url);

			// Remove line breaks to enable full regex search
			$html = preg_replace('!\s+!m', ' ', $html);

			$episodeRegEx = "/twist.moe\/a\/[^\/]+\/(\d+)/";
			preg_match_all($episodeRegEx, $html, $matches, PREG_SET_ORDER);

			$latestEpisode = -1;

			foreach($matches as $match) {
				$latestEpisode = intval($match[1]);
			}

			$anime["episodes"]["available"] = $latestEpisode;

			// Cache it
			apc_store($key, $latestEpisode, $cacheTime);
		}
		
		$anime["animeProvider"] = $animeProviderInfo;
		return $anime;
	}
}
?>