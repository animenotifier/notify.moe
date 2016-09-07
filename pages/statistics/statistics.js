let increment = function(obj, key) {
	if(obj.hasOwnProperty(key))
		obj[key] += 1
	else
		obj[key] = 1
}

let filterLessThan = function(map, threshold) {
	map.Others = 0
	
	Object.keys(map).forEach(key => {
		if(key === 'Others')
			return
		
		if(map[key] < threshold) {
			map.Others += map[key]
			delete map[key]
		}
	})
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
	let countries = {}

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
			
		if(user.location && user.location.countryName && user.location.countryName != '-')
			increment(countries, user.location.countryName)
		else
			increment(countries, 'Unknown')

		increment(providers.list, user.providers.list)
		increment(providers.anime, user.providers.anime)
		increment(providers.airingDate, user.providers.airingDate)

		recordCount++
	})
	
	const onePercentMark = recordCount / 100
	
	filterLessThan(browsers, onePercentMark)
	filterLessThan(countries, onePercentMark * 4)

	response.render({
		users: {
			total: recordCount,
			gender,
			genderSpecified: gender.male + gender.female,
			titleLanguages,
			notificationsEnabled,
			browsers,
			countries,
			sortBy
		},
		anime: {
			total: arn.animeCount
		},
		providers
	})
}
