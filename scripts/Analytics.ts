export default class Analytics {
	push() {
		const analytics = {
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
			const connection = navigator["connection"] as any

			analytics.connection = {
				downLink: connection.downlink,
				roundTripTime: connection.rtt,
				effectiveType: connection.effectiveType
			}
		}

		fetch("/dark-flame-master", {
			method: "POST",
			credentials: "same-origin",
			body: JSON.stringify(analytics)
		})
	}
}
