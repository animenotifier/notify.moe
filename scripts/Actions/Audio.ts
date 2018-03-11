import { AnimeNotifier } from "../AnimeNotifier"

var audioContext: AudioContext
var audioNode: AudioBufferSourceNode
var playID = 0
var audioPlayer = document.getElementById("audio-player")
var audioPlayerPlay = document.getElementById("audio-player-play")
var audioPlayerPause = document.getElementById("audio-player-pause")

// Play audio file
export function playAudio(arn: AnimeNotifier, element: HTMLElement) {
	if(!audioContext) {
		audioContext = new AudioContext()
	}

	playID++

	// Stop current track
	stopAudio(arn)

	arn.currentSoundTrackId = element.dataset.soundtrackId
	element.classList.add("playing")

	// Request
	let request = new XMLHttpRequest()
	request.open("GET", element.dataset.audioSrc, true)
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
					stopAudio(arn)
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

// Stop audio
export function stopAudio(arn: AnimeNotifier) {
	if(!audioNode) {
		return
	}

	audioNode.stop()
	audioNode.disconnect()
	audioNode = null

	arn.currentSoundTrackId = undefined

	// Remove CSS class "playing"
	let playingElements = document.getElementsByClassName("playing")

	for(let playing of playingElements) {
		playing.classList.remove("playing")
	}

	// Fade out sidebar player
	audioPlayer.classList.add("fade-out")
}

// Toggle audio
export function toggleAudio(arn: AnimeNotifier, element: HTMLElement) {
	// If we're clicking on the same track again, stop playing.
	// Otherwise, start the track we clicked on.
	if(arn.currentSoundTrackId && element.dataset.soundtrackId === arn.currentSoundTrackId) {
		stopAudio(arn)
	} else {
		playAudio(arn, element)
	}
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