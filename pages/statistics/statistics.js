let increment = function(obj, key) {
	if(obj.hasOwnProperty(key))
		obj[key] += 1
	else
		obj[key] = 1
}

exports.get = function*(request, response) {
	let recordCount = 0
	let notificationsEnabled = 0

	let gender = {
		male: 0,
		female: 0,
		unknown: 0
	}

	let providers = {
		list: {},
		anime: {},
		airingDate: {}
	}

	let titleLanguages = {
		romaji: 0,
		english: 0,
		japanese: 0
	}

	let sortBy = {
		alphabetically: 0,
		airingDate: 0
	}

	let browsers = {}

	yield arn.forEach('Users', function(user) {
		if(!arn.isActiveUser(user))
			return

		if(user.gender === 'male' || user.gender === 'female')
			gender[user.gender] += 1
		//else
		//	gender.unknown += 1

		titleLanguages[user.titleLanguage] += 1
		sortBy[user.sortBy] += 1

		if(Object.keys(user.pushEndpoints).length > 0)
			notificationsEnabled += 1

		if(user.agent && user.agent.family)
			increment(browsers, user.agent.family)

		increment(providers.list, user.providers.list)
		increment(providers.anime, user.providers.anime)
		increment(providers.airingDate, user.providers.airingDate)

		recordCount++
	})
	
	browsers.Others = 0
	
	const onePercentMark = recordCount / 100
	
	Object.keys(browsers).forEach(key => {
		if(key === 'Others')
			return
		
		if(browsers[key] < onePercentMark) {
			browsers.Others += browsers[key]
			delete browsers[key]
		}
	})

	response.render({
		users: {
			total: recordCount,
			gender,
			genderSpecified: gender.male + gender.female,
			titleLanguages,
			notificationsEnabled,
			browsers,
			sortBy
		},
		anime: {
			total: arn.animeCount
		},
		providers
	})
}