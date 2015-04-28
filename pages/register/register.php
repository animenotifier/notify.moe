<?php
	set_include_path($_SERVER['DOCUMENT_ROOT']);
	require_once("fw/login.php");
?>

<style scoped>
	#register-page {
		width: 100%;
		float: left;
	}

	#register-page form {
		width: 100%;
		max-width: 25em;
		display: inline-block;
	}
</style>

<?php
	if(isset($loggedIn) && $loggedIn) {
		echo "Already logged in as <a href='/+" . $_SESSION["accountId"] . "' class='ajax'>" . $_SESSION["accountId"] . "</a>.";
		return;
	}

	if(AeroSpikeLogin::$lastError != null) {
		echo '<div class="error">' . AeroSpikeLogin::$lastError . "</div>";
	}
?>

<section id="register-page">
	<div style="width: 100%; text-align: center;">
		<form name="register" id="register" action="/register" method="post" accept-charset="utf-8" autocomplete="on">
			<div>
				<input type="email" placeholder="E-Mail" name="email" id="email" value="<?php echo @$_POST["email"] ?: ""; ?>" autofocus required>
			</div>
			<div>
				<input type="password" placeholder="Password" name="password" id="password" value="<?php echo @$_POST["password"] ?: ""; ?>" autocomplete="off" required>
			</div>
			<div>
				<input type="text" placeholder="User name" name="username" id="username" maxlength="25" value="<?php echo @$_POST["username"] ?: ""; ?>" required>
			</div>
			<input type="hidden" name="register" value="1">
			
			<div class="center">
				<input type="submit" id="submit-button" value="Register" class="hover-box button">
			</div>
		</form>
	</div>
</section>