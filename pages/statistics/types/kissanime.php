<?php
	$count = 0;

	$db->scan('arn', 'Users', function($record) {
		global $count;
		global $db;
		
		$user = $record['bins'];
		$userName = $user['userName'];
		$providerName = $user['providers']['list'];

		if($user['providers']['anime'] === 'KissAnime') {
			echo "<a href='/+$userName' class='ajax'>$userName</a><br>";
			$count++;
		}
	});

	echo "<br>$count users use KissAnime.";
?>