export function uploadWithProgress(url, options: RequestInit, onProgress: ((ev: ProgressEvent) => any) | null): Promise<string> {
	return new Promise((resolve, reject) => {
		let xhr = new XMLHttpRequest()

		xhr.onload = e => resolve(xhr.responseText)
		xhr.onerror = reject

		if(onProgress && xhr.upload) {
			xhr.upload.onprogress = onProgress
			xhr.upload.onerror = reject
		} else {
			console.error("Could not attach progress event listener")
		}

		xhr.open(options.method || "GET", url, true)

		for(let k in options.headers || {}) {
			xhr.setRequestHeader(k, options.headers[k])
		}

		xhr.send(options.body)
	})
}
