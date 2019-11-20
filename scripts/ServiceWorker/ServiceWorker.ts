// pack:ignore

// This is the service worker for notify.moe.
// When installed, it will intercept all requests made by the browser
// and return a cache-first response.

//         request                 request
// Browser -------> Service Worker ------> notify.moe Server
//         <-------
//         response (cache)
//                                 <------
//                                 response (network)
//         <-------
//         response (network)
//
// -> Diff cache with network response.

// By always returning cache first,
// we avoid latency problems on high latency connections like mobile
// networks. While the cache is being served, we start a real network
// request to the server to see if the resource changed. We compare the
// E-Tag of the cached and latest version of the resource. If the E-Tag
// of the current document changed, we send a message to the client
// that will cause the client to reload (diff) the page. It is not a real
// page reload as we will only calculate a DOM diff on the contents.
// If the style or script resources changed after being served, we need
// to force a real page reload.

// Promises
// const CACHEREFRESH = new Map<string, Promise<void>>()

// E-Tags that we served for a given URL
// const ETAGS = new Map<string, string>()

// MyCache is the cache used by the service worker.
class MyCache {
	version: string

	constructor(version: string) {
		this.version = version
	}

	clear() {
		return caches.delete(this.version)
	}

	async store(request: RequestInfo, response: Response) {
		try {
			// This can fail if the disk space quota has been exceeded.
			const cache = await caches.open(this.version)
			await cache.put(request, response)
		} catch(err) {
			console.log("Disk quota exceeded, can't store in cache:", request, response, err)
		}
	}

	async serve(request: RequestInfo): Promise<Response> {
		const cache = await caches.open(this.version)
		const matching = await cache.match(request)

		if(matching) {
			return matching
		}

		return Promise.reject("no-match")
	}
}

// Globals
let cache = new MyCache("v-7")
let reloads = new Map<string, Promise<Response>>()

// When these patterns are matched for the request URL, we exclude them from being
// served cache-first and instead serve them via a network request.
// Note that the service worker URL is automatically excluded from fetch events
// and therefore doesn't need to be added here.
let excludeCache = new Set<string>([
	"/api/",            // API requests
	"/paypal/",         // PayPal stuff
	"/import/",         // List imports
	"/from/",           // Infinite scrolling
	"chrome-extension", // Chrome extension

	// Authorization paths /auth/ and /logout are not listed here because they are handled in a special way.
])

// onInstall
async function onInstall(_: InstallEvent) {
	console.log("service worker install")

	// Skip waiting for the old service worker to shutdown
	await self.skipWaiting()

	// Install cache
	await installCache()
}

// onActivate
function onActivate(_: any) {
	console.log("service worker activate")

	// Only keep current version of the cache and delete old caches
	const cacheWhitelist = [cache.version]

	// Query existing cache keys
	const deleteOldCache = caches.keys().then(keyList => {
		// Create a deletion for every key that's not whitelisted
		const deletions = keyList.map(key => {
			if(cacheWhitelist.indexOf(key) === -1) {
				return caches.delete(key)
			}

			return Promise.resolve(false)
		})

		// Wait for deletions
		return Promise.all(deletions)
	})

	// Immediate claim helps us gain control over a new client immediately
	const immediateClaim = self.clients.claim()

	return Promise.all([
		deleteOldCache,
		immediateClaim
	])
}

// onRequest intercepts all browser requests.
// Simply returning, without calling evt.respondWith(),
// will let the browser deal with the request normally.
async function onRequest(evt: FetchEvent) {
	const request = evt.request as Request

	// If it's not a GET request, fetch it normally.
	// Let the browser handle XHR upload requests via POST,
	// so that we can receive upload progress events.
	if(request.method !== "GET") {
		return
	}

	// Video files are always loaded over the network.
	// We are defaulting to the normal browser handler here
	// so we can see the HTTP 206 partial responses in DevTools
	// and it also seems to have slightly smoother video playback.
	if(request.url.includes("/videos/")) {
		return
	}

	return //evt.respondWith(fetch(request))

	// // Exclude certain URLs from being cached.
	// for(let pattern of this.excludeCache.keys()) {
	// 	if(request.url.includes(pattern)) {
	// 		return
	// 	}
	// }

	// // If the request has cache set to "force-cache", return a cache-only response.
	// // This is used in reloads to avoid generating a 2nd request after a cache refresh.
	// if(request.headers.get("X-Force-Cache") === "true") {
	// 	return evt.respondWith(this.cache.serve(request))
	// }

	// // --------------------------------------------------------------------------------
	// // Cross-origin requests.
	// // --------------------------------------------------------------------------------

	// // These hosts don't support CORS. Always load via network.
	// if(request.url.startsWith("https://img.youtube.com/")) {
	// 	return
	// }

	// // Use CORS for cross-origin requests.
	// if(!request.url.startsWith("https://notify.moe/") && !request.url.startsWith("https://beta.notify.moe/")) {
	// 	request = new Request(request.url, {
	// 		credentials: "omit",
	// 		mode: "cors"
	// 	})
	// } else {
	// 	// let relativePath = trimPrefix(request.url, "https://notify.moe")
	// 	// relativePath = trimPrefix(relativePath, "https://beta.notify.moe")
	// 	// console.log(relativePath)
	// }

	// // --------------------------------------------------------------------------------
	// // Network refresh.
	// // --------------------------------------------------------------------------------

	// // Save response in cache.
	// let saveResponseInCache = (response: Response) => {
	// 	let contentType = response.headers.get("Content-Type")

	// 	// Don't cache anything other than text, styles, scripts, fonts and images.
	// 	if(!contentType.includes("text/") && !contentType.includes("application/javascript") && !contentType.includes("image/") && !contentType.includes("font/")) {
	// 		return response
	// 	}

	// 	// Save response in cache.
	// 	if(response.ok) {
	// 		let clone = response.clone()
	// 		this.cache.store(request, clone)
	// 	}

	// 	return response
	// }

	// let onResponse = (response: Response | null) => {
	// 	return response
	// }

	// // Refresh resource via a network request.
	// let refresh = fetch(request).then(saveResponseInCache)

	// // --------------------------------------------------------------------------------
	// // Final response.
	// // --------------------------------------------------------------------------------

	// // Clear cache on authentication and fetch it normally.
	// if(request.url.includes("/auth/") || request.url.includes("/logout")) {
	// 	return evt.respondWith(this.cache.clear().then(() => refresh))
	// }

	// // If the request has cache set to "no-cache",
	// // return the network-only response even if it fails.
	// if(request.headers.get("X-No-Cache") === "true") {
	// 	return evt.respondWith(refresh)
	// }

	// // Styles and scripts will be served via network first and fallback to cache.
	// if(request.url.endsWith("/styles") || request.url.endsWith("/scripts")) {
	// 	evt.respondWith(this.networkFirst(request, refresh, onResponse))
	// 	return refresh
	// }

	// // --------------------------------------------------------------------------------
	// // Default behavior for most requests.
	// // --------------------------------------------------------------------------------

	// // // Respond via cache first.
	// // evt.respondWith(this.cacheFirst(request, refresh, onResponse))
	// // return refresh

	// // Serve via network first and fallback to cache.
	// evt.respondWith(this.networkFirst(request, refresh, onResponse))
	// return refresh
}

// onMessage is called when the service worker receives a message from a client (browser tab).
async function onMessage(evt: ServiceWorkerMessageEvent) {
	const message = JSON.parse(evt.data)
	const clientId = (evt.source as any).id
	const client = await MyClient.get(clientId)

	client.onMessage(message)
}

// onPush is called on push events and requires the payload to contain JSON information about the notification.
function onPush(evt: PushEvent) {
	const payload = evt.data ? evt.data.json() : {}

	// Notify all clients about the new notification so they can update their notification counter
	broadcast({
		type: "new notification"
	})

	// Display the notification
	return self.registration.showNotification(payload.title, {
		body: payload.message,
		icon: payload.icon,
		data: payload.link,
		badge: "https://media.notify.moe/images/brand/256.png"
	})
}

async function onPushSubscriptionChange(evt: PushSubscriptionChangeEvent) {
	const userResponse = await fetch("/api/me", {credentials: "same-origin"})
	const user = await userResponse.json()

	await fetch(`/api/pushsubscriptions/${user.id}/remove`, {
		method: "POST",
		credentials: "same-origin",
		body: JSON.stringify({
			endpoint: evt.oldSubscription.endpoint
		})
	})

	const subscription = evt.newSubscription || await self.registration.pushManager.subscribe(evt.oldSubscription.options)

	if(!subscription || !subscription.endpoint) {
		return
	}

	const rawKey = subscription.getKey("p256dh")
	const key = rawKey ? btoa(String.fromCharCode.apply(null, new Uint8Array(rawKey))) : ""

	const rawSecret = subscription.getKey("auth")
	const secret = rawSecret ? btoa(String.fromCharCode.apply(null, new Uint8Array(rawSecret))) : ""

	const pushSubscription = {
		endpoint: subscription.endpoint,
		p256dh: key,
		auth: secret,
		platform: navigator.platform,
		userAgent: navigator.userAgent,
		screen: {
			width: window.screen.width,
			height: window.screen.height
		}
	}

	await fetch(`/api/pushsubscriptions/${user.id}/add`, {
		method: "POST",
		credentials: "same-origin",
		body: JSON.stringify(pushSubscription)
	})
}

// onNotificationClick is called when the user clicks on a notification.
function onNotificationClick(evt: NotificationEvent) {
	const notification = evt.notification
	notification.close()

	return self.clients.matchAll().then(function(clientList) {
		// If we have a link, use that link to open a new window.
		const url = notification.data

		if(url) {
			return self.clients.openWindow(url)
		}

		// If there is at least one client, focus it.
		if(clientList.length > 0) {
			return (clientList[0] as WindowClient).focus()
		}

		// Otherwise open a new window
		return self.clients.openWindow("https://notify.moe")
	})
}

// Broadcast sends a message to all clients (open tabs etc.)
function broadcast(msg: object) {
	const msgText = JSON.stringify(msg)

	self.clients.matchAll().then(function(clientList) {
		for(const client of clientList) {
			client.postMessage(msgText)
		}
	})
}

// installCache is called when the service worker is installed for the first time.
function installCache() {
	const urls = [
		"/",
		"/scripts",
		"/styles",
		"/manifest.json",
		"https://media.notify.moe/images/elements/noise-strong.png",
	]

	const requests = urls.map(async url => {
		const request = new Request(url, {
			credentials: "same-origin",
			mode: "cors"
		})

		const response = await fetch(request)
		await cache.store(request, response)
	})

	return Promise.all(requests)
}

// Serve network first.
// Fall back to cache.
async function networkFirst(request: Request, network: Promise<Response>, onResponse: (r: Response) => Response): Promise<Response> {
	let response: Response | null

	try {
		response = await network
		// console.log("Network HIT:", request.url)
	} catch(error) {
		// console.log("Network MISS:", request.url, error)

		try {
			response = await cache.serve(request)
		} catch(error) {
			return Promise.reject(error)
		}
	}

	return onResponse(response)
}

// Serve cache first.
// Fall back to network.
async function cacheFirst(request: Request, network: Promise<Response>, onResponse: (r: Response) => Response): Promise<Response> {
	let response: Response | null

	try {
		response = await cache.serve(request)
		// console.log("Cache HIT:", request.url)
	} catch(error) {
		// console.log("Cache MISS:", request.url, error)

		try {
			response = await network
		} catch(error) {
			return Promise.reject(error)
		}
	}

	return onResponse(response)
}

// MyClient represents a single tab in the browser.
class MyClient {
	// MyClient.idToClient is a Map of clients
	static idToClient = new Map<string, MyClient>()

	// MyClient.get retrieves a client by ID
	static async get(id: string): Promise<MyClient> {
		let client = MyClient.idToClient.get(id)

		if(!client) {
			client = new MyClient(await self.clients.get(id))
			MyClient.idToClient.set(id, client)
		}

		return client
	}

	// The actual client
	client: ServiceWorkerClient

	constructor(client: ServiceWorkerClient) {
		this.client = client
	}

	onMessage(message: any) {
		switch(message.type) {
			case "loaded":
				this.onDOMContentLoaded(message.url)
				break

			case "broadcast":
				message.type = message.realType
				delete message.realType
				broadcast(message)
				break
		}
	}

	// postMessage sends a message to the client.
	postMessage(message: object) {
		this.client.postMessage(JSON.stringify(message))
	}

	// onDOMContentLoaded is called when the client sent this service worker
	// a message that the page has been loaded.
	onDOMContentLoaded(_: string) {
		// ...
	}
}

// trimPrefix removes the prefix from the text.
function trimPrefix(text, prefix) {
	if(text.startsWith(prefix)) {
		return text.slice(prefix.length)
	}

	return text
}

// Register event listeners
self.addEventListener("install", (evt: InstallEvent) => evt.waitUntil(onInstall(evt)))
self.addEventListener("activate", (evt: any) => evt.waitUntil(onActivate(evt)))
self.addEventListener("fetch", (evt: FetchEvent) => evt.waitUntil(onRequest(evt)))
self.addEventListener("message", (evt: any) => evt.waitUntil(onMessage(evt)))
self.addEventListener("push", (evt: PushEvent) => evt.waitUntil(onPush(evt)))
self.addEventListener("pushsubscriptionchange", (evt: any) => evt.waitUntil(onPushSubscriptionChange(evt)))
self.addEventListener("notificationclick", (evt: NotificationEvent) => evt.waitUntil(onNotificationClick(evt)))
