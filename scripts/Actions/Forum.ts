import AnimeNotifier from "../AnimeNotifier"

// Edit post
export function editPost(arn: AnimeNotifier, element: HTMLElement) {
	let postId = element.dataset.id

	let render = document.getElementById("render-" + postId)
	let toolbar = document.getElementById("toolbar-" + postId)
	let title = document.getElementById("title-" + postId)
	let source = document.getElementById("source-" + postId)
	let edit = document.getElementById("edit-toolbar-" + postId)

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
	let source = document.getElementById("source-" + postId) as HTMLTextAreaElement
	let title = document.getElementById("title-" + postId) as HTMLInputElement
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

	arn.post(endpoint + "/delete")
	.then(() => arn.reloadContent())
	.catch(err => arn.statusMessage.showError(err))
}

// Create post
export function createPost(arn: AnimeNotifier, element: HTMLElement) {
	let textarea = document.getElementById("new-post-text") as HTMLTextAreaElement
	let {parentId, parentType} = element.dataset

	let post = {
		text: textarea.value,
		parentId,
		parentType,
		tags: []
	}

	arn.post("/api/new/post", post)
	.then(() => arn.reloadContent())
	.then(() => textarea.value = "")
	.catch(err => arn.statusMessage.showError(err))
}

// Create thread
export function createThread(arn: AnimeNotifier) {
	let title = document.getElementById("title") as HTMLInputElement
	let text = document.getElementById("text") as HTMLTextAreaElement
	let category = document.getElementById("tag") as HTMLInputElement

	let thread = {
		title: title.value,
		text: text.value,
		tags: [category.value]
	}

	arn.post("/api/new/thread", thread)
	.then(() => arn.app.load("/forum/" + thread.tags[0]))
	.catch(err => arn.statusMessage.showError(err))
}

// Reply to a post
export async function reply(arn: AnimeNotifier, element: HTMLElement) {
	let apiEndpoint = arn.findAPIEndpoint(element)
	let repliesId = `replies-${element.dataset.postId}`
	let replies = document.getElementById(repliesId)

	// Delete old reply area
	let oldReplyArea = document.getElementById("new-post")

	if(oldReplyArea) {
		oldReplyArea.remove()
	}

	// Delete old reply button
	let oldPostActions = document.getElementsByClassName("new-post-actions")[0]

	if(oldPostActions) {
		oldPostActions.remove()
	}

	// Fetch new reply UI
	try {
		let response = await fetch(`${apiEndpoint}/reply/ui`)
		let html = await response.text()
		replies.innerHTML = html + replies.innerHTML
		arn.onNewContent(replies)
		arn.assignActions()
	} catch(err) {
		arn.statusMessage.showError(err)
	}
}

// Cancel replying to a post
export function cancelReply(arn: AnimeNotifier, element: HTMLElement) {
	arn.reloadContent()
}

// Lock thread
export function lockThread(arn: AnimeNotifier, element: HTMLButtonElement) {
	setThreadLock(arn, element, true)
}

// Unlock thread
export function unlockThread(arn: AnimeNotifier, element: HTMLButtonElement) {
	setThreadLock(arn, element, false)
}

// Set thread locked state
function setThreadLock(arn: AnimeNotifier, element: HTMLButtonElement, state: boolean) {
	let verb = state ? "lock" : "unlock"

	if(!confirm(`Are you sure you want to ${verb} this Thread?`)) {
		return
	}

	let endpoint = arn.findAPIEndpoint(element)

	arn.post(`${endpoint}/${verb}`)
	.then(() => arn.reloadContent())
	.catch(err => arn.statusMessage.showError(err))
}