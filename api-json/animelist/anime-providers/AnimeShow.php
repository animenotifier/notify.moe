<?php
class AnimeShow implements AnimeProvider {
	// Special
	public $special;
	
	// Constructor
	function __construct() {
		$this->special = json_decode(getHTML("https://raw.githubusercontent.com/freezingwind/animereleasenotifier.com/master/api/providers/anime/AnimeShow/special.json"), true);
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

			/*if($available === -1) {
				$nativeURL = $this->getLinkFromAnimeShow($animeTitle);
				$available = $this->getAvailableEpisode($nativeURL, $lookUpTitle);
			}*/

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
	private function getNativeTitleFromGoogle($title, $site) {
		global $config;
		
		$title = preg_replace('/[^[:alnum:]\'*]/ui', '+', $title);
		$customSearchEngineId = '002450170332278128138:nxs5wgw2vrg';
		$googleURL = "http://ajax.googleapis.com/ajax/services/search/web?v=1.0&key=" . $config['googleAPIKey'] . '$cx=' . $customSearchEngineId . "&userip=" . $_SERVER['REMOTE_ADDR'] . "&q=site:$site+" . $title;
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
			if(strpos($result['titleNoFormatting'], 'Episodes - ') !== false) {
				$url = $result['url'];
				return $url;
			}
		}

		return '';
	}
}
?>