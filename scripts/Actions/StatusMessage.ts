import AnimeNotifier from "../AnimeNotifier"

// Close status message
export function closeStatusMessage(arn: AnimeNotifier) {
	arn.statusMessage.close()
}
