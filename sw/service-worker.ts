// pack:ignore

var CACHE = "v-alpha"

self.addEventListener("install", (evt: any) => {
	console.log("The service worker is being installed.")

	evt.waitUntil(installCache())
})

self.addEventListener("activate", (evt: any) => {
	evt.waitUntil((self as any).clients.claim())
})

self.addEventListener("fetch", async (evt: any) => {
	let request = evt.request

	console.log("Serving:", request.url, request, request.method)

	// Do not use cache in some cases
	if(request.method !== "GET" || request.url.includes("/auth/") || request.url.includes("chrome-extension")) {
		return evt.waitUntil(evt.respondWith(fetch(request)))
	}
	
	// Start fetching the request
	let refresh = fetch(request).then(response => {
		let clone = response.clone()

		// Save the new version of the resource in the cache
		caches.open(CACHE).then(cache => {
			cache.put(request, clone)
		})

		return response
	})

	// Try to serve cache first and fall back to network response
	let networkOrCache = fromCache(request).catch(error => {
		console.log("Cache MISS:", request.url)
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
				console.log("Cache HIT:", request.url)
				return Promise.resolve(matching)
			}

			return Promise.reject("no-match")
		})
	})
}

function updateCache(request, response) {
	return caches.open(CACHE).then(cache => {
		cache.put(request, response)
	})
}
