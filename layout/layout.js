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

	nav.push({
		title: 'Forum',
		url: 'forum',
		icon: 'comment'
	})

	if(!user) {
		nav.push({
			title: 'Users',
			url: 'users',
			icon: 'globe'
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

		nav.push({
			title: '',
			url: 'faq',
			icon: 'question-circle',
			float: 'right',
			tooltip: 'FAQ'
		})

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
		maintenance: app.maintenance,
		embedded,
		request
	})
}