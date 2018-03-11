import { AnimeNotifier } from "../AnimeNotifier"

var audioContext: AudioContext
var audioNode: AudioBufferSourceNode

// Play audio file
export function playAudio(arn: AnimeNotifier, button: HTMLButtonElement) {
	if(!audioContext) {
		audioContext = new AudioContext()
	}

	if(audioNode) {
		audioNode.stop()
	}

	let request = new XMLHttpRequest()
	request.open("GET", button.dataset.src, true)
	request.responseType = "arraybuffer"
	request.onload = () => {
		audioContext.decodeAudioData(request.response, buffer => {
			console.log("play")
			audioNode = audioContext.createBufferSource()
			audioNode.buffer = buffer
			audioNode.connect(audioContext.destination)
			audioNode.start(0)
		}, console.error)
	}
	request.send()

	// Show audio player
	document.getElementById("audio-player").classList.remove("fade-out")
	document.getElementById("audio-player-play").classList.add("fade-out")
	document.getElementById("audio-player-pause").classList.remove("fade-out")
}

// Pause audio
export function pauseAudio(arn: AnimeNotifier, button: HTMLButtonElement) {
	if(!audioNode) {
		return
	}

	audioNode.playbackRate.setValueAtTime(0.0, 0)

	document.getElementById("audio-player-play").classList.remove("fade-out")
	document.getElementById("audio-player-pause").classList.add("fade-out")
}

// Resume audio
export function resumeAudio(arn: AnimeNotifier, button: HTMLButtonElement) {
	if(!audioNode) {
		return
	}

	audioNode.playbackRate.setValueAtTime(1.0, 0)

	document.getElementById("audio-player-play").classList.add("fade-out")
	document.getElementById("audio-player-pause").classList.remove("fade-out")
}