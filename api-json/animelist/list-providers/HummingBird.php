<?php
class HummingBird implements ListProvider {
	public $specialOffsets = [
		'Aldnoah.Zero 2' => 12,
		'JoJo\'s Bizarre Adventure: Stardust Crusaders - Egypt Arc' => 24,
		'Seitokai Yakuindomo* OVA' => 13,
	];

	// Constructor
	function __construct($user) {
		// ...
	}

	// Get anime list URL
	public function getAnimeListUrl($userName) {
		return "http://hummingbird.me/users/$userName/library";
	}

	// Get anime list
	public function getAnimeList($userName, $completedOnly = false) {
		$requestedStatus = $completedOnly ? "completed" : "currently-watching";

		// Cached JSON data
		$cacheTime = 20 * 60;
		$key = $userName . ":hummingbird-json-$requestedStatus";
		$json = apc_fetch($key, $found);

		// API URL
		if(!$found) {
			$apiUrl = "https://hummingbirdv1.p.mashape.com/users/{userName}/library?status=$requestedStatus";
			$apiUrl = str_replace("{userName}", $userName, $apiUrl);

			// Execute
			$json = getJSON($apiUrl, 'X-Mashape-Key: nr5IdgBU8pmshScE5qxAH92MmFwWp1oqx4mjsnA5igw5vcKlXu');

			// Cache it
			apc_add($key, $json, $cacheTime);
		}
		
		// Parse JSON
		$data = json_decode($json, true);

		// Error parsing JSON?
		if($data === null || count($data) === 0) {
			return array();
		}

		// Watching list
		$watching = array();
		foreach($data as $entry) {
			if($entry['status'] != $requestedStatus)
				continue;

			$anime = $entry['anime'];
			$title = $anime['title'];

			// Offset
			$episodesOffset = 0;

			if(array_key_exists($title, $this->specialOffsets))
				$episodesOffset = $this->specialOffsets[$title];

			$episodesWatched = $entry['episodes_watched'];
			$nextEpisodeToWatch = $episodesWatched + 1 + $episodesOffset;

			$newEntry = array(
				'title' => $title,
				'image' => $anime['cover_image'],
				'url' => $anime['url'],
				'airingDate' => [
					'timeStamp' => '',
					'remaining' => '',
				],
				'episodes' => array(
					'watched' => $episodesWatched,
					'next' => $nextEpisodeToWatch,
					'available' => 0,
					'max' => $anime['episode_count'] ? $anime['episode_count'] : -1,
					'offset' => $episodesOffset,
				)
			);

			$watching[] = $newEntry;
		}

		return $watching;
	}

	// Clear cache
	public function clearCache($userName) {
		$key = $userName . ":hummingbird-json";
		return apc_delete($key);
	}
}
?>