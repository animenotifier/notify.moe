import AnimeNotifier from "../AnimeNotifier"

// Toggle play video
export function togglePlayVideo(arn: AnimeNotifier, element: HTMLElement) {
	let mediaId = element.dataset.mediaId

	if(!mediaId) {
		console.error("Missing data-media-id:", element)
		return
	}

	let container = document.getElementById(mediaId)

	if(!container) {
		console.error("Invalid data-media-id:", element)
		return
	}

	let video = container.getElementsByTagName("video")[0]
	video.volume = arn.audioPlayer.volume

	if(video.readyState >= 2) {
		togglePlayVideoElement(video)
		return
	}

	video.addEventListener("canplay", () => {
		togglePlayVideoElement(video)
	})

	video.load()
}

function togglePlayVideoElement(video: HTMLVideoElement) {
	if(video.paused) {
		video.play()
	} else {
		video.pause()
	}
}

// Toggle fullscreen
export function toggleFullscreen(_: AnimeNotifier, button: HTMLElement) {
	let elementId = button.dataset.id

	if(!elementId) {
		console.error("Missing data-id:", button)
		return
	}

	let element = document.getElementById(elementId)

	if(!element) {
		console.error("Invalid data-id:", button)
		return
	}

	let requestFullscreen = element.requestFullscreen || element["mozRequestFullScreen"] || element["webkitRequestFullScreen"] || element["msRequestFullscreen"]
	let exitFullscreen = document.exitFullscreen || document["mozCancelFullScreen"] || document["webkitExitFullscreen"] || document["msExitFullscreen"]
	let fullscreen = document.fullscreen || document["webkitIsFullScreen"] || document["mozFullScreen"]

	if(fullscreen) {
		exitFullscreen.call(document)
	} else {
		requestFullscreen.call(element)
	}
}