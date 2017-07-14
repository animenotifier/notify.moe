// pack:ignore

const CACHE = "v-1"
const RELOADS = new Map<string, Promise<Response>>()
const ETAGS = new Map<string, string>()

self.addEventListener("install", (evt: any) => {
	console.log("Service worker install")

	evt.waitUntil(
		(self as any).skipWaiting().then(() => {
			return installCache()
		})
	)
})

self.addEventListener("activate", (evt: any) => {
	console.log("Service worker activate")

	evt.waitUntil(
		(self as any).clients.claim()
	)
})

// controlling service worker
self.addEventListener("message", (evt: any) => {
	let message = JSON.parse(evt.data)
	
	let url = message.url
	let refresh = RELOADS.get(url)
	let servedETag = ETAGS.get(url)

	if(!refresh || !servedETag) {
		return
	}
	
	evt.waitUntil(
		refresh.then((response: Response) => {
			// If the fresh copy was used to serve the request instead of the cache,
			// we don't need to tell the client to do a refresh.
			if(response.bodyUsed) {
				return
			}

			let eTag = response.headers.get("ETag")

			if(eTag === servedETag) {
				return
			}

			ETAGS.set(url, eTag)

			let message = {
				type: "new content",
				url
			}

			return evt.source.postMessage(JSON.stringify(message))
		})
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

	let servedETag = undefined
	
	// Start fetching the request
	let refresh = fetch(request).then(response => {
		let clone = response.clone()

		// Save the new version of the resource in the cache
		caches.open(CACHE).then(cache => {
			return cache.put(request, clone)
		})

		return response
	})

	// Save in map
	RELOADS.set(request.url, refresh)

	// Forced reload
	if(request.headers.get("X-Reload") === "true") {
		return evt.waitUntil(refresh)
	}

	// Try to serve cache first and fall back to network response
	let networkOrCache = fromCache(request).then(response => {
		servedETag = response.headers.get("ETag")
		ETAGS.set(request.url, servedETag)
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
