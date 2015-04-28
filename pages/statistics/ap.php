<?php
	$count = 0;

	$db->scan('arn', 'Users', function($record) {
		global $count;
		
		$user = $record['bins'];
		$userName = $user['userName'];
		$providerName = $user['providers']['list'];

		if($user['animeLists'][$providerName]['userName'] == '') {
			//echo $user['email'] . "<br>";
			echo "<a href='/+$userName' class='ajax'>$userName</a><br>";
			$count++;
		}
	});

	echo "$count users did not specify their list provider username.";
?>