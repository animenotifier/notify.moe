export default class Analytics {
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
			},
			connection: {
				downLink: 0,
				roundTripTime: 0,
				effectiveType: ""
			}
		}

		if("connection" in navigator) {
			analytics.connection = {
				downLink: navigator["connection"].downlink,
				roundTripTime: navigator["connection"].rtt,
				effectiveType: navigator["connection"].effectiveType
			}
		}

		fetch("/dark-flame-master", {
			method: "POST",
			credentials: "same-origin",
			body: JSON.stringify(analytics)
		})
	}
}