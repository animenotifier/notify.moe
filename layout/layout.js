'use strict'

exports.render = function(request, render) {
	let user = request.user
	let nav = []

	nav.push({
		title: 'Dash',
		url: '',
		icon: 'dashboard'
	})

	if(user) {
		nav.push({
			title: 'Profile',
			url: '+' + user.nick,
			icon: 'user'
		})
	}

	// nav.push({
	// 	title: 'Anime',
	// 	url: 'anime',
	// 	icon: 'eye-open'
	// })

	nav.push({
		title: 'Users',
		url: 'users',
		icon: 'globe'
	})

	nav.push({
		title: 'FAQ',
		url: 'faq',
		icon: 'question-sign'
	})

	nav.push({
		title: 'Changes',
		url: 'changes',
		icon: 'refresh'
	})

	if(user) {
		// nav.push({
		// 	title: 'Youtube',
		// 	url: 'videos',
		// 	icon: 'facetime-video'
		// })
		//
		// nav.push({
		// 	title: 'News',
		// 	url: 'news',
		// 	icon: 'pushpin'
		// })

		// nav.push({
		// 	title: 'Stats',
		// 	url: 'statistics',
		// 	icon: 'stats'
		// })

		nav.push({
			title: '',
			url: 'logout',
			icon: 'log-out',
			ajax: false,
			float: 'right'
		})
	}

	render({
		user,
		nav
	})
}