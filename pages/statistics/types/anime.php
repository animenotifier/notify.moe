<style scoped>
	.number {
		font-weight: bold;
	}

	.top-anime-list {

	}

	.top-anime {
		display: block;
		text-align: left;
		border: 1px solid rgba(0, 0, 0, 0.25);
		background: rgba(0, 0, 0, 0.05);
		border-radius: 3px;
		margin: 0.2em auto;
		padding: 0.5em 0.75em;

		max-width: 1000px;
	}

	.watch-count {
		float: right;
	}
</style>

<?php
	$watchingAnimeCount = 0;
	$usersCounted = 0;
	$last60Seconds = 0;
	$last24Hours = 0;
	$top = [];

	// Helper
	$minute = 60;
	$hour = 60 * $minute;
	$day = 24 * $hour;

	// Scan
	$db->scan('arn', 'AnimeLists', function($record) {
		global $usersCounted;
		global $watchingAnimeCount;
		global $last60Seconds;
		global $last24Hours;
		global $top;

		global $minute;
		global $day;

		$user = $record['bins'];
		$watching = $user['watching'];
		$watchingCount = count($watching);

		if(!array_key_exists('timeStamp', $user))
			return;

		if(time() - $minute <= $user['timeStamp'])
			$last60Seconds++;

		if(time() - $day <= $user['timeStamp'])
			$last24Hours++;

		if($watchingCount === 0)
			return;

		foreach($watching as $anime) {
			$title = $anime['title'];

			if(array_key_exists($title, $top)) {
				$top[$title] = $top[$title] + 1;
			} else {
				$top[$title] = 1;
			}
		}

		$watchingAnimeCount += $watchingCount;
		$usersCounted++;
	});

	if($usersCounted == 0)
		return;

	$averageAnimeCount = $watchingAnimeCount / floatval($usersCounted);
?>

<div>
	<?php $maxCount = 5; ?>
	Currently the <span class='number'>Top <?php echo $maxCount; ?></span> most watched anime titles are:
</div>

<p></p>

<div class="top-anime-list">
	<?php
		arsort($top);

		$iterationCount = 0;
		foreach($top as $title => $userCount) {
			echo "<div class='top-anime'>$title <div class='watch-count'>$userCount</div></div>";

			$iterationCount++;
			if($iterationCount === $maxCount)
				break;
		}
	?>
</div>

<p></p>
<p></p>

<div>
	The average person is watching <span class='number'><?php echo round($averageAnimeCount); ?></span> anime per week.
</div>

<p></p>

<div>
	<span class='number'><?php echo $last60Seconds; ?></span> users have updated their list in the last 60 seconds.
</div>

<div>
	<span class='number'><?php echo $last24Hours; ?></span> users have updated their list in the last 24 hours.
</div>