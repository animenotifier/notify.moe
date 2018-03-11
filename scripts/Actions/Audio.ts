import { AnimeNotifier } from "../AnimeNotifier"

// Play audio file
export function playAudio(arn: AnimeNotifier, button: HTMLButtonElement) {
	if(!arn.audio) {
		arn.audio = document.createElement("audio") as HTMLAudioElement
		let source = document.createElement("source") as HTMLSourceElement
		source.src = button.dataset.src
		arn.audio.appendChild(source)
	}

	arn.audio.play()
}