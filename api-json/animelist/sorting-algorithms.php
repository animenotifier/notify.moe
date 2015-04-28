<?php
	// Sort
	$sortingAlgorithms = [
		'airingDate' => function($a, $b) {
			$aTS = intval($a['airingDate']['timeStamp']);
			$bTS = intval($b['airingDate']['timeStamp']);

			if($aTS == -1)
				return 999999999;

			if($bTS == -1)
				return -999999999;

			return ($aTS - $bTS);
		},

		'title' => function($a, $b) {
			return strnatcmp(strtolower($a['title']), strtolower($b['title']));
		}
	];
?>