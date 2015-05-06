<style scoped>
	.number {
		font-weight: bold;
	}
</style>

<?php
	$activatedAccounts = 0;
	$totalAccounts = 0;

	$db->scan('arn', 'Accounts', function($record) {
		global $activatedAccounts;
		global $totalAccounts;

		$user = $record['bins'];

		$totalAccounts++;

		if($user['activated'] === 1)
			$activatedAccounts++;
	});

	echo "<div><span class='number'>$totalAccounts</span> accounts have been registered.</div>";
	echo "<div><span class='number'>$activatedAccounts</span> accounts have been activated.</div>";
?>