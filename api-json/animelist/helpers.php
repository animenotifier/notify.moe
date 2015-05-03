<?php
	// getHTML
	function getHTML($url, $cookie = '', $agent = 'Anime Release Notifier') {
		; ///1.0 (+https://animereleasenotifier.com/)
		//$agent = 'Atarashii/1.0';
		//$agent = 'Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.115 Safari/537.36';

		$header = array(
			"User-Agent: $agent",
			'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8',
			'Accept-Language: en-US,en;q=0.8,de-DE;q=0.6,de;q=0.4,ja;q=0.2',
			'Accept-Charset: utf-8;q=0.7,*;q=0.7',
			//'Accept-Encoding: gzip, deflate, sdch',
			'Cache-Control: max-age=0',
			"Cookie: $cookie",
			'Keep-Alive: 115',
			'Connection: keep-alive',
			"Referer: $url"
		);

		$curl = curl_init();

		curl_setopt($curl, CURLOPT_URL, $url);
		curl_setopt($curl, CURLOPT_RETURNTRANSFER, 1);
		curl_setopt($curl, CURLOPT_CONNECTTIMEOUT, 5);
		curl_setopt($curl, CURLOPT_USERAGENT, $agent);
		curl_setopt($curl, CURLOPT_HTTPHEADER, $header);
		//curl_setopt($curl, CURLOPT_SSL_VERIFYPEER, false);

		$data = curl_exec($curl);
		$status = curl_getinfo($curl, CURLINFO_HTTP_CODE);

		if($status !== 200)
			$data = '';

		curl_close($curl);

		return $data;
	}

	// getJSON
	function getJSON($url, $additionalHeader = null) {
		$curl = curl_init($url);
		curl_setopt($curl, CURLOPT_CUSTOMREQUEST, "GET");
		curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);

		if($additionalHeader !== null)
			curl_setopt($curl, CURLOPT_HTTPHEADER, array('Accept: application/json', $additionalHeader));
		else
			curl_setopt($curl, CURLOPT_HTTPHEADER, array('Accept: application/json'));

		return curl_exec($curl);
	}

	// getLinkFromGoogle
	function getLinkFromGoogle($title, $site, $linkNeedsToContain) {
		global $config;
		
		$title = preg_replace('/[^[:alnum:]\'*]/ui', '+', $title);

		$googleURL = "http://ajax.googleapis.com/ajax/services/search/web?v=1.0&key=" . $config['googleAPIKey'] . "&q=site:$site+" . $title;
		$googleResults = getHTML($googleURL);

		// Parse JSON
		$googleData = json_decode($googleResults, true);

		$results = $googleData['responseData']['results'];

		if(count($results) === 0)
			return '';

		foreach($results as $result) {
			if(strpos($result['url'], $linkNeedsToContain) !== false)
				return $result['url'];
		}

		return '';
	}

	// plural
	function plural($count, $singular) {
		return ($count === 1 || $count === -1) ? "$count $singular" : ("$count $singular" . 's');
	}
?>