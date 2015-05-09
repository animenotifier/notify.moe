<?php
	require_once("../../config.php");
	require_once("../../fw/helper.php");
	require_once("helpers.php");
	require_once("sorting-algorithms.php");
	require_once("list-providers/interface.php");
	require_once("anime-providers/interface.php");
	require_once("time-providers/interface.php");

	date_default_timezone_set("UTC");

	//apc_clear_cache('user');
	if(!array_key_exists('userName', $_GET) || $_GET['userName'] == "") {
		echo "Please specify your username in the settings.";
		exit;
	}

	$userName = $_GET['userName'];

	// Content type
	header('Content-Type: application/json');

	// CORS
	header('Access-Control-Allow-Origin: *');

	$json = getAnimeListJSON($userName);
	
	// Output
	echo $json;
	
	// Get anime list JSON
	function getAnimeListJSON($userName) {
		global $config;
		global $sortingAlgorithms;
		
		$db = new Aerospike($config["aeroSpike"]);
		
		if(!$db->isConnected()) {
			echo "Failed to connect to the database [{$db->errorno()}]: {$db->error()}\n";
			exit(1);
		}
		
		$key = $db->initKey("arn", "Users", $userName);
		
		$status = $db->get($key, $record);
	
		if($status == Aerospike::ERR_RECORD_NOT_FOUND) {
			echo "The user ". $key['key']. " does not exist in the database.\n";
			exit(2);
		}
	
		if($status != Aerospike::OK) {
			echo $db->error();
			exit(3);
		}
	
		$user = $record["bins"];
		$providers = $user["providers"];
		$animeProviderName = @$_GET['animeProvider'] ?: $providers['anime'];

		if($animeProviderName === 'KissAnime')
			$animeProviderName = 'Nyaa';

		$listProviderName = $providers['list'];
		$timeProviderName = $providers['time'];
		
		$listProviderUserName = $user["animeLists"][$listProviderName]["userName"];
	
		// Include files
		require_once("list-providers/$listProviderName.php");
		require_once("anime-providers/$animeProviderName.php");
		require_once("time-providers/$timeProviderName.php");
	
		// Initialize list provider
		$listProvider = new $listProviderName($user);
		
		if(array_key_exists('clearListCache', $_GET) && $_GET['clearListCache'] == 1)
			$listProvider->clearCache($listProviderUserName);
		
		$onlyCompleted = array_key_exists('completed', $_GET) && $_GET['completed'] == 1;
		
		// Get watching list from list provider
		$watching = $listProvider->getAnimeList($listProviderUserName, $onlyCompleted);
	
		if(!$onlyCompleted && count($watching) > 0) {
			// Initialize anime provider
			$animeProvider = new $animeProviderName();
	
			// Initialize time provider
			if($timeProviderName === $listProviderName)
				$timeProvider = $listProvider;
			else
				$timeProvider = new $timeProviderName($user);
			
			$i = 0;
			foreach($watching as $entry) {
				$entry['listProvider'] = $listProviderName;
				$entry = $animeProvider->getAnimeInfo($entry);
				$entry = $timeProvider->getAiringDate($entry);
				$watching[$i] = $entry;
				$i++;
			}
	
			// MAL block message
			if($watching[0]['url'] === 'https://animereleasenotifier.com/incapsula') {
				$watching[0]['animeProvider']['url'] = $watching[0]['url'];
			}
	
			usort($watching, $sortingAlgorithms[$user['sortBy']]);
		} else {
			usort($watching, $sortingAlgorithms['title']);
		}

		$requestCount = 0;
		$animeListKey = $db->initKey('arn', 'AnimeLists', $userName);

		// Old anime list available?
		$status = $db->get($animeListKey, $animeListRecord);

		if($status === Aerospike::OK) {
			$oldAnimeList = $animeListRecord['bins'];
			$requestCount = intval($oldAnimeList['requestCount']) + 1;
		}
	
		// User data
		$user = array(
			'name' => $listProviderUserName,
			'listUrl' => $listProvider->getAnimeListUrl($listProviderUserName),
			'requestCount' => $requestCount,
			'timeStamp' => time(),
			'watching' => $watching
		);

		// Save in database
		if(!$onlyCompleted) {
			$db->put($animeListKey, $user);
		}
		
		// Close DB connection
		$db->close();

		return json_encode($user);
	}
?>