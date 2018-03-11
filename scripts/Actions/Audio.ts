import { AnimeNotifier } from "../AnimeNotifier"

var audioContext: AudioContext
var audioNode: AudioBufferSourceNode
var playID = 0
var audioPlayer = document.getElementById("audio-player")
var audioPlayerPlay = document.getElementById("audio-player-play")
var audioPlayerPause = document.getElementById("audio-player-pause")

// Play audio file
export function playAudio(arn: AnimeNotifier, button: HTMLButtonElement) {
	if(!audioContext) {
		audioContext = new AudioContext()
	}

	playID++

	// Stop existing audioNode
	if(audioNode) {
		audioNode.stop()
		audioNode.disconnect()
		audioNode = null
	}

	// Request
	let request = new XMLHttpRequest()
	request.open("GET", button.dataset.audioSrc, true)
	request.responseType = "arraybuffer"

	request.onload = () => {
		audioContext.decodeAudioData(request.response, buffer => {
			audioNode = audioContext.createBufferSource()
			audioNode.buffer = buffer
			audioNode.connect(audioContext.destination)
			audioNode.start(0)

			let currentPlayCount = playID

			audioNode.onended = (event: MediaStreamErrorEvent) => {
				if(currentPlayCount === playID) {
					audioPlayer.classList.add("fade-out")
					audioNode.disconnect()
				}
			}
		}, console.error)
	}

	request.send()

	// Show audio player
	audioPlayer.classList.remove("fade-out")
	audioPlayerPlay.classList.add("fade-out")
	audioPlayerPause.classList.remove("fade-out")
}

// Pause audio
export function pauseAudio(arn: AnimeNotifier, button: HTMLButtonElement) {
	if(!audioNode) {
		return
	}

	audioNode.playbackRate.setValueAtTime(0.0, 0)

	audioPlayerPlay.classList.remove("fade-out")
	audioPlayerPause.classList.add("fade-out")
}

// Resume audio
export function resumeAudio(arn: AnimeNotifier, button: HTMLButtonElement) {
	if(!audioNode) {
		return
	}

	audioNode.playbackRate.setValueAtTime(1.0, 0)

	audioPlayerPlay.classList.add("fade-out")
	audioPlayerPause.classList.remove("fade-out")
}