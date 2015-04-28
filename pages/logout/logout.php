<?php
	// Setup
	set_include_path($_SERVER['DOCUMENT_ROOT']);
	require_once("fw/login.php");
	
	if(isset($_SESSION["accountId"])) {
		session_destroy();
		unset($_SESSION["accountId"]);
		unset($_SESSION["email"]);

		if(!isset($_SESSION["accountId"])) {
			$loggedIn = false;
			echo "Logged out successfully.";
		}
	} else {
		echo "Not logged in.";
	}
?>

<script>
	// Reset page cache
	cache = [];
	location.replace("/");
</script>