'use strict'

exports.render = function(request, render) {
	let user = request.user
	let nav = []
	let embedded = request.params.indexOf('embedded') !== -1

	nav.push({
		title: 'Dash',
		url: '',
		icon: 'dashboard'
	})

	if(user) {
		nav.push({
			title: 'Profile',
			url: embedded ? '+/watching' : '+' + user.nick,
			icon: 'user'
		})
	}

	nav.push({
		title: 'Anime',
		url: 'anime',
		icon: 'television'
	})
	
	if(!arn.production) {
		nav.push({
			title: 'Forum',
			url: 'forum',
			icon: 'comment'
		})
	}

	if(user) {
		nav.push({
			title: 'Settings',
			url: 'settings',
			icon: 'cog'
		})
	} else {
		nav.push({
			title: 'FAQ',
			url: 'faq',
			icon: 'question-circle'
		})
	}

	if(user) {
		if(!embedded) {
			nav.push({
				title: '',
				url: 'logout',
				icon: 'sign-out',
				ajax: false,
				float: 'right',
				tooltip: 'Logout'
			})
		}

		/*nav.push({
			title: '',
			url: 'https://www.paypal.me/blitzprog',
			icon: 'heart',
			float: 'right',
			tooltip: 'Donate'
		})*/

		nav.push({
			title: '',
			url: 'faq',
			icon: 'question-circle',
			float: 'right',
			tooltip: 'FAQ'
		})

		/*
		if(!embedded) {
			nav.push({
				title: '',
				url: 'api',
				icon: 'code',
				float: 'right',
				tooltip: 'API (for developers)'
			})
		}

		nav.push({
			title: '',
			url: 'statistics',
			icon: 'bar-chart',
			float: 'right',
			tooltip: 'Statistics'
		})

		nav.push({
			title: '',
			url: 'changes',
			icon: 'refresh',
			float: 'right',
			tooltip: 'Changes'
		})

		nav.push({
			title: '',
			url: 'roadmap',
			icon: 'road',
			float: 'right',
			tooltip: 'Roadmap'
		})*/
		
		nav.push({
			title: '',
			url: 'others',
			icon: 'table',
			float: 'right',
			tooltip: 'Others'
		})
		
		nav.push({
			title: '',
			url: 'users',
			icon: 'globe',
			float: 'right',
			tooltip: 'Users'
		})

		if(!embedded && (user.role === 'admin' || user.role === 'editor')) {
			nav.push({
				title: '',
				url: 'admin',
				icon: 'wrench',
				float: 'right',
				tooltip: 'Admin Panel'
			})
		}
	}

	render({
		user,
		nav,
		maintenance: arn.maintenance,
		embedded,
		request
	})
}