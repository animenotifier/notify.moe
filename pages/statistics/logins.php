<p>
<?php
	$list = array();

	$db->scan('arn', 'Accounts', function($record) {
		global $list;

		$account = $record['bins'];
		$list[] = $account;
	});

	usort($list, function($a, $b) {
		return strnatcmp($b['lastLogin'], $a['lastLogin']);
	});

	foreach($list as $account) {
		$userName = $account['userName'];
		$loginDate = $account['lastLogin'];
		$loginDate = str_replace('T', ' ', $loginDate);
		$loginDate = str_replace('+0000', '', $loginDate);

		echo "$loginDate: <a href='/+$userName' class='ajax'>$userName</a><br>";
	}
?>
</p>