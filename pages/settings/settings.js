'use strict'

exports.get = function(request, response) {
	let user = request.user

	response.render({
		user,
		viewUser: user,
		getProperty: function(obj, desc) {
			let arr = desc.split('.')

			while(arr.length && obj)
				obj = obj[arr.shift()]

			return obj
		},
		canEdit: true
	})
}