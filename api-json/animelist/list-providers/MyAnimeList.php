<?php
class MyAnimeList implements ListProvider {
	public $specialOffsets = [
		'Kuroko no Basket 3rd Season' => 50
	];

	// Constructor
	function __construct($user) {
		// ...
	}

	// Get anime list URL
	public function getAnimeListUrl($userName) {
		return "http://myanimelist.net/animelist/$userName&status=1";
	}

	public function getMALBlockList() {
		return array(array(
			"title" => 'MAL is blocking ARN from accessing your list.',
			"image" => '',
			"url" => "https://animereleasenotifier.com/incapsula",
			'airingDate' => [
				'timeStamp' => '',
				'remaining' => '',
			],
			"episodes" => array(
				"watched" => 1,
				"next" => 2,
				"available" => 1,
				"max" => 1,
				"offset" => 0,
			)
		));
	}

	// Get anime list
	public function getAnimeList($userName, $completedOnly = false) {
		$requestedStatus = $completedOnly ? 2 : 1;

		// Cached XML data
		$cacheTime = 240 * 60;
		$key = $userName . ":myanimelist-xml";
		$xml = apc_fetch($key, $found);

		// Get HTML
		if(!$found) {
			$apiUrl = "http://myanimelist.net/malappinfo.php?u=$userName&status=all&type=anime";
			$cookies = "";
			$xml = getHTML($apiUrl, $cookies);

			// Cache it
			apc_add($key, $xml, $cacheTime);
		}

		// Incapsula error
		if(strpos($xml, 'Incapsula') !== FALSE)
			return $this->getMALBlockList();

		$doc = simplexml_load_string($xml);

		// Parsing error
		if(!is_object($doc))
			return $this->getMALBlockList();

		$animeList = $doc->anime;
		$watching = array();

		foreach($animeList as $anime) {
			if($anime->my_status != $requestedStatus)
				continue;

			$title = strval($anime->series_title);

			$episodesOffset = 0;

			if(array_key_exists($title, $this->specialOffsets))
				$episodesOffset = $this->specialOffsets[$title];

			$episodesWatched = intval($anime->my_watched_episodes);
			$nextEpisodeToWatch = $episodesWatched + 1 + $episodesOffset;

			$newEntry = array(
				"title" => $title,
				"image" => strval($anime->series_image),
				"url" => "",
				'airingDate' => [
					'timeStamp' => '',
					'remaining' => '',
				],
				"episodes" => array(
					"watched" => $episodesWatched,
					"next" => $nextEpisodeToWatch,
					"available" => 0,
					"max" => $anime->series_episodes ? intval($anime->series_episodes) : -1,
					"offset" => $episodesOffset,
				)
			);

			$watching[] = $newEntry;
		}

		return $watching;
	}

	// Clear cache
	public function clearCache($userName) {
		$key = $userName . ":myanimelist-xml";
		return apc_delete($key);
	}
}
?>