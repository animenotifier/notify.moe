<?php
	// Setup
	set_include_path($_SERVER['DOCUMENT_ROOT']);
	require_once('config.php');

	$userName = 'Aky';

	$db = new Aerospike($config["aeroSpike"]);

	// 1
	$before = microtime(true);

	for ($i = 0; $i < 10000; $i++) {
		$key = $db->initKey('arn', 'Users', $userName);
		$db->get($key, $record);
		$user = $record['bins'];
	}

	$after = microtime(true);
	echo ($after - $before) / $i . " sec/GET<br>";

	// 2
	$before = microtime(true);

	for ($i = 0; $i < 10000; $i++) {
		$key = $db->initKey('arn', 'Users', $userName);
		$db->get($key, $record);
		$user = $record['bins'];
	}

	$after = microtime(true);
	echo ($after - $before) / $i . " sec/GET<br>";
?>