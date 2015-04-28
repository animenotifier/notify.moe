<?php
	$db->scan('arn', 'Users', function($record) {
		$user = $record['bins'];
		$userName = $user['userName'];
		$tagLine = @$user['tagLine'];

		if($tagLine) {
			echo "<a href='/+$userName' class='ajax'>$tagLine</a><br>";
		}
	});
?>