import AnimeNotifier from "./AnimeNotifier"
import { Anime } from "./Types/Anime"

export default class AudioPlayer {
	arn: AnimeNotifier

	// Web audio
	audioContext: AudioContext
	audioNode: AudioBufferSourceNode | null
	gainNode: GainNode

	// Parameters
	volume = 0.5
	volumeTimeConstant = 0.01
	volumeSmoothingDelay = 0.05
	targetSpeed = 1.0
	playId = 0

	// Save last request so that we can cancel it
	lastRequest: XMLHttpRequest | null

	// DOM elements
	audioPlayer: HTMLElement
	audioPlayerPlay: HTMLButtonElement
	audioPlayerPause: HTMLButtonElement
	trackLink: HTMLLinkElement
	animeInfo: HTMLElement
	animeLink: HTMLLinkElement
	animeImage: HTMLImageElement

	constructor(arn: AnimeNotifier) {
		this.arn = arn
		this.audioPlayer = document.getElementById("audio-player") as HTMLElement
		this.audioPlayerPlay = document.getElementById("audio-player-play") as HTMLButtonElement
		this.audioPlayerPause = document.getElementById("audio-player-pause") as HTMLButtonElement
		this.trackLink = document.getElementById("audio-player-track-title") as HTMLLinkElement
		this.animeInfo = document.getElementById("audio-player-anime-info") as HTMLElement
		this.animeLink = document.getElementById("audio-player-anime-link") as HTMLLinkElement
		this.animeImage = document.getElementById("audio-player-anime-image") as HTMLImageElement
	}

	// Play audio file
	play(trackId: string, trackUrl: string) {
		if(typeof AudioContext === "undefined") {
			alert("Your browser doesn't support web audio!")
			return
		}

		if(!this.audioContext) {
			this.audioContext = new AudioContext()
			this.gainNode = this.audioContext.createGain()
			this.gainNode.gain.setTargetAtTime(this.volume, this.audioContext.currentTime + this.volumeSmoothingDelay, this.volumeTimeConstant)
		}

		this.playId++
		const currentPlayId = this.playId

		if(this.lastRequest) {
			this.lastRequest.abort()
			this.lastRequest = null
		}

		// Stop current track
		this.stop()

		this.arn.currentMediaId = trackId
		this.arn.markPlayingMedia()
		this.arn.loading(true)

		// Mark as loading
		this.audioPlayer.classList.add("loading-network")
		this.audioPlayer.classList.remove("decoding-audio")
		this.audioPlayer.classList.remove("decoded")

		// Request
		const request = new XMLHttpRequest()
		request.open("GET", trackUrl, true)
		request.responseType = "arraybuffer"

		request.onload = () => {
			if(currentPlayId !== this.playId) {
				return
			}

			// Mark as loading finished, now decoding starts
			this.audioPlayer.classList.add("decoding-audio")
			this.arn.loading(false)

			// Connect gain node to audio context when the response arrives.
			// Connecting it instantly on audio context creation would fail
			// without any error and also wouldn't work on audio source changes.
			this.gainNode.connect(this.audioContext.destination)

			// Show status message
			this.arn.statusMessage.showInfo("Decoding audio...", -1)

			this.audioContext.decodeAudioData(request.response, async buffer => {
				if(currentPlayId !== this.playId) {
					return
				}

				// Close status message
				this.arn.statusMessage.showInfo(`Playing "${this.trackLink.textContent}"`)

				// Mark as ready
				this.audioPlayer.classList.add("decoded")

				this.audioNode = this.audioContext.createBufferSource()
				this.audioNode.buffer = buffer
				this.audioNode.connect(this.gainNode)
				this.audioNode.playbackRate.setValueAtTime(this.targetSpeed, 0)
				this.audioNode.start(0)

				this.audioNode.onended = (_: MediaStreamErrorEvent) => {
					if(currentPlayId !== this.playId) {
						return
					}

					this.next()
				}
			}, err => {
				if(currentPlayId !== this.playId) {
					return
				}

				this.arn.statusMessage.showError("Your browser can't decode this audio file.")
				console.error(err)
			})
		}

		request.onprogress = e => {
			if(!e.lengthComputable) {
				return
			}

			const progress = e.loaded / e.total * 100
			this.arn.statusMessage.showInfo(`Loading audio...${progress.toFixed(1)}%`, -1)
		}

		request.onerror = () => {
			this.arn.loading(false)
		}

		// Show status message
		this.arn.statusMessage.showInfo("Loading audio...", -1)

		this.lastRequest = request
		request.send()

		// Resume audio context if it was paused before
		if(this.audioContext.state === "suspended") {
			this.audioContext.resume()
		}

		// Update track info
		this.updateTrackInfo(trackId)

		// Show audio player
		this.audioPlayer.classList.remove("fade-out")
		this.audioPlayerPlay.classList.add("fade-out")
		this.audioPlayerPause.classList.remove("fade-out")
	}

	// Pause
	pause() {
		if(!this.audioNode || !this.audioContext || this.audioContext.state === "suspended") {
			return
		}

		this.audioNode.playbackRate.setValueAtTime(0.0, 0)
		this.audioContext.suspend()

		this.audioPlayerPlay.classList.remove("fade-out")
		this.audioPlayerPause.classList.add("fade-out")
	}

	// Resume
	resume() {
		if(!this.audioNode) {
			this.next()
			return
		}

		if(!this.audioContext || this.audioContext.state === "running") {
			return
		}

		this.audioNode.playbackRate.setValueAtTime(this.targetSpeed, 0)
		this.audioContext.resume()

		this.audioPlayerPlay.classList.add("fade-out")
		this.audioPlayerPause.classList.remove("fade-out")
	}

	// Stop
	stop() {
		this.arn.currentMediaId = ""

		// Remove CSS class "playing"
		const playingElements = document.getElementsByClassName("playing")

		for(const playing of playingElements) {
			playing.classList.remove("playing")
		}

		// Fade out sidebar player
		// audioPlayer.classList.add("fade-out")

		// Remove title
		this.trackLink.href = ""
		this.trackLink.textContent = ""

		// Hide anime info
		this.animeLink.href = ""
		this.animeInfo.classList.add("hidden")
		this.animeImage.classList.add("hidden")

		// Show play button
		this.audioPlayerPlay.classList.remove("fade-out")
		this.audioPlayerPause.classList.add("fade-out")

		if(this.gainNode) {
			this.gainNode.disconnect()
		}

		if(this.audioNode) {
			this.audioNode.stop()
			this.audioNode.disconnect()
			this.audioNode = null
		}
	}

	// Previous track
	previous() {
		this.arn.statusMessage.showInfo("Previous track feature is currently work in progress! Check back later :)")
	}

	// Next track
	async next() {
		// Get random track
		const response = await fetch("/api/next/soundtrack")
		const track = await response.json()

		this.play(track.id, "https://notify.moe/audio/" + track.file)
		// arn.statusMessage.showInfo("Now playing: " + track.title)

		return track
	}

	// Set volume
	setVolume(volume: number) {
		this.volume = volume

		if(!this.gainNode) {
			return
		}

		this.gainNode.gain.setTargetAtTime(volume, this.audioContext.currentTime + this.volumeSmoothingDelay, this.volumeTimeConstant)
	}

	// Add speed
	addSpeed(speed: number) {
		if(!this.audioNode || this.audioContext.state === "suspended") {
			return
		}

		this.targetSpeed += speed

		if(this.targetSpeed < 0.5) {
			this.targetSpeed = 0.5
		} else if(this.targetSpeed > 2) {
			this.targetSpeed = 2
		}

		this.audioNode.playbackRate.setValueAtTime(this.targetSpeed, 0)
		this.arn.statusMessage.showInfo("Playback speed: " + Math.round(this.targetSpeed * 100) + "%")
	}

	// Play or pause
	playPause() {
		if(!this.audioNode) {
			this.next()
			return
		}

		if(this.audioContext.state === "suspended") {
			this.resume()
		} else {
			this.pause()
		}
	}

	// Update track info
	async updateTrackInfo(trackId: string) {
		// Set track title
		const trackInfoResponse = await fetch("/api/soundtrack/" + trackId)
		const track = await trackInfoResponse.json()
		this.trackLink.href = "/soundtrack/" + track.id
		this.trackLink.textContent = track.title.canonical || track.title.native

		// Find anime ID by looking at the tags
		let animeId = ""

		if(!track.tags) {
			return
		}

		for(const tag of (track.tags as string[])) {
			if(tag.startsWith("anime:")) {
				animeId = tag.split(":")[1]
				break
			}
		}

		if(animeId === "") {
			return
		}

		// Set anime info
		this.animeInfo.classList.remove("hidden")
		const animeResponse = await fetch("/api/anime/" + animeId)
		const anime = await animeResponse.json() as Anime
		this.animeLink.title = anime.title.canonical
		this.animeLink.href = "/anime/" + anime.id
		this.animeImage.dataset.src = "//media.notify.moe/images/anime/medium/" + anime.id + ".jpg?" + anime.image.lastModified.toString()
		this.animeImage.classList.remove("hidden")
		this.animeImage["became visible"]()
	}
}
