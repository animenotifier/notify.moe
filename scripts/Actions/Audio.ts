import { AnimeNotifier } from "../AnimeNotifier"

var audioContext: AudioContext
var audioNode: AudioBufferSourceNode
var gainNode: GainNode
var volume = 0.5
var volumeTimeConstant = 0.01
var volumeSmoothingDelay = 0.05
var playId = 0
var audioPlayer = document.getElementById("audio-player")
var audioPlayerPlay = document.getElementById("audio-player-play")
var audioPlayerPause = document.getElementById("audio-player-pause")
var trackLink = document.getElementById("audio-player-track-title") as HTMLLinkElement
var animeInfo = document.getElementById("audio-player-anime-info") as HTMLElement
var animeLink = document.getElementById("audio-player-anime-link") as HTMLLinkElement
var animeImage = document.getElementById("audio-player-anime-image") as HTMLImageElement

// Play audio
export function playAudio(arn: AnimeNotifier, element: HTMLElement) {
	playAudioFile(arn, element.dataset.soundtrackId, element.dataset.audioSrc)
}

// Play audio file
function playAudioFile(arn: AnimeNotifier, trackId: string, trackUrl: string) {
	if(!audioContext) {
		audioContext = new AudioContext()
		gainNode = audioContext.createGain()
		gainNode.gain.setTargetAtTime(volume, audioContext.currentTime + volumeSmoothingDelay, volumeTimeConstant)
	}

	playId++
	let currentPlayId = playId

	// Stop current track
	stopAudio(arn)

	arn.currentSoundTrackId = trackId
	arn.markPlayingSoundTrack()
	arn.loading(true)

	// Request
	let request = new XMLHttpRequest()
	request.open("GET", trackUrl, true)
	request.responseType = "arraybuffer"

	request.onload = () => {
		arn.loading(false)

		if(currentPlayId !== playId) {
			return
		}

		audioContext.decodeAudioData(request.response, async buffer => {
			if(currentPlayId !== playId) {
				return
			}

			audioNode = audioContext.createBufferSource()
			audioNode.buffer = buffer
			audioNode.connect(gainNode)
			gainNode.connect(audioContext.destination)
			audioNode.start(0)

			audioNode.onended = (event: MediaStreamErrorEvent) => {
				if(currentPlayId !== playId) {
					return
				}

				playNextTrack(arn)
				// stopAudio(arn)
			}

			// Set track title
			let trackInfoResponse = await fetch("/api/soundtrack/" + trackId)
			let track = await trackInfoResponse.json()
			trackLink.href = "/soundtrack/" + track.id
			trackLink.innerText = track.title

			let animeId = ""

			for(let tag of (track.tags as string[])) {
				if(tag.startsWith("anime:")) {
					animeId = tag.split(":")[1]
					break
				}
			}

			// Set anime info
			if(animeId !== "") {
				animeInfo.classList.remove("hidden")
				let animeResponse = await fetch("/api/anime/" + animeId)
				let anime = await animeResponse.json()
				animeLink.title = anime.title.canonical
				animeLink.href = "/anime/" + anime.id
				animeImage.dataset.src = "//media.notify.moe/images/anime/medium/" + anime.id + anime.imageExtension
				animeImage.classList.remove("hidden")
				animeImage["became visible"]()
			}
		}, console.error)
	}

	request.onerror = () => {
		arn.loading(false)
	}

	request.send()

	// Show audio player
	audioPlayer.classList.remove("fade-out")
	audioPlayerPlay.classList.add("fade-out")
	audioPlayerPause.classList.remove("fade-out")
}

// Stop audio
export function stopAudio(arn: AnimeNotifier) {
	arn.currentSoundTrackId = undefined

	// Remove CSS class "playing"
	let playingElements = document.getElementsByClassName("playing")

	for(let playing of playingElements) {
		playing.classList.remove("playing")
	}

	// Fade out sidebar player
	// audioPlayer.classList.add("fade-out")

	// Remove title
	trackLink.href = ""
	trackLink.innerText = ""

	// Hide anime info
	animeLink.href = ""
	animeInfo.classList.add("hidden")
	animeImage.classList.add("hidden")

	if(gainNode) {
		gainNode.disconnect()
	}

	if(audioNode) {
		audioNode.stop()
		audioNode.disconnect()
		audioNode = null
	}
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

// Play previous track
export async function playPreviousTrack(arn: AnimeNotifier) {
	alert("Previous track is currently work in progress! Check back later :)")
}

// Play next track
export async function playNextTrack(arn: AnimeNotifier) {
	// Get random track
	let response = await fetch("/api/next/soundtrack")
	let track = await response.json()

	playAudioFile(arn, track.id, "https://notify.moe/audio/" + track.file)
	arn.statusMessage.showInfo("Now playing: " + track.title)

	return track
}

// Set volume
export function setVolume(arn: AnimeNotifier, element: HTMLInputElement) {
	volume = parseFloat(element.value) / 100.0

	if(gainNode) {
		gainNode.gain.setTargetAtTime(volume, audioContext.currentTime + volumeSmoothingDelay, volumeTimeConstant)
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
		playNextTrack(arn)
		return
	}

	audioNode.playbackRate.setValueAtTime(1.0, 0)

	audioPlayerPlay.classList.add("fade-out")
	audioPlayerPause.classList.remove("fade-out")
}