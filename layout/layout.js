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

	nav.push({
		title: 'Users',
		url: 'users',
		icon: 'globe'
	})

	nav.push({
		title: 'Anime',
		url: 'anime',
		icon: 'film'
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

	nav.push({
		title: '',
		url: 'faq',
		icon: 'question-sign',
		float: 'right'
	})

	nav.push({
		title: '',
		url: 'changes',
		icon: 'refresh',
		float: 'right'
	})

	nav.push({
		title: '',
		url: 'roadmap',
		icon: 'road',
		float: 'right'
	})

	render({
		user,
		nav
	})
}