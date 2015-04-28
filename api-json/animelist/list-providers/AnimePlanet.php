<?php
class AnimePlanet implements ListProvider {
	public $specialOffsets = [
		"Aldnoah Zero (2015)" => 12
	];

	// Constructor
	function __construct($user) {
		// ...
	}

	// Get anime list URL
	public function getAnimeListUrl($userName) {
		return "http://www.anime-planet.com/users/$userName/anime/watching";
	}

	// Get anime list
	public function getAnimeList($userName, $completedOnly = false) {
		$requestedStatus = $completedOnly ? 'watched' : 'watching';

		// Cached XML data
		$cacheTime = 20 * 60;
		$key = $userName . ":animeplanet-html-$requestedStatus";
		$html = apc_fetch($key, $found);
		$animeRegEx =
		'/<tr>\s*<td class="tableTitle"><a title=[\'"]<h5><a href=[\'"](\/anime\/.*?)[\'"]>(.*?)<\/a><\/h5>.*?\((\d+)\+? eps\).*?src=[\'"]([^\'"]+)[\'"].*?<td class="tableEps">([^<]+)<\/td>.*?<\/td>\s*<\/tr>/';

		// Get HTML
		if(!$found) {
			$apiUrl = "http://www.anime-planet.com/users/$userName/anime/$requestedStatus";
			$html = getHTML($apiUrl);

			// Remove line breaks to enable full regex search
			$html = preg_replace('!\s+!m', ' ', $html);

			// Cache it
			apc_add($key, $html, $cacheTime);
		}

		$watching = array();
		preg_match_all($animeRegEx, $html, $matches, PREG_SET_ORDER);
		
		foreach($matches as $match) {
			$title = $match[2];

			// Offset
			$episodesOffset = 0;

			if(array_key_exists($title, $this->specialOffsets))
				$episodesOffset = $this->specialOffsets[$title];
			
			$episodesWatched = intval(trim($match[5]));
			$nextEpisodeToWatch = $episodesWatched + 1 + $episodesOffset;

			$newEntry = array(
				"title" => $title,
				"image" => 'http://www.anime-planet.com' . $match[4],
				"url" => 'http://www.anime-planet.com' . $match[1],
				'airingDate' => [
					'timeStamp' => '',
					'remaining' => '',
				],
				"episodes" => array(
					"watched" => $episodesWatched,
					"next" => $nextEpisodeToWatch,
					"available" => 0,
					"max" => intval(trim($match[3])),
					"offset" => $episodesOffset,
				)
			);

			$watching[] = $newEntry;
		}

		return $watching;
	}

	// Clear cache
	public function clearCache($userName) {
		$key = $userName . ":animeplanet-html";
		return apc_delete($key);
	}
}
?>