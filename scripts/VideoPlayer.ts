import AnimeNotifier from "./AnimeNotifier"

export default class VideoPlayer {
	arn: AnimeNotifier

	constructor(arn: AnimeNotifier) {
		this.arn = arn
	}

	play(_: HTMLVideoElement) {

	}

	playPause() {

	}
}
