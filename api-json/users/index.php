<?php
	// Setup
	set_include_path($_SERVER['DOCUMENT_ROOT']);
	require_once('config.php');

	$userName = $_GET["userName"];

	$db = new Aerospike($config["aeroSpike"]);
	$key = $db->initKey("arn", "Users", $userName);

	$db->get($key, $record);
	$user = $record['bins'];

	// Don't show mail
	unset($user['email']);

	foreach($user['animeLists'] as $providerName => $provider) {
		unset($user['animeLists'][$providerName]['auth']);
	}
	

	// Output
	header('Content-Type: application/json');
	echo json_encode($user);
?>