import AnimeNotifier from "../AnimeNotifier"

// Toggle play video
export function togglePlayVideo(arn: AnimeNotifier, element: HTMLElement) {
	let container = document.getElementById(element.dataset.mediaId)
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
export function toggleFullscreen(arn: AnimeNotifier, button: HTMLElement) {
	let elementId = button.dataset.id
	let element = document.getElementById(elementId)
	let requestFullscreen = element.requestFullscreen || element["mozRequestFullScreen"] || element["webkitRequestFullScreen"] || element["msRequestFullscreen"]
	let exitFullscreen = document.exitFullscreen || document["mozCancelFullScreen"] || document["webkitExitFullscreen"] || document["msExitFullscreen"]
	let fullscreen = document.fullscreen || document["webkitIsFullScreen"] || document["mozFullScreen"]

	if(fullscreen) {
		exitFullscreen.call(document)
	} else {
		requestFullscreen.call(element)
	}
}