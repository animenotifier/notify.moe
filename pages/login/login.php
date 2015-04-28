<style scoped>
	<?php
		include("login.css");
	?>
</style>

<?php
	// Setup
	set_include_path($_SERVER['DOCUMENT_ROOT']);
	require_once("fw/login.php");

	if(!isset($loggedIn) || !$loggedIn) {
		include("login-form.php");
	}

	if(isset($wasLoggedIn) && $wasLoggedIn) {
		echo "Already logged in as <a href='/+" . $_SESSION["accountId"] . "' class='ajax'>" . $_SESSION["accountId"] . "</a>.";
	} else if(isset($logInSuccessful)) {
		if($logInSuccessful === true) {
			echo "Successfully logged in as <a href='/+" . $_SESSION["accountId"] . "' class='ajax'>" . $_SESSION["accountId"] . "</a>.";
		} else if($logInSuccessful === false) {
			echo "Wrong mail or password.";
		}
	}

	/*if(isset($loggedIn) && $loggedIn) {
		echo '<br/><br/><a href="/logout">Logout</a>';
	}*/
?>