// import AnimeNotifier from "../AnimeNotifier"

// // User map
// export function showUserMap(arn: AnimeNotifier, button: HTMLElement) {
// 	let script = document.createElement("script")
// 	script.src = "https://www.gstatic.com/charts/loader.js"
// 	script.onload = () => {
// 		let google = window["google"]
// 		google.charts.load("current", {
// 			"packages": ["geochart"],
// 			"mapsApiKey": ""
// 		})
// 		google.charts.setOnLoadCallback(drawUserMap)
// 	}

// 	document.head.appendChild(script)
// }

// function drawUserMap() {
// 	let data = window["google"].visualization.arrayToDataTable([
// 		["Country", "Popularity"],
// 		["Germany", 200],
// 		["United States", 300],
// 		["Brazil", 400],
// 		["Canada", 500],
// 		["France", 600],
// 		["RU", 700]
// 	])

// 	let options = {}
// 	let element = document.getElementById("user-map")
// 	let chart = new window["google"].visualization.GeoChart(element)
// 	console.log(element, data)
// 	chart.draw(data, options)
// }
