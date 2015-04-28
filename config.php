<?php
	$config = array(
		'title' => 'Anime Release Notifier',
		'pageURL' => '/pages/{name}/{name}.php',
		'favIcon' => '/animereleasenotifier.png',
		'keywords' => 'anime,release,notifier',
		'author' => 'Eduard Urbach',
		'dateFormat' => 'F j, Y',
		'font' => 'Open Sans',

		'loginSystem' => 'AeroSpikeLogin',
		'googleAPIKey' => 'AIzaSyC02wQE4rM945X-Yp1vQkT8RJm3a5Qmplk',

		'cssFiles' => array(
			'css/animereleasenotifier.css',
			'css/loading-animation.css',
			'css/layout.css',
			'css/elements.css',
			'css/headers.css',
			'css/forms.css',
			'css/colors.css',
			'css/anime.css'
		),

		'aeroSpike' => [
			'hosts' => [
				[
					'addr' => '127.0.0.1',
					'port' => 3000
				]
			]
		],

		'frontPage' => 'home',

		'pages' => array(
			'home' => array(
				'title' => 'About',
				'url' => '',
				'description' => 'Fetches your anime watching list and notifies you when a new anime episode is available. It also displays the time until a new episode is released.

Supports anime lists from:
 - anilist.co
 - anime-planet.com
 - myanimelist.net
 - hummingbird.me

アニメリリースチェッカー。'
			),

			'users' => array(
				'title' => 'Users',
				'description' => 'Users',
				'visible' => false
			),

			'incapsula' => array(
				'title' => 'MAL and Incapsula',
				'description' => 'MAL and Incapsula',
				'visible' => false
			),

			'profile' => array(
				'title' => 'Profile',
				'description' => 'Profile',
				'visible' => false
			),

			'donate' => array(
				'title' => 'Donate',
				'description' => 'Donations'
			),

			'statistics' => array(
				'title' => 'Statistics',
				'description' => 'Statistics',
			),

			'all' => array(
				'title' => 'Users',
				'description' => 'Users',
			),

			'roadmap' => array(
				'title' => 'Roadmap',
				'description' => 'Roadmap',
				'visible' => false
			),

			'feedback' => array(
				'title' => 'Feedback',
				'description' => 'Feedback',
				'visible' => false
			),

			'irc' => array(
				'title' => 'IRC',
				'description' => 'IRC',
				'visible' => false
			),

			'pc' => array(
				'title' => 'PC/Chrome',
				'description' => 'PC/Chrome',
				'visible' => false
			),

			'logout' => array(
				'title' => 'Logout',
				'description' => 'Logout'
			),
			
			'login' => array(
				'title' => 'Login',
				'description' => 'Login'
			),

			'register' => array(
				'title' => 'Register',
				'description' => 'Registration'
			),

			'resendmail' => array(
				'title' => 'Resend activation mail',
				'description' => 'Resend activation mail',
				'visible' => false
			),

			'legal' => array(
				'title' => 'Legal notice',
				'description' => 'Legal notice',
				'visible' => false
			),

			'recruit' => array(
				'title' => 'Recruitment',
				'description' => 'Recruitment',
				'visible' => false
			),

			'preview' => array(
				'title' => '',
				'description' => '',
				'visible' => false
			),
		)
	);
?>