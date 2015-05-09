<style scoped>
	.number {
		font-weight: bold;
	}
</style>

<?php
	date_default_timezone_set("UTC");

	$now = new DateTime();
	$activatedAccounts = 0;
	$totalAccounts24h = 0;
	$totalAccounts = 0;
	$loggedInAccounts = 0;

	$db->scan('arn', 'Accounts', function($record) {
		global $now;
		global $activatedAccounts;
		global $totalAccounts;
		global $totalAccounts24h;
		global $loggedInAccounts;

		$user = $record['bins'];

		$totalAccounts++;

		$created = (new DateTime($user['created']))->diff($now);
		if($created->d > 0 || $created->m > 0 || $created->y > 0)
			return;

		$totalAccounts24h++;

		if($user['activated'] === 1)
			$activatedAccounts++;

		$lastLogin = (new DateTime($user['lastLogin']))->diff($now);
		if($lastLogin->d > 0 || $lastLogin->m > 0 || $lastLogin->y > 0)
			return;

		$loggedInAccounts++;
	});

	echo "<div><span class='number'>$totalAccounts24h</span> new accounts have been registered.</div>";
	echo "<div><span class='number'>$activatedAccounts</span> new accounts have been activated.</div>";
	echo "<div><span class='number'>$loggedInAccounts</span> users have logged in.</div>";
	echo "<p></p>";
	echo "<div style='opacity: 0.4; font-size: 0.9em;'>In total <span class='number'>$totalAccounts</span> accounts have been registered over the whole lifetime of ARN.</div>";
?>