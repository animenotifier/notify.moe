<?php
	// Setup
	set_include_path($_SERVER['DOCUMENT_ROOT']);
	require_once('fw/helper.php');

	sendFeedback($_POST['userName'], $_POST['email'], $_POST['message']);

	// Send mail
	function sendFeedback($userName, $email, $message) {
		$adminMail = 'e.urbach@gmail.com';

		$userName = filter_var($userName, FILTER_SANITIZE_STRING);
		$email = filter_var($email, FILTER_SANITIZE_EMAIL);

		$headers  = "MIME-Version: 1.0\r\n";
		$headers .= "Content-Type: text/plain; charset=UTF-8\r\n";
		$headers .= "From: $userName <$email>\r\n";

		sendSlackMessage("Feedback from: https://animereleasenotifier.com/+$userName\n\n$message", 'feedback');

		return mail(
			$adminMail,
			"Anime Release Notifier - Feedback from $userName",
			"From: https://animereleasenotifier.com/+$userName\n\n" . $message,
			$headers,
			'-f noreply@animereleasenotifier.com'
		);
	}
?>