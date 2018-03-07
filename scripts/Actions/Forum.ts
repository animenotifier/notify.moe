import { AnimeNotifier } from "../AnimeNotifier"

// Edit post
export function editPost(arn: AnimeNotifier, element: HTMLElement) {
	let postId = element.dataset.id

	let render = arn.app.find("render-" + postId)
	let toolbar = arn.app.find("toolbar-" + postId)
	let title = arn.app.find("title-" + postId)
	let source = arn.app.find("source-" + postId)
	let edit = arn.app.find("edit-toolbar-" + postId)

	render.classList.toggle("hidden")
	toolbar.classList.toggle("hidden")
	source.classList.toggle("hidden")
	edit.classList.toggle("hidden")

	if(title) {
		title.classList.toggle("hidden")
	}
}

// Save post
export function savePost(arn: AnimeNotifier, element: HTMLElement) {
	let postId = element.dataset.id
	let source = arn.app.find("source-" + postId) as HTMLTextAreaElement
	let title = arn.app.find("title-" + postId) as HTMLInputElement
	let text = source.value

	let updates: any = {
		Text: text,
	}

	// Add title for threads only
	if(title) {
		updates.Title = title.value
	}

	let apiEndpoint = arn.findAPIEndpoint(element)

	arn.post(apiEndpoint, updates)
	.then(() => arn.reloadContent())
	.catch(err => arn.statusMessage.showError(err))
}

// Delete post
export function deletePost(arn: AnimeNotifier, element: HTMLElement) {
	if(!confirm(`Are you sure you want to delete this Post?`)) {
		return
	}

	let endpoint = arn.findAPIEndpoint(element)

	arn.post(endpoint + "/delete", "")
	.then(() => arn.reloadContent())
	.catch(err => arn.statusMessage.showError(err))
}

// Forum reply
export function forumReply(arn: AnimeNotifier) {
	let textarea = arn.app.find("new-reply") as HTMLTextAreaElement
	let thread = arn.app.find("thread")

	let post = {
		text: textarea.value,
		threadId: thread.dataset.id,
		tags: []
	}

	arn.post("/api/new/post", post)
	.then(() => arn.reloadContent())
	.then(() => textarea.value = "")
	.catch(err => arn.statusMessage.showError(err))
}

// Group post
export function newGroupPost(arn: AnimeNotifier) {
	// TODO: ...
}

// Create thread
export function createThread(arn: AnimeNotifier) {
	let title = arn.app.find("title") as HTMLInputElement
	let text = arn.app.find("text") as HTMLTextAreaElement
	let category = arn.app.find("tag") as HTMLInputElement

	let thread = {
		title: title.value,
		text: text.value,
		tags: [category.value]
	}

	arn.post("/api/new/thread", thread)
	.then(() => arn.app.load("/forum/" + thread.tags[0]))
	.catch(err => arn.statusMessage.showError(err))
}