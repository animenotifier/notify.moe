<?php
class AniList implements ListProvider, TimeProvider {
	private $accessToken;
	private $now;

	public $specialOffsets = [
		'Aldnoah.Zero 2' => 12,
		'Fate/stay night: Unlimited Blade Works 2nd Season' => 12,
		'Saenai Heroine no Sodate-kata' => -1
	];

	// Constructor
	function __construct($user) {
		$this->now = new DateTime();

		// Cached access token
		$cacheTime = 30 * 60;
		$key = 'anilist-access-token'; //$user['userName'] . ":anilist-access-token";
		$this->accessToken = apc_fetch($key, $found);

		if(!$found) {
			$authUrl = "https://anilist.co/api/auth/access_token";

			// Execute curl request
			$c = curl_init($authUrl);
			curl_setopt($c, CURLOPT_POST, 3);

			//if(array_key_exists('auth', $user['animeLists']['AniList']))
			//	$auth = $user['animeLists']['AniList']['auth'];

			/*if(isset($auth) && $auth !== null && $auth !== '')
				curl_setopt($c, CURLOPT_POSTFIELDS, "grant_type=authorization_pin&client_id=akyoto-wbdln&client_secret=zS3MidMPmolyHRYNOvSR1&code=$auth");
			else*/
				curl_setopt($c, CURLOPT_POSTFIELDS, "grant_type=client_credentials&client_id=akyoto-wbdln&client_secret=zS3MidMPmolyHRYNOvSR1");
			
			curl_setopt($c, CURLOPT_RETURNTRANSFER, true);
			curl_setopt($c, CURLOPT_CONNECTTIMEOUT, 4);
			curl_setopt($c, CURLOPT_TIMEOUT, 20);
			$authJson = curl_exec($c);

			$authData = json_decode($authJson);
			$this->accessToken = $authData->access_token;

			// Cache it
			apc_add($key, $this->accessToken, $cacheTime);
		}
	}

	// Get airing date
	public function getAiringDate($anime) {
		global $listProviderName;
		$nextEpisode = $anime['episodes']['next'] - $anime['episodes']['offset'];
		$remaining = '';

		// Shortcut if AniList is also the list provider
		if($listProviderName == 'AniList') {
			$timeStamp = $this->getAiringDateForAnimeID($anime['id'], $nextEpisode);
		} else {
			$animeTitle = $anime['title'];
			$animeTitle = '"' . str_replace('/', ' ', $animeTitle) . '"';
			$animeTitle = urlencode($animeTitle);

			// Cached
			$cacheTime = 12 * 60 * 60;
			$key = "anilist:$animeTitle:search-results-json";
			$json = apc_fetch($key, $found);

			if(!$found) {
				$searchURL = "https://anilist.co/api/anime/search/$animeTitle";
				$json = getJSON($searchURL . "?access_token=$this->accessToken");

				// Cache it
				apc_add($key, $json, $cacheTime);
			}
			
			$searchResults = json_decode($json);

			if(count($searchResults) == 0) {
				$timeStamp = -1;
			} else if(is_object($searchResults)) {
				$timeStamp = -1;//$this->getAiringDateForAnimeID($searchResults->id, $nextEpisode);
			} else {
				$timeStamp = $this->getAiringDateForAnimeID($searchResults[0]->id, $nextEpisode);
			}
		}

		// MAL block message
		if($anime['url'] === 'https://animereleasenotifier.com/incapsula') {
			$timeStamp = 1425531315;
		}
		
		$anime['airingDate']['timeStamp'] = $timeStamp;

		if($timeStamp == -1) {
			$anime['airingDate']['remaining'] = '';
			$anime['airingDate']['remainingString'] = '';
			return $anime;
		}

		$airingDate = (new DateTime())->setTimestamp($timeStamp);
		$diff = $this->now->diff($airingDate);

		$seconds = $diff->s;
		$minutes = $diff->i;
		$hours = $diff->h;
		$days = $diff->days;
		// TODO: daysRounded

		if($days == 0) {
			if($hours == 0) {
				if($minutes != 0)
					$remaining = plural($minutes, "minute");
				else
					$remaining = plural($seconds, "second");
			} else {
				$remaining = plural($hours, "hour");
			}
		} else {
			$remaining = plural($days, "day");
		}

		if($airingDate < $this->now) {
			$remaining .= ' ago';
		}

		$anime['airingDate']['remaining'] = $remaining;
		$anime['airingDate']['remainingString'] = $remaining;

		return $anime;
	}

	// Get airing date by anime ID
	function getAiringDateForAnimeID($animeId, $nextEpisode) {
		// Cached
		$cacheTime = 12 * 60 * 60;
		$key = "anilist:$animeId:$nextEpisode:airing-date";
		$timeStamp = apc_fetch($key, $found);

		if(!$found) {
			$airingDateURL = "https://anilist.co/api/anime/$animeId/airing";
			$json = getJSON($airingDateURL . "?access_token=$this->accessToken");
			$airingDates = json_decode($json, true);

			if(is_array($airingDates) && array_key_exists($nextEpisode, $airingDates))
				$timeStamp = $airingDates[$nextEpisode];
			else
				$timeStamp = -1;

			// Cache it
			apc_add($key, $timeStamp, $cacheTime);
		}

		return $timeStamp;
	}

	// Get anime list URL
	public function getAnimeListUrl($userName) {
		return "https://anilist.co/animelist/$userName";
	}

	// Get anime list
	public function getAnimeList($userName, $completedOnly = false) {
		$requestedStatus = $completedOnly ? 'completed' : 'watching';

		// Cached anime list
		$cacheTime = 20 * 60;
		$key = $userName . ":anilist-json";
		$json = apc_fetch($key, $found);

		// API URL
		if(!$found) {
			$apiUrl = "https://anilist.co/api/user/{userName}/animelist";
			$apiUrl = str_replace("{userName}", $userName, $apiUrl);

			// Execute
			$json = getJSON($apiUrl . "?access_token=$this->accessToken");

			// Cache it
			apc_add($key, $json, $cacheTime);
		}
		
		// Parse JSON
		$data = json_decode($json, true);

		// Watching list
		$watching = array();
		$lists = $data["lists"];

		if(!$lists || !array_key_exists($requestedStatus, $lists))
			return $watching;

		$aniListWatching = $lists[$requestedStatus];
		
		foreach($aniListWatching as $entry) {
			$anime = $entry["anime"];
			$title = $anime['title_romaji'];

			$episodesOffset = 0;

			if(array_key_exists($title, $this->specialOffsets))
				$episodesOffset = $this->specialOffsets[$title];

			$episodesWatched = $entry["episodes_watched"];
			$nextEpisodeToWatch = $episodesWatched + 1 + $episodesOffset;

			$newEntry = array(
				'title' => $title,
				'image' => str_replace('http://', 'https://', $anime['image_url_lge']),
				'url' => 'https://anilist.co/anime/' . $anime['id'],
				'id' => $anime['id'],
				'airingDate' => [
					'timeStamp' => '',
					'remaining' => '',
					'remainingString' => '',
				],
				'episodes' => array(
					'watched' => $episodesWatched ? $episodesWatched : 0,
					'next' => $nextEpisodeToWatch,
					'available' => 0,
					'max' => $anime['total_episodes'] ? $anime['total_episodes'] : -1,
					'offset' => $episodesOffset,
				)
			);

			$watching[] = $newEntry;
		}

		return $watching;
	}

	// Clear cache
	public function clearCache($userName) {
		$key = $userName . ":anilist-json";
		return apc_delete($key);
	}
}
?>