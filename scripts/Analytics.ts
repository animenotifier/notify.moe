export function uploadAnalytics() {
	const analytics = {
		general: {
			timezoneOffset: new Date().getTimezoneOffset()
		},
		screen: {
			availableHeight: screen.availHeight,
			availableWidth: screen.availWidth,
			height: screen.height,
			pixelRatio: window.devicePixelRatio,
			width: screen.width
		},
		system: {
			cpuCount: navigator.hardwareConcurrency,
			platform: navigator.platform
		},
		connection: {
			downLink: 0,
			effectiveType: "",
			roundTripTime: 0
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
