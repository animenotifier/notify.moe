export default function uploadWithProgress(url, options: RequestInit, onProgress: ((ev: ProgressEvent) => any) | null): Promise<string> {
	return new Promise((resolve, reject) => {
		const xhr = new XMLHttpRequest()

		xhr.onload = () => {
			if(xhr.status >= 400) {
				return reject(xhr.responseText)
			}

			resolve(xhr.responseText)
		}
		xhr.onerror = reject

		if(onProgress && xhr.upload) {
			xhr.upload.onprogress = onProgress
			xhr.upload.onerror = reject
		} else {
			console.error("Could not attach progress event listener")
		}

		xhr.open(options.method || "GET", url, true)

		if(options.headers) {
			for(const key of Object.keys(options.headers)) {
				xhr.setRequestHeader(key, options.headers[key])
			}
		}

		xhr.send(options.body)
	})
}
