import { AnimeNotifier } from "../AnimeNotifier"

// Play audio
export function playAudio(arn: AnimeNotifier, element: HTMLElement) {
	arn.audioPlayer.play(element.dataset.soundtrackId, element.dataset.audioSrc)
}

// Pause audio
export function pauseAudio(arn: AnimeNotifier, button: HTMLButtonElement) {
	arn.audioPlayer.pause()
}

// Resume audio
export function resumeAudio(arn: AnimeNotifier, button: HTMLButtonElement) {
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
	let volume = parseFloat(element.value) / 100.0
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
	if(arn.currentSoundTrackId && element.dataset.soundtrackId === arn.currentSoundTrackId) {
		stopAudio(arn)
	} else {
		playAudio(arn, element)
	}
}
