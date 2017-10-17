import { AnimeNotifier } from "../AnimeNotifier"

// New soundtrack
export function newSoundTrack(arn: AnimeNotifier, button: HTMLButtonElement) {
	arn.post("/api/new/soundtrack", "")
	.then(response => response.json())
	.then(track => arn.app.load(`/soundtrack/${track.id}/edit`))
	.catch(err => arn.statusMessage.showError(err))
}