export function fetchWithProgress(url, opts: RequestInit, onProgress: ((this: XMLHttpRequest, ev: ProgressEvent) => any) | null): Promise<string> {
	return new Promise((resolve, reject) => {
		let xhr = new XMLHttpRequest()
		xhr.open(opts.method || "GET", url)

		for(let k in opts.headers || {}) {
			xhr.setRequestHeader(k, opts.headers[k])
		}

		xhr.onload = e => resolve(xhr.responseText)
		xhr.onerror = reject

		if(xhr.upload && onProgress) {
			xhr.upload.onprogress = onProgress
		}

		xhr.send(opts.body)
	})
}
