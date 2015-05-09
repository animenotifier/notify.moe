<div id="listProviders" style="width: 100%; height: 500px;"></div>
<div id="animeProviders" style="width: 100%; height: 500px;"></div>
<div id="combinations" style="width: 100%; height: 500px;"></div>

<?php
	$providers = array();
	$animeProviders = array();
	$combinations = array();

	$db->scan('arn', 'Users', function($record) {
		global $db;
		global $providers;
		global $animeProviders;
		global $combinations;

		$user = $record['bins'];
		$userName = $user['userName'];
		$providerName = $user['providers']['list'];
		$animeProviderName = $user['providers']['anime'];

		$listUserName = $user['animeLists'][$providerName]['userName'];

		// When a user name is not specified, skip this account
		if(strlen($listUserName) == 0)
			return;

		/*// April's fool
		if($animeProviderName == "Nyaa")
			$animeProviderName = "Cool";
		else
			$animeProviderName = "Uncool";

		if($providerName == "AniList" || $providerName == "HummingBird")
			$providerName = "Cool";
		else
			$providerName = "Uncool";
		// End April's fool*/

		$animeListKey = $db->initKey('arn', 'AnimeLists', $userName);
		$status = $db->get($animeListKey, $animeListRecord);

		if($status != Aerospike::OK)
			return;

		$animeList = $animeListRecord['bins'];

		if(!array_key_exists('timeStamp', $animeList))
			return;

		$lastUpdate = $animeList['timeStamp'];

		if(time() - $lastUpdate > 24 * 60 * 60)
			return;

		if(array_key_exists($providerName, $providers)) {
			$providers[$providerName] += 1;
		} else {
			$providers[$providerName] = 1;
		}

		if(array_key_exists($animeProviderName, $animeProviders)) {
			$animeProviders[$animeProviderName] += 1;
		} else {
			$animeProviders[$animeProviderName] = 1;
		}

		$combinationName = "$providerName + $animeProviderName";

		if(array_key_exists($combinationName, $combinations)) {
			$combinations[$combinationName] += 1;
		} else {
			$combinations[$combinationName] = 1;
		}
	});
?>

<script>
	//google.setOnLoadCallback(drawChart);
	document.addEventListener('DOMContentLoaded', setup);

	function setup() {
		setPageHandler(function(pageId) {
			document.removeEventListener('DOMContentLoaded', setup);
		});

		$.ajax({
			url: 'https://www.google.com/jsapi?callback',
			cache: true,
			dataType: 'script',
			success: function() {
				google.load('visualization', '1', {packages:['corechart'], 'callback' : function() {
					drawChart();
				}});

				return true;
			}
		});
	}

	function drawChart() {
		var listProviderData = google.visualization.arrayToDataTable([
			['List provider', 'Users'],
			<?php
				arsort($providers);
				foreach($providers as $providerName => $userCount) {
					echo "['$providerName', $userCount],";
				}
			?>
		]);

		var animeProviderData = google.visualization.arrayToDataTable([
			['Anime provider', 'Users'],
			<?php
				arsort($animeProviders);
				foreach($animeProviders as $providerName => $userCount) {
					echo "['$providerName', $userCount],";
				}
			?>
		]);

		var combinationData = google.visualization.arrayToDataTable([
			['Combination', 'Users'],
			<?php
				arsort($combinations);
				foreach($combinations as $providerName => $userCount) {
					echo "['$providerName', $userCount],";
				}
			?>
		]);

		var listProviderOptions = {
			title: 'List providers',
			fontName: 'Open Sans',
			backgroundColor: { fill:'transparent' },
			legend: {position: 'right', textStyle: {color: 'black', fontSize: 16}},
			titleTextStyle: {color: 'black', fontSize: 20},
			pieSliceBorderColor: 'transparent',
			pieSliceText: 'percentage'
		};

		var listProviderChart = new google.visualization.PieChart(document.getElementById('listProviders'));
		listProviderChart.draw(listProviderData, listProviderOptions);

		listProviderOptions.title = 'Anime providers';

		var animeProviderChart = new google.visualization.PieChart(document.getElementById('animeProviders'));
		animeProviderChart.draw(animeProviderData, listProviderOptions);

		listProviderOptions.title = 'Popular combinations';

		var combinationChart = new google.visualization.PieChart(document.getElementById('combinations'));
		combinationChart.draw(combinationData, listProviderOptions);
	}
</script>