'use strict'

exports.render = function(request, render) {
	let user = request.user
	let nav = []

	nav.push({
		title: user ? 'Dashboard' : 'About',
		url: '',
		icon: 'dashboard'
	})

	if(user) {
		nav.push({
			title: 'Profile',
			url: 'profile/' + user.nick,
			icon: 'user'
		})
	}

	nav.push({
		title: 'Stats',
		url: 'statistics',
		icon: 'stats'
	})

	if(user) {
		nav.push({
			title: 'Logout',
			url: 'logout',
			icon: 'log-out',
			ajax: false
		})
	}

	render({
		user,
		nav
	})
}