import AnimeNotifier from "../AnimeNotifier"

// Play video
export function playVideo(arn: AnimeNotifier, video: HTMLVideoElement) {
	video.volume = arn.audioPlayer.volume

	if(video.readyState >= 2) {
		togglePlayVideo(video)
		return
	}

	video.load()

	video.addEventListener("loadeddata", () => {
		togglePlayVideo(video)
	})
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

	if(document.fullscreen) {
		document.exitFullscreen()
	} else {
		element.requestFullscreen()
	}
}