// pack:ignore

const CACHE = "v-1"
const RELOADS = new Map<string, Promise<Response>>()
const ETAGS = new Map<string, string>()
const CACHEREFRESH = new Map<string, Promise<void>>()

self.addEventListener("install", (evt: InstallEvent) => {
	console.log("Service worker install")

	evt.waitUntil(
		(self as any).skipWaiting().then(() => {
			return installCache()
		})
	)
})

self.addEventListener("activate", (evt: any) => {
	console.log("Service worker activate")

	// Delete old cache
	let cacheWhitelist = [CACHE]

	let deleteOldCache = caches.keys().then(keyList => {
		return Promise.all(keyList.map(key => {
			if(cacheWhitelist.indexOf(key) === -1) {
				return caches.delete(key)
			}
		}))
	})

	let immediateClaim = (self as any).clients.claim()

	// Immediate claim
	evt.waitUntil(
		Promise.all([
			deleteOldCache,
			immediateClaim
		])
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
			// we don"t need to tell the client to do a refresh.
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

			// If a subpage has refreshed, refresh the main page cache, too.
			// if(url.includes("/_/")) {
				
			// }

			let cacheRefresh = CACHEREFRESH.get(url)

			if(!cacheRefresh) {
				return evt.source.postMessage(JSON.stringify(message))
			}

			return cacheRefresh.then(() => {
				evt.source.postMessage(JSON.stringify(message))
			})
		})
	)
})

self.addEventListener("fetch", async (evt: FetchEvent) => {
	let request = evt.request as Request
	let isAuth = request.url.includes("/auth/") || request.url.includes("/logout")
	let ignoreCache = request.url.includes("/api/") || request.url.includes("/paypal/") || request.url.includes("chrome-extension")

	// Delete existing cache on authentication
	if(isAuth) {
		caches.delete(CACHE)
	}

	// Do not use cache in some cases
	if(request.method !== "GET" || isAuth || ignoreCache) {
		return evt.waitUntil(evt.respondWith(fetch(request)))
	}

	let servedETag = undefined
	
	// Start fetching the request
	let refresh = fetch(request).then(response => {
		let clone = response.clone()

		// Save the new version of the resource in the cache
		let cacheRefresh = caches.open(CACHE).then(cache => {
			return cache.put(request, clone)
		})

		CACHEREFRESH.set(request.url, cacheRefresh)

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

self.addEventListener("push", (evt: PushEvent) => {
	var payload = evt.data ? evt.data.json() : {}

	evt.waitUntil(
		(self as any).registration.showNotification(payload.title, {
			body: payload.message,
			icon: payload.icon,
			image: payload.image,
			data: payload.link
		})
	)
})

self.addEventListener("pushsubscriptionchange", (evt: any) => {
	console.log("pushsubscriptionchange", evt)
})

self.addEventListener("notificationclick", (evt: NotificationEvent) => {
	let notification = evt.notification
	notification.close()

	evt.waitUntil(
		(self as any).clients.matchAll().then(function(clientList) {
			// If we have a link, use that link to open a new window.
			let url = notification.data

			if(url) {
				return (self as any).clients.openWindow(url)
			}

			// If there is at least one client, focus it.
			if(clientList.length > 0) {
				return clientList[0].focus()
			}

			// Otherwise open a new window
			return (self as any).clients.openWindow("https://notify.moe")
		})
	)
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
