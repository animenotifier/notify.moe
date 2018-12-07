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

	let progressElement = video.parentElement.getElementsByClassName("video-current-progress")[0] as HTMLElement
	let timeElement = video.parentElement.getElementsByClassName("video-time")[0]

	video.addEventListener("timeupdate", () => {
		let time = video.currentTime
		let minutes = Math.trunc(time / 60)
		let seconds = Math.trunc(time) % 60
		let paddedSeconds = ("00" + seconds).slice(-2)
		timeElement.textContent = `${minutes}:${paddedSeconds}`
		progressElement.style.transform = `scaleX(${time / video.duration})`
	})

	video.addEventListener("waiting", () => {
		arn.statusMessage.showInfo("Buffering...", -1)
	})

	video.addEventListener("playing", () => {
		arn.statusMessage.close()
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

	if(document.fullscreen) {
		document.exitFullscreen()
	} else {
		element.requestFullscreen()
	}
}