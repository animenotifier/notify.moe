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

	// Header
	include('pages/statistics/header.php');

	// Include the selected type
	if(!isset($params) || count($params) == 0 || $params[0] == '') {
		require_once('pages/statistics/types/providers.php');
	} else {
		$type = $params[0];
		require_once("pages/statistics/types/$type.php");
	}
?>