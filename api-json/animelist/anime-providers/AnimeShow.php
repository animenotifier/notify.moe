<?php
class AnimeShow implements AnimeProvider {
	// Special
	public $special;
	
	// Constructor
	function __construct() {
		// Cached
		$cacheTime = 2 * 60;
		$key = "animeShow-exceptions";
		$json = apc_fetch($key, $found);

		if(!$found) {
			$json = getHTML("https://raw.githubusercontent.com/freezingwind/animereleasenotifier.com/master/api/providers/anime/AnimeShow/special.json");

			// Cache it
			if(!empty($json))
				apc_add($key, $json, $cacheTime);
		}

		$this->special = json_decode($json, true);
	}
	
	// Get available episode
	public function getAvailableEpisode($nativeURL, $lookUpTitle) {
		$html = getHTML($nativeURL);

		if($html == '') {
			$key = 'animeshow-errors';
			$jsonString = apc_fetch($key, $found);

			if(!$found)
				$jsonString = '{}';

			$json = json_decode($jsonString, true);

			if(array_key_exists($lookUpTitle, $json))
				$json[$lookUpTitle] = $json[$lookUpTitle] + 1;
			else
				$json[$lookUpTitle] = 1;
			
			apc_store($key, json_encode($json));
		}

		if(preg_match("/Episode ([0-9]{1,3})<\/h2>/", $html, $matches) === 1) {
			return intval($matches[1]);
		}
		
		return -1;
	}
	
	// Get anime info
	public function getAnimeInfo($anime) {
		$animeTitle = $anime['title'];
		$nativeTitle = $animeTitle;

		$nativeTitle = preg_replace('/[^[:alnum:]\'*]/ui', '-', $nativeTitle);
		
		$nativeTitle = str_replace('--', '-', $nativeTitle);
		$nativeTitle = trim($nativeTitle, '-');

		$lookUpTitle = $nativeTitle;
		
		if($this->special != null && array_key_exists($nativeTitle, $this->special))
			$nativeTitle = $this->special[$nativeTitle];

		// Doesn't exist on the platform?
		if(empty($nativeTitle)) {
			$anime['episodes']['available'] = -1;
			$anime["animeProvider"] = array(
				'url' => '',
				'nextEpisodeUrl' => '',
				'videoUrl' => ''
			);
			
			return $anime;
		}

		// Native URLs
		$nextEpisodeToWatch = $anime['episodes']['next'];
		$nativeURL = 'http://animeshow.tv/' . $nativeTitle . '/';
		$nativeNextEpisodeURL = "http://animeshow.tv/$nativeTitle-episode-$nextEpisodeToWatch/";

		// Cached last episode available
		$cacheTime = 10 * 60;
		$key = $anime["title"] . ":animeShow-episodes-available";
		$available = apc_fetch($key, $found);

		if(!$found) {
			$available = $this->getAvailableEpisode($nativeURL, $lookUpTitle);

			// Google fetch
			if($available === -1 && !empty($this->special)) {
				$googleCacheTime = 12 * 60 * 60;
				$googleKey = $animeTitle . ":animeShow-google-result";
				$nativeTitle = apc_fetch($key, $googleCached);

				if(!$googleCached) {
					$nativeTitle = $this->getNativeTitleFromGoogle($animeTitle);

					// Cache it
					apc_add($googleKey, $nativeTitle, $googleCacheTime);
				}

				if(!empty($nativeTitle)) {
					$nativeURL = 'http://animeshow.tv/' . $nativeTitle . '/';
					$nativeNextEpisodeURL = "http://animeshow.tv/$nativeTitle-episode-$nextEpisodeToWatch/";
					$available = $this->getAvailableEpisode($nativeURL, $lookUpTitle);

					if($available !== -1) {
						sendSlackMessage("[A] $animeTitle\n[U] $nativeURL\n[E] https://github.com/freezingwind/animereleasenotifier.com/edit/master/api/providers/anime/AnimeShow/special.json\n```\"$lookUpTitle\":\n\"$nativeTitle\",```", 'animeshow');
					}
				}
			}

			// Cache it
			apc_add($key, $available, $cacheTime);
		}

		$anime['episodes']['available'] = $available;
		$anime['episodes']['offset'] = 0; // TEMPORARY WORKAROUND
		$anime["animeProvider"] = array(
			'url' => $nativeURL,
			'nextEpisodeUrl' => $nativeNextEpisodeURL,
			'videoUrl' => ''
		);
		
		return $anime;
	}

	/*// Get link from AnimeShow
	private function getLinkFromAnimeShow($title) {
		//$title = preg_replace('/[^[:alnum:]\'*]/ui', '+', $title);
		$apiURL = 'http://animeshow.tv/pages/search-data.php';

		// Execute curl request
		$c = curl_init($apiURL);
		curl_setopt($c, CURLOPT_POST, 3);
		curl_setopt($c, CURLOPT_POSTFIELDS, "search=$title");
		curl_setopt($c, CURLOPT_RETURNTRANSFER, true);
		curl_setopt($c, CURLOPT_CONNECTTIMEOUT, 4);
		curl_setopt($c, CURLOPT_TIMEOUT, 4);
		$html = curl_exec($c);

		if(preg_match('/href="([^"]+)"/', $html, $matches) === 1) {
			return $matches[1];
		} else {
			return null;
		}
	}*/
	
	// Get native title from Google
	private function getNativeTitleFromGoogle($title) {
		global $config;
		
		$title = preg_replace('/[^[:alnum:]\'*]/ui', '+', $title);
		$customSearchEngineId = '002450170332278128138:nxs5wgw2vrg';
		$googleURL = "http://ajax.googleapis.com/ajax/services/search/web?v=1.0&key=" . $config['googleAPIKey'] . '&cx=' . $customSearchEngineId . "&userip=" . $_SERVER['REMOTE_ADDR'] . "&q=site:animeshow.tv+" . $title;
		$googleResults = getHTML($googleURL);

		// Parse JSON
		$googleData = json_decode($googleResults, true);
		$responseData = $googleData['responseData'];
		
		// Quota exceeded
		if($responseData === null)
			return '';
		
		$results = $responseData['results'];

		if(count($results) === 0)
			return '';

		foreach($results as $result) {
			$url = $result['url'];

			if(preg_match("/animeshow.tv\/([-\p{L}\d]+)-episode-(\d{1,3})/", $url, $matches) === 1 || preg_match("/animeshow.tv\/([-\p{L}\d]+)\//", $url, $matches) === 1) {
				// If we find a 'genre' link that probably means the anime doesn't exist so we'll give up
				if(substr($matches[1], 0, 5) === 'genre')
					return '';

				return $matches[1];
			}
		}

		return '';
	}
}
?>