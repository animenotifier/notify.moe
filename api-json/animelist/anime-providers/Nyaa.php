<?php
class Nyaa implements AnimeProvider {
	// Special
	public $special;

	function __construct() {
		// Cached
		$cacheTime = 2 * 60;
		$key = "nyaa-exceptions";
		$json = apc_fetch($key, $found);

		if(!$found) {
			$json = getHTML("https://raw.githubusercontent.com/freezingwind/animereleasenotifier.com/master/api/providers/anime/Nyaa/special.json");

			// Cache it
			if(!empty($json))
				apc_add($key, $json, $cacheTime);
		}
		
		$this->special = json_decode($json, true);
	}

	public function getAnimeInfo($anime) {
		$searchTitle = $anime["title"];

		$searchTitle = preg_replace("/[^[:alnum:]!']/ui", ' ', $searchTitle);
		$searchTitle = str_replace(" TV", "", $searchTitle);
		$searchTitle = str_replace("  ", " ", $searchTitle);
		$searchTitle = trim($searchTitle, " ");

		$lookUpTitle = $searchTitle;
		
		if($this->special != null && array_key_exists($searchTitle, $this->special))
			$searchTitle = $this->special[$searchTitle];

		// Replace spaces with plus's
		$searchTitle = str_replace(" ", "+", $searchTitle);

		$quality = "";
		$subs = "";
		$nyaaSuffix = trim(str_replace("++", "+", "&cats=1_37&filter=0&sort=2&term=" . $searchTitle . "+" . $quality . "+" . $subs), "+");
		
		$url = "http://www.nyaa.se/?page=search" . $nyaaSuffix;
		$rssUrl = "http://www.nyaa.se/?page=rss" . $nyaaSuffix;

		$nyaa = [
			"url" => $url,
			"rssUrl" => $rssUrl,
			"nextEpisodeUrl" => $url . "+" . str_pad($anime["episodes"]["next"] + $anime["episodes"]["offset"], 2, "0", STR_PAD_LEFT),
			"videoUrl" => ""
		];

		// Cached XML data
		$cacheTime = 10 * 60;
		$key = $searchTitle . ":" .  $quality . ":" .  $subs . ":nyaa-xml";
		$xml = apc_fetch($key, $found);

		// Get HTML
		if(!$found) {
			$xml = getHTML($rssUrl);

			// Cache it
			apc_add($key, $xml, $cacheTime);
		}

		$anime["episodes"]["available"] = -1;
		$anime["animeProvider"] = $nyaa;

		// Error fetching XML
		if($xml === null || $xml === '')
			return $anime;

		$doc = simplexml_load_string($xml);

		// Error parsing XML
		if(!is_object($doc) || !is_object($doc->channel))
			return $anime;

		$items = $doc->channel->item;
		$episodeRegEx = "/[ _]?-?[ _]E?p?([0-9]{1,3})(v\d)?[ _][^a-zA-Z0-9-]/";
		//              .. Log Horizon 2 - 19 [720p].mkv
		$highestEpisodeAvailable = -1;

		foreach($items as $item) {
			$title = strval($item->title);

			if(preg_match($episodeRegEx, $title, $matches) === 1) {
				$episodeNumber = intval($matches[1]);

				if($episodeNumber > $highestEpisodeAvailable) {
					$highestEpisodeAvailable = $episodeNumber;
				}
			}
		}

		// On error
		if($highestEpisodeAvailable === -1) {
			$key = 'nyaa-errors';
			$jsonString = apc_fetch($key, $found);

			if(!$found)
				$jsonString = '{}';

			$json = json_decode($jsonString, true);

			if(array_key_exists($lookUpTitle, $json))
				$json[$lookUpTitle] = $json[$lookUpTitle] + 1;
			else
				$json[$lookUpTitle] = 1;
			
			apc_store($key, json_encode($json));

			// Slack message
			/*$key = "slack-nyaa:$searchTitle";
			$cacheTime = 24 * 60 * 60;

			apc_fetch($key, $found);

			if(!$found) {
				$message = "*Nyaa - Couldn't determine latest episode!*\n```\"" . $anime['title'] . "\":\n\"\",```\n" . $nyaa['url'];

				sendSlackMessage($message, 'arn');

				apc_add($key, 1, $cacheTime);
			}*/
		}

		$anime["episodes"]["available"] = $highestEpisodeAvailable;
		return $anime;
	}
}
?>