const prettyBytes = require('pretty-bytes')
const setRegex = /\| "arn" \| "use-default"\s* \| (\d+)\s* \| \d+\s* \| "(\w+)"\s* \| (\d+)\s* \|/g

exports.get = function*(request, response) {
	if(!arn.auth(request, response, 'admin'))
		return

	let user = request.user
	let statusText = yield arn.execute('aql -c \'show sets\'')
	let sets = []
	
	let match = null
	while(match = setRegex.exec(statusText)) {
		sets.push({
			name: match[2],
			objects: parseInt(match[1]),
			memoryUsage: parseInt(match[3])
		})
	}
	
	sets.sort((a, b) => a.name.localeCompare(b.name))

	response.render({
		user,
		sets,
		prettyBytes
	})
}