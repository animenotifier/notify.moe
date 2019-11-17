import AnimeNotifier from "../AnimeNotifier"

// Play audio
export function playAudio(arn: AnimeNotifier, element: HTMLElement) {
	const {mediaId, audioSrc} = element.dataset

	if(!mediaId || !audioSrc) {
		console.error("Invalid media ID or audio source:", element)
		return
	}

	arn.audioPlayer.play(mediaId, audioSrc)
}

// Pause audio
export function pauseAudio(arn: AnimeNotifier) {
	arn.audioPlayer.pause()
}

// Resume audio
export function resumeAudio(arn: AnimeNotifier) {
	arn.audioPlayer.resume()
}

// Stop audio
export function stopAudio(arn: AnimeNotifier) {
	arn.audioPlayer.stop()
}

// Play previous track
export async function playPreviousTrack(arn: AnimeNotifier) {
	arn.audioPlayer.previous()
}

// Play next track
export async function playNextTrack(arn: AnimeNotifier) {
	arn.audioPlayer.next()
}

// Set volume
export function setVolume(arn: AnimeNotifier, element: HTMLInputElement) {
	const volume = parseFloat(element.value) / 100.0
	arn.audioPlayer.setVolume(volume)
}

// Play or pause audio
export function playPauseAudio(arn: AnimeNotifier) {
	arn.audioPlayer.playPause()
}

// Toggle audio
export function toggleAudio(arn: AnimeNotifier, element: HTMLElement) {
	// If we're clicking on the same track again, stop playing.
	// Otherwise, start the track we clicked on.
	if(arn.currentMediaId && element.dataset.mediaId === arn.currentMediaId) {
		stopAudio(arn)
	} else {
		playAudio(arn, element)
	}
}
