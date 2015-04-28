<?php
	$emailValue = "";

	if(array_key_exists("email", $_POST)) {
		$emailValue = $_POST["email"];
		$emailValue = filter_var($emailValue, FILTER_SANITIZE_EMAIL);
	}
?>

<section id="login-page">
	<div style="width: 100%; text-align: center;">
		<form name="login" id="login" action="/login" method="post" accept-charset="utf-8" onload="loadMail();">
			<div>
				<input type="email" placeholder="E-Mail" name="email" id="email" value="<?php echo $emailValue; ?>" autofocus required>
			</div>
			<div>
				<input type="password" placeholder="Password" name="password" id="password" required>
			</div>
			<div class="center">
				<input type="submit" id="submit-button" value="Log in" class="hover-box button">
				<!--<a href="#">Lost your password?</a>
				<a href="#">Register</a>-->
			</div>
		</form>
	</div>

	<br>

	<!--<div id="signInButton">
		<span
			class="g-signin"
			data-callback="googleSignIn"
			data-clientid="941298467524-al3j1d1im1cv2j1dvo0uq624dag3hb4k.apps.googleusercontent.com"
			data-cookiepolicy="single_host_origin"
			data-scope="profile https://www.googleapis.com/auth/userinfo.email">
		</span>
	</div>-->

	<script src="https://apis.google.com/js/client:platform.js" async defer></script>

	<!--<a href="/resetpassword" class="hover-box button ajax" style="float: right;">Forgot your password?</a>-->
</section>

<script>
	var rememberMail = function () {
		loadMail();
		
		$("#login").submit(saveMail);

		document.removeEventListener('DOMContentLoaded', rememberMail);
	};

	document.addEventListener('DOMContentLoaded', rememberMail);

	// Load mail
	function loadMail() {
		var email = localStorage["email"];

		if(typeof email != "undefined") {
			$("#email").val(email);
			$("#password").focus();
		}
	}

	// Save mail
	function saveMail() {
		email = $("#email").val();
		localStorage["email"] = email;
	};

	function googleSignIn(authResult) {
		if (authResult['status']['signed_in']) {
			// Update the app to reflect a signed in user
			// Hide the sign-in button now that the user is authorized, for example:
			$("#signInButton").hide();

			var accessToken = authResult['access_token'];
			console.log(authResult);
			$.getJSON("https://www.googleapis.com/plus/v1/people/me?fields=gender%2Cbirthday%2Curl%2Cnickname%2Clanguage%2CdisplayName%2Cemails%2Fvalue%2Cid&key=AIzaSyC02wQE4rM945X-Yp1vQkT8RJm3a5Qmplk&access_token=" + accessToken, function(user) {
				console.log(user);
			});
		} else {
			// Update the app to reflect a signed out user
			// Possible error values:
			//   "user_signed_out" - User is signed-out
			//   "access_denied" - User denied access to your app
			//   "immediate_failed" - Could not automatically log in the user
			console.log('Sign-in state: ' + authResult['error']);
		}
	}
</script>