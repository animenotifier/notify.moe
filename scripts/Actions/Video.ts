import AnimeNotifier from "../AnimeNotifier"

// Toggle play video
export function togglePlayVideo(arn: AnimeNotifier, element: HTMLElement) {
	const mediaId = element.dataset.mediaId

	if(!mediaId) {
		console.error("Missing data-media-id:", element)
		return
	}

	const container = document.getElementById(mediaId)

	if(!container) {
		console.error("Invalid data-media-id:", element)
		return
	}

	const video = container.getElementsByTagName("video")[0]
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
	const elementId = button.dataset.id

	if(!elementId) {
		console.error("Missing data-id:", button)
		return
	}

	const element = document.getElementById(elementId)

	if(!element) {
		console.error("Invalid data-id:", button)
		return
	}

	const requestFullscreen = element.requestFullscreen || element["mozRequestFullScreen"] || element["webkitRequestFullScreen"] || element["msRequestFullscreen"]
	const exitFullscreen = document.exitFullscreen || document["mozCancelFullScreen"] || document["webkitExitFullscreen"] || document["msExitFullscreen"]
	const fullscreen = document.fullscreen || document["webkitIsFullScreen"] || document["mozFullScreen"]

	if(fullscreen) {
		exitFullscreen.call(document)
	} else {
		requestFullscreen.call(element)
	}
}
