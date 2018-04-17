export function fetchWithProgress(url, options: RequestInit, onProgress: ((this: XMLHttpRequest, ev: ProgressEvent) => any) | null): Promise<string> {
	return new Promise((resolve, reject) => {
		let xhr = new XMLHttpRequest()

		xhr.addEventListener("load", e => resolve(xhr.responseText))
		xhr.addEventListener("error", reject)

		if(onProgress && xhr.upload) {
			xhr.upload.addEventListener("progress", onProgress)
		}

		xhr.open(options.method || "GET", url, true)

		for(let k in options.headers || {}) {
			xhr.setRequestHeader(k, options.headers[k])
		}

		xhr.send(options.body)
	})
}
