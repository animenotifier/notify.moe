<?php
	// Setup
	set_include_path($_SERVER['DOCUMENT_ROOT']);
	require_once('fw/page.php');
	require_once('fw/login.php');

	if(!$loggedIn) {
		echo "Not logged in.";
		return;
	}
?>

<style scoped>
	textarea {
		border-radius: 2px;
		padding: 0.75em;
	}
</style>

<form id="feedback">
	<textarea id="message" placeholder="Did something happen?" rows="5" cols="50" required></textarea>
	<br>
	<input type="submit" id="submit-button" value="Send message" class="hover-box button">
</form>

<script>
	document.addEventListener('DOMContentLoaded', init);

	function init() {
		setPageHandler(function(pageId) {
			document.removeEventListener('DOMContentLoaded', init);
		});

		$("#submit-button").click(function(e) {
			e.preventDefault();

			var postData = {
				userName: "<?php echo $_SESSION['accountId']; ?>",
				email: "<?php echo $_SESSION['email']; ?>",
				message: $("#message").val()
			};

			$loadingAnimation.fadeIn(fadeSpeed);
			
			$.post("/pages/feedback/mail.php", postData, function(response) {
				$loadingAnimation.fadeOut(fadeSpeed);
				$("#feedback").html("Thank you, your message has been sent!");
			});
		});
	}
</script>