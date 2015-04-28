<?php
class KissAnime implements AnimeProvider {
	// Special
	public $special;

	function __construct() {
		$this->special = json_decode(getHTML("https://raw.githubusercontent.com/freezingwind/animereleasenotifier.com/master/api/providers/anime/KissAnime/special.json"), true);
	}

	private function endsWith($haystack, $needle) {
		return substr($haystack, -strlen($needle)) === $needle;
	}

	public function getAvailableEpisode($kissAnimeURL, $lookUpTitle) {
		$html = getHTML($kissAnimeURL);
		// __cfduid=debda4a8d516aafa2f76f0322efa57b811427788208; test_enable_cookie=cookie_value; cf_clearance=a133541e9b79cdcd259408efda9bf0a2f94b3026-1427788216-604800

		if($html == '') {
			$key = 'kissanime-errors';
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

		if(preg_match("/ ([0-9]{3})<\/a>/", $html, $matches) === 1) {
			return intval($matches[1]);
		}

		return -1;
	}

	public function getAnimeInfo($anime) {
		$animeTitle = $anime['title'];
		$kaTitle = $animeTitle;

		$kaTitle = preg_replace('/[^[:alnum:]\'*]/ui', '-', $kaTitle);

		$kaTitle = str_replace('◎', '', $kaTitle);
		$kaTitle = str_replace('-TV', '', $kaTitle);
		$kaTitle = str_replace('--', '-', $kaTitle);
		$kaTitle = trim($kaTitle, '-');

		if($this->endsWith($kaTitle, '\'\''))
			$kaTitle = substr($kaTitle, 0, -2) . '-3';

		if($this->endsWith($kaTitle, '\''))
			$kaTitle = substr($kaTitle, 0, -1) . '-2';

		$lookUpTitle = $kaTitle;
		
		if($this->special != null && array_key_exists($kaTitle, $this->special))
			$kaTitle = $this->special[$kaTitle];

		// KissAnime URLs
		$nextEpisodeToWatch = $anime["episodes"]["next"];
		$kissAnimeURL = "http://kissanime.com/Anime/" . $kaTitle;
		
		$episodePadLength = 3;

		// Special case
		if($kaTitle === "Genshiken-2")
			$episodePadLength = 2;
		
		$kissAnimeNextEpisodeURL = $kissAnimeURL . "/Episode-" . substr("000" . $nextEpisodeToWatch, -$episodePadLength);

		$kissAnime = array(
			'url' => $kissAnimeURL,
			'nextEpisodeUrl' => $kissAnimeNextEpisodeURL,
			'videoUrl' => ''
		);

		// Cached last episode available
		$cacheTime = 10 * 60;
		$key = $anime["title"] . ":kissAnime-episodes-available";
		$available = apc_fetch($key, $found);

		if(!$found) {
			$available = $this->getAvailableEpisode($kissAnimeURL, $lookUpTitle);

			if($available === -1) {
				$kissAnimeURL = getLinkFromGoogle($animeTitle, 'kissanime.com', 'kissanime.com/Anime/');
				$available = $this->getAvailableEpisode($kissAnimeURL, $lookUpTitle);
			}

			// Cache it
			apc_add($key, $available, $cacheTime);
		}

		$anime['episodes']['available'] = $available;

		// On error
		//if($available === -1) {
			// Slack message
			/*$key = "slack-kissanime:$animeTitle";
			$cacheTime = 24 * 60 * 60;

			apc_fetch($key, $found);

			if(!$found) {
				$message = "*KissAnime - Couldn't determine latest episode!*\n```\"" . $kaTitle . "\":\n\"\",```\n" . $kissAnime['url'];

				sendSlackMessage($message, 'arn-bot-kissanime');

				apc_add($key, 1, $cacheTime);
			}*/
		//}

		//$anime["episodes"]["available"] = 1;

		// Cached video URL
		$cacheTime = 12 * 60 * 60;
		$key = $anime["title"] . "[" . $nextEpisodeToWatch . "]:kissAnime-video-url";
		$kissAnime["videoUrl"] = apc_fetch($key, $found);
		$kissAnime["videoHash"] = apc_fetch($key . '-hash', $foundHash);

		if(!$found) {
			$html = getHTML($kissAnimeNextEpisodeURL);

			// Remove line breaks to enable full regex search
			$html = preg_replace('!\s+!m', ' ', $html);

			/*if(preg_match("/redirector\.googlevideo\.com(.*?)(?=lh1)/", $html, $matches) === 1) {
				$videoURL = urldecode("https://redirector.googlevideo.com" . $matches[1] . "lh1");
				$kissAnime["videoHash"] = '';
				$kissAnime["videoUrl"] = $videoURL;
			} else*/
			if(preg_match('/<option\s+value="([^"]{100,})"/', $html, $matches) === 1) {
				$kissAnime["videoHash"] = $matches[1];
				$kissAnime["videoUrl"] = "";
			} else {
				$kissAnime["videoHash"] = '';
				$kissAnime["videoUrl"] = "";
			}

			// Cache it
			apc_add($key, $kissAnime["videoUrl"], $cacheTime);
			apc_add($key . '-hash', $kissAnime["videoHash"], $cacheTime);
		}

		if(!$foundHash)
			$kissAnime["videoHash"] = '';

		$anime["animeProvider"] = $kissAnime;
		return $anime;
	}
}
?>