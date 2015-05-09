<?php
	$count = 0;

	$db->scan('arn', 'Users', function($record) {
		global $count;
		global $db;
		
		$user = $record['bins'];
		$userName = $user['userName'];
		$providerName = $user['providers']['list'];

		if($user['animeLists'][$providerName]['userName'] == '') {
			//echo $user['email'] . "<br>";
			echo "<a href='/+$userName' class='ajax'>$userName</a><br>";
			$count++;
		}

		/*$animeListKey = $db->initKey("arn", "AnimeLists", $userName);
		$db->remove($animeListKey);*/

		/*if($user['providers']['anime'] === 'KissAnime' || $user['providers']['anime'] === 'AnimeTwist') {
			$key = $db->initKey('arn', 'Users', $userName);
			$user['providers']['anime'] = 'AnimeShow';
			$db->put($key, $user);
		}*/

		/*if(!array_key_exists('balance', $user)) {
			$key = $db->initKey('arn', 'Users', $userName);
			$user['balance'] = 0.00;
			$db->put($key, $user);
		}*/
	});

	echo "$count users did not specify their list provider username.";
?>