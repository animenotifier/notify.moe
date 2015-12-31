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
			url: 'user/' + user.nick,
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
		title: 'Stats',
		url: 'statistics',
		icon: 'stats'
	})

	if(user) {
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