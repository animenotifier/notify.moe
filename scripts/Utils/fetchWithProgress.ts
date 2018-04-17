export function fetchWithProgress(url, options: RequestInit, onProgress: ((this: XMLHttpRequest, ev: ProgressEvent) => any) | null): Promise<string> {
	return new Promise((resolve, reject) => {
		let xhr = new XMLHttpRequest()
		xhr.open(options.method || "GET", url)

		for(let k in options.headers || {}) {
			xhr.setRequestHeader(k, options.headers[k])
		}

		xhr.onload = e => resolve(xhr.responseText)
		xhr.onerror = reject

		if(onProgress) {
			xhr.addEventListener("progress", onProgress)

			if(xhr.upload) {
				xhr.upload.onprogress = onProgress
			}
		}

		xhr.send(options.body)
	})
}
