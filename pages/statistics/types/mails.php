<div id="listProviders" style="width: 100%; height: 500px;"></div>

<?php
	$providers = array();

	$db->scan('arn', 'Accounts', function($record) {
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
				arsort($providers);
				foreach($providers as $providerName => $userCount) {
					if($userCount >= 10)
						echo "['$providerName', $userCount],";
				}
			?>
		]);

		var listProviderOptions = {
			title: 'E-Mail providers',
			fontName: 'Open Sans',
			backgroundColor: { fill:'transparent' },
			legend: {position: 'right', textStyle: {color: 'black', fontSize: 16}},
			titleTextStyle: {color: 'black', fontSize: 20},
			pieSliceBorderColor: 'transparent'
		};

		var listProviderChart = new google.visualization.PieChart(document.getElementById('listProviders'));
		listProviderChart.draw(listProviderData, listProviderOptions);
	}
</script>