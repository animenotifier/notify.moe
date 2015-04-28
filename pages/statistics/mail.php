<style scoped>
	h2 {
		text-align: center;
	}
</style>

<h2>Live statistics</h2>
<div id="listProviders" style="width: 100%; height: 500px;"></div>

<?php
	$providers = array();

	$db->scan('arn', 'Users', function($record) {
		global $providers;

		$user = $record['bins'];
		$providerName = $user['email'];
		$providerName = split('@', $providerName)[1];
		$providerName = explode('.', $providerName)[0];

		if($providerName == "googlemail")
			$providerName = "gmail";

		if(array_key_exists($providerName, $providers)) {
			$providers[$providerName] += 1;
		} else {
			$providers[$providerName] = 1;
		}
	});
?>

<script>
	//google.setOnLoadCallback(drawChart);
	document.addEventListener('DOMContentLoaded', setup);

	function setup() {
		document.removeEventListener('DOMContentLoaded', setup);

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
			['E-Mail provider', 'Users'],
			<?php
				foreach($providers as $providerName => $userCount) {
					echo "['$providerName', $userCount],";
				}
			?>
		]);

		var listProviderOptions = {
			title: 'E-Mail providers',
			fontName: 'Open Sans',
			backgroundColor: { fill:'transparent' },
			legend: {position: 'right', textStyle: {color: 'white', fontSize: 16}},
			titleTextStyle: {color: 'white', fontSize: 20},
			pieSliceBorderColor: 'transparent'
		};

		var listProviderChart = new google.visualization.PieChart(document.getElementById('listProviders'));
		listProviderChart.draw(listProviderData, listProviderOptions);
	}
</script>