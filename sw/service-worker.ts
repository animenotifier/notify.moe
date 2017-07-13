// pack:ignore

var CACHE = "v-1"

self.addEventListener("install", (evt: any) => {
	evt.waitUntil(
		(self as any).skipWaiting().then(() => {
			return installCache()
		})
	)
})

self.addEventListener("activate", (evt: any) => {
	evt.waitUntil(
		(self as any).clients.claim()
	)
})

self.addEventListener("fetch", async (evt: any) => {
	let request = evt.request
	let isAuth = request.url.includes("/auth/") || request.url.includes("/logout")

	// Delete existing cache on authentication
	if(isAuth) {
		caches.delete(CACHE)
	}

	// Do not use cache in some cases
	if(request.method !== "GET" || isAuth || request.url.includes("chrome-extension")) {
		return evt.waitUntil(evt.respondWith(fetch(request)))
	}

	let servedCachedResponse = false
	
	// Start fetching the request
	let refresh = fetch(request).then(response => {
		let clone = response.clone()

		// Save the new version of the resource in the cache
		caches.open(CACHE).then(cache => {
			return cache.put(request, clone)
		}).then(() => {
			if(!servedCachedResponse) {
				return
			}

			let contentType = clone.headers.get("Content-Type")

			if(contentType && contentType.startsWith("text/html") && clone.headers.get("ETag") && request.headers.get("X-Reload") !== "true") {
				reloadContent(clone)
			}
		})

		return response
	})

	// Forced reload
	if(request.headers.get("X-Reload") === "true") {
		return evt.waitUntil(refresh)
	}

	// Try to serve cache first and fall back to network response
	let networkOrCache = fromCache(request).then(response => {
		servedCachedResponse = true
		return response
	}).catch(error => {
		// console.log("Cache MISS:", request.url)
		return refresh
	})

	return evt.waitUntil(evt.respondWith(networkOrCache))
})

function installCache() {
	return caches.open(CACHE).then(cache => {
		return cache.addAll([
			"./",
			"./scripts",
			"https://fonts.gstatic.com/s/ubuntu/v10/2Q-AW1e_taO6pHwMXcXW5w.ttf"
		])
	})
}

function fromCache(request) {
	return caches.open(CACHE).then(cache => {
		return cache.match(request).then(matching => {
			if(matching) {
				// console.log("Cache HIT:", request.url)
				return Promise.resolve(matching)
			}

			return Promise.reject("no-match")
		})
	})
}

function reloadContent(response) {
	return (self as any).clients.matchAll().then(clients => {
		clients.forEach(client => {
			var message = {
				type: 'content changed',
				url: response.url,
				eTag: response.headers.get('ETag')
			}

			client.postMessage(JSON.stringify(message))
		})
	})
}
