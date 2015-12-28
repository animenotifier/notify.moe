'use strict'

exports.render = function(request, render) {
	let user = request.user
	let nav = []

	nav.push({
		title: 'About',
		url: ''
	})

	if(user) {
		nav.push({
			title: 'Profile',
			url: 'profile/' + user.nick
		})
	}

	nav.push({
		title: 'Statistics',
		url: 'statistics'
	})

	if(user) {
		nav.push({
			title: 'Logout',
			url: 'logout',
			ajax: false
		})
	}

	render({
		user,
		nav
	})
}