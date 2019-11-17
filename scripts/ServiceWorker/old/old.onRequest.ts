// pack:ignore

// // Clear cache on authentication and fetch it normally
// if(request.url.includes("/auth/") || request.url.includes("/logout")) {
// 	return evt.respondWith(caches.delete(this.cache.version).then(() => fetch(request)))
// }

// // Exclude certain URLs from being cached
// for(let pattern of EXCLUDECACHE.keys()) {
// 	if(request.url.includes(pattern)) {
// 		return evt.respondWith(this.fromNetwork(request))
// 	}
// }

// // If the request included the header "X-CacheOnly", return a cache-only response.
// // This is used in reloads to avoid generating a 2nd request after a cache refresh.
// if(request.headers.get("X-CacheOnly") === "true") {
// 	return evt.respondWith(this.fromCache(request))
// }

// // Save the served E-Tag when onResponse is called
// let servedETag = undefined

// let onResponse = (response: Response | null) => {
// 	if(response) {
// 		servedETag = response.headers.get("ETag")
// 		ETAGS.set(request.url, servedETag)
// 	}

// 	return response
// }

// let saveResponseInCache = response => {
// 	let clone = response.clone()

// 	// Save the new version of the resource in the cache
// 	let cacheRefresh = this.cache.store(request, clone).catch(err => {
// 		console.error(err)
// 		// TODO: Tell client that the quota is exceeded (disk full).
// 	})

// 	CACHEREFRESH.set(request.url, cacheRefresh)
// 	return response
// }

// // Start fetching the request
// let network =
// 	fetch(request)
// 	.then(saveResponseInCache)
// 	.catch(error => {
// 		console.log("Fetch error:", error)
// 		throw error
// 	})

// // Save in map
// this.reloads.set(request.url, network)

// if(request.headers.get("X-Reload") === "true") {
// 	return evt.respondWith(network)
// }

// // Scripts and styles are server pushed on the initial response
// // so we can use a network-first response without an additional round-trip.
// // This causes the browser to always load the most recent scripts and styles.
// if(request.url.endsWith("/styles") || request.url.endsWith("/scripts")) {
// 	return evt.respondWith(this.networkFirst(request, network, onResponse))
// }

// return evt.respondWith(this.cacheFirst(request, network, onResponse))
