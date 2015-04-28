<?php
	// Riak setup
	set_include_path($_SERVER['DOCUMENT_ROOT']);
	require_once('fw/login.php');

	echo $loginSystem->activateAccount($_GET['email'], $_GET['token']);
?>