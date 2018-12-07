import AnimeNotifier from "../AnimeNotifier"

// Play video
export function playVideo(arn: AnimeNotifier, video: HTMLVideoElement) {
	video.volume = arn.audioPlayer.volume

	if(video.readyState >= 2) {
		togglePlayVideo(video)
		return
	}

	video.addEventListener("canplay", () => {
		togglePlayVideo(video)
	})

	video.load()
}

function togglePlayVideo(video: HTMLVideoElement) {
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

	if(document.fullscreen) {
		exitFullscreen.call(document)
	} else {
		requestFullscreen.call(element)
	}
}