<?php
	// Setup
	set_include_path($_SERVER['DOCUMENT_ROOT']);
	require_once('fw/login.php');

	if(!$loggedIn) {
		echo "Not logged in.";
		return;
	}

	$userName = $_SESSION["accountId"];

	$db = new Aerospike($config["aeroSpike"]);
	$key = $db->initKey("arn", "Users", $userName);

	$listProvider = filter_var($_POST['listProvider'], FILTER_SANITIZE_STRING);
	$animeProvider = filter_var($_POST['animeProvider'], FILTER_SANITIZE_STRING);
	$timeProvider = filter_var($_POST['timeProvider'], FILTER_SANITIZE_STRING);

	$oldListProvider = filter_var($_POST['oldListProvider'], FILTER_SANITIZE_STRING);
	$oldListUserName = filter_var($_POST['oldListUserName'], FILTER_SANITIZE_STRING);

	$sortBy = filter_var($_POST['sortBy'], FILTER_SANITIZE_STRING);
	$tagLine = filter_var($_POST['tagLine'], FILTER_SANITIZE_STRING);
	$website = filter_var($_POST['website'], FILTER_SANITIZE_URL);
	$auth = filter_var($_POST['auth'], FILTER_SANITIZE_STRING);

	$db->get($key, $record);
	$user = $record['bins'];

	$user['animeLists'][$oldListProvider]['userName'] = $oldListUserName;
	$user['animeLists'][$oldListProvider]['auth'] = $auth;
	
	$user['providers'] = [
		'list' => $listProvider,
		'anime' => $animeProvider,
		'time' => $timeProvider,
	];

	$bins = [
		'providers' => $user['providers'],
		'animeLists' => $user['animeLists'],
		'sortBy' => $sortBy,
		'tagLine' => $tagLine,
		'website' => $website
	];

	$db->put($key, $bins);
	$db->close();
?>