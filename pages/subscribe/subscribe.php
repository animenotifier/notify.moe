<?php
	// Setup
	set_include_path($_SERVER['DOCUMENT_ROOT']);
	require_once('fw/login.php');

	if(!$loggedIn) {
		echo "Not logged in!";
		return;
	}

	$userName = $_SESSION['accountId'];
?>

<a href="https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=MSTSA6EXDH2GY&item_name=ARN Basic 1 Year (for <?php echo $userName; ?>)&custom=<?php echo $userName; ?>" target="_blank">ARN Basic 1 Year</a>