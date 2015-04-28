<?php
	// Setup
	set_include_path($_SERVER['DOCUMENT_ROOT']);
	require_once('fw/page.php');
	require_once('config.php');

	$db = new Aerospike($config["aeroSpike"]);

	if(!$db->isConnected()) {
		echo "Failed to connect to the database [{$db->errorno()}]: {$db->error()}\n";
		exit(1);
	}

	//require_once('pages/statistics/providers.php');

	if(!isset($params) || count($params) == 0 || $params[0] == '') {
		require_once('pages/statistics/providers.php');
		return;
	}

	$type = $params[0];

	if($type == "mail") {
		require_once('pages/statistics/mail.php');
		return;
	}

	if($type == "ap") {
		require_once('pages/statistics/ap.php');
		return;
	}

	if($type == "logins") {
		require_once('pages/statistics/logins.php');
		return;
	}

	if($type == "taglines") {
		require_once('pages/statistics/taglines.php');
		return;
	}
?>