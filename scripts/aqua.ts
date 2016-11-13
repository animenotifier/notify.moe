function find(id: string) {
	return document.getElementById(id)
}

function get(url: string, body?: Object): Promise<string> {
	return new Promise(function(resolve, reject) {
		resolve("")
	})
}

function post(url: string, body?: Object): Promise<string> {
	return new Promise(function(resolve, reject) {
		resolve("")
	})
}

export {
	find,
	get,
	post
}