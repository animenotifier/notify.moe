<?php
	// Setup
	set_include_path($_SERVER['DOCUMENT_ROOT']);
	require_once('fw/login.php');

	if(!array_key_exists('email', $_SESSION)) {
		echo "Not logged in.";
		return;
	}

	$email = $_SESSION['email'];

	if($loginSystem->sendActivationMail($email)) {
		echo "Sent activation mail to <strong>$email</strong>.";
	} else {
		echo "Error sending activation mail: " . AeroSpikeLogin::$lastError;
	}
?>