// pack:ignore

// onDOMContentLoaded(url: string) {
	// let refresh = serviceWorker.reloads.get(url)
	// let servedETag = ETAGS.get(url)

	// // If the user requests a sub-page we should prefetch the full page, too.
	// if(url.includes("/_/") && !url.includes("/_/search/")) {
	// 	var prefetch = true

	// 	for(let pattern of EXCLUDECACHE.keys()) {
	// 		if(url.includes(pattern)) {
	// 			prefetch = false
	// 			break
	// 		}
	// 	}

	// 	if(prefetch) {
	// 		this.prefetchFullPage(url)
	// 	}
	// }

	// if(!refresh || !servedETag) {
	// 	return Promise.resolve()
	// }

	// return refresh.then(async (response: Response) => {
	// 	// When the actual network request was used by the client, response.bodyUsed is set.
	// 	// In that case the client is already up to date and we don"t need to tell the client to do a refresh.
	// 	if(response.bodyUsed) {
	// 		return
	// 	}

	// 	// Get the ETag of the cached response we sent to the client earlier.
	// 	let eTag = response.headers.get("ETag")

	// 	// Update ETag
	// 	ETAGS.set(url, eTag)

	// 	// If the ETag changed, we need to do a reload.
	// 	if(eTag !== servedETag) {
	// 		return this.reloadContent(url)
	// 	}

	// 	// Do nothing
	// 	return Promise.resolve()
	// })
// }

// prefetchFullPage(url: string) {
// 	let fullPage = new Request(url.replace("/_/", "/"))

// 	let fullPageRefresh = fetch(fullPage, {
// 		credentials: "same-origin"
// 	}).then(response => {
// 		// Save the new version of the resource in the cache
// 		let cacheRefresh = caches.open(serviceWorker.cache.version).then(cache => {
// 			return cache.put(fullPage, response)
// 		})

// 		CACHEREFRESH.set(fullPage.url, cacheRefresh)
// 		return response
// 	})

// 	// Save in map
// 	serviceWorker.reloads.set(fullPage.url, fullPageRefresh)
// }

// async reloadContent(url: string) {
// 	let cacheRefresh = CACHEREFRESH.get(url)

// 	if(cacheRefresh) {
// 		await cacheRefresh
// 	}

// 	return this.postMessage({
// 		type: "new content",
// 		url
// 	})
// }

// async reloadPage(url: string) {
// 	let networkFetch = serviceWorker.reloads.get(url.replace("/_/", "/"))

// 	if(networkFetch) {
// 		await networkFetch
// 	}

// 	return this.postMessage({
// 		type: "reload page",
// 		url
// 	})
// }

// reloadStyles() {
// 	return this.postMessage({
// 		type: "reload styles"
// 	})
// }
