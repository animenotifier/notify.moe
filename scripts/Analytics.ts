export class Analytics {
	push() {
		let analytics = {
			general: {
				timezoneOffset: new Date().getTimezoneOffset()
			},
			screen: {
				width: screen.width,
				height: screen.height,
				availableWidth: screen.availWidth,
				availableHeight: screen.availHeight,
				pixelRatio: window.devicePixelRatio
			},
			system: {
				cpuCount: navigator.hardwareConcurrency,
				platform: navigator.platform
			}
		}

		fetch("/dark-flame-master", {
			method: "POST",
			credentials: "same-origin",
			body: JSON.stringify(analytics)
		})
	}
}