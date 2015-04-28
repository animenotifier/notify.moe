<?php
	// Setup
	set_include_path($_SERVER['DOCUMENT_ROOT']);
	require_once('fw/page.php');
	require_once('fw/login.php');
?>

<style scoped>
	h3 {
		margin-top: 1em;
		margin-bottom: 1em;
		text-align: center;
	}

	.user-link {
		display: block;
		float: left;
		width: 240px;
		border: 1px solid transparent;
		border-radius: 3px;
		margin: 0.5em;
		background: transparent;
		font-size: 1.25em;
		overflow: hidden;
	}

	.user-link:hover {
		border-color: rgba(0, 0, 0, 0.25);
		/*box-shadow: 0 0 8px rgba(0, 0, 0, 0.5);*/
	}

	.user-list-provider {
		color: black;
		opacity: 0.5;
		font-size: 0.7em;
	}

	.user-link img {
		
	}
</style>

<?php
	$users = array();

	function showAllUsers($db) {
		global $users;

		// Fetch
		$db->scan('arn', 'Users', function($record) {
			global $users;

			$user = $record['bins'];
			$users[] = $user;
		});

		// Sort
		usort($users, function($a, $b) {
			$result = strnatcmp($a['providers']['list'], $b['providers']['list']);

			if($result == 0)
				return strnatcmp(strtolower($a['userName']), strtolower($b['userName']));

			return $result;
		});

		// Display
		$provider = $users[0]['providers']['list'];
		echo "<h3>$provider</h3>";

		foreach($users as $user) {
			$userName = $user['userName'];
			$email = $user['email'];
			$userProvider = $user['providers']['list'];

			if($provider !== $userProvider) {
				$provider = $userProvider;
				echo "<h3>$provider</h3>";
			}

			$a = $user['animeLists'][$userProvider]['userName'];

			if(!$a)
				$a = $userProvider;

			$gravatarHash = md5(strtolower($email));
			echo "<a href='/+$userName' class='user-link ajax'><img src='https://www.gravatar.com/avatar/$gravatarHash?s=64&d=blank' alt='$userName' width='64' height='64'><br/>$userName<br/><span class='user-list-provider'>" . $a . "</span></a>";
		}
	}

	showAllUsers($loginSystem->db);
?>