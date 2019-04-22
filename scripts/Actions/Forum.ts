import AnimeNotifier from "../AnimeNotifier"

// Edit post
export function editPost(arn: AnimeNotifier, element: HTMLElement) {
	let postId = element.dataset.id

	if(!postId) {
		console.error("Post missing post ID:", postId)
		return
	}

	let render = document.getElementById("render-" + postId) as HTMLElement
	let toolbar = document.getElementById("toolbar-" + postId) as HTMLElement
	let source = document.getElementById("source-" + postId) as HTMLElement
	let edit = document.getElementById("edit-toolbar-" + postId) as HTMLElement

	render.classList.toggle("hidden")
	toolbar.classList.toggle("hidden")
	source.classList.toggle("hidden")
	edit.classList.toggle("hidden")

	let title = document.getElementById("title-" + postId)

	if(title) {
		title.classList.toggle("hidden")
	}
}

// Save post
export async function savePost(arn: AnimeNotifier, element: HTMLElement) {
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

	try {
		await arn.post(apiEndpoint, updates)
		arn.reloadContent()
	} catch(err) {
		arn.statusMessage.showError(err)
	}
}

// Delete post
export async function deletePost(arn: AnimeNotifier, element: HTMLElement) {
	if(!confirm(`Are you sure you want to delete this Post?`)) {
		return
	}

	let endpoint = arn.findAPIEndpoint(element)

	try {
		await arn.post(endpoint + "/delete")
		arn.reloadContent()
	} catch(err) {
		arn.statusMessage.showError(err)
	}
}

// Create post
export async function createPost(arn: AnimeNotifier, element: HTMLElement) {
	let textarea = document.getElementById("new-post-text") as HTMLTextAreaElement
	let {parentId, parentType} = element.dataset

	let post = {
		text: textarea.value,
		parentId,
		parentType,
		tags: []
	}

	try {
		await arn.post("/api/new/post", post)
		await arn.reloadContent()
		textarea.value = ""
	} catch(err) {
		arn.statusMessage.showError(err)
	}
}

// Create thread
export async function createThread(arn: AnimeNotifier) {
	let title = document.getElementById("title") as HTMLInputElement
	let text = document.getElementById("text") as HTMLTextAreaElement
	let category = document.getElementById("tag") as HTMLInputElement

	let thread = {
		title: title.value,
		text: text.value,
		tags: [category.value]
	}

	try {
		await arn.post("/api/new/thread", thread)
		await arn.app.load("/forum/" + thread.tags[0])
	} catch(err) {
		arn.statusMessage.showError(err)
	}
}

// Reply to a post
export async function reply(arn: AnimeNotifier, element: HTMLElement) {
	let apiEndpoint = arn.findAPIEndpoint(element)
	let repliesId = `replies-${element.dataset.postId}`
	let replies = document.getElementById(repliesId)

	if(!replies) {
		console.error("Missing replies container:", element)
		return
	}

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
	return setThreadLock(arn, element, true)
}

// Unlock thread
export function unlockThread(arn: AnimeNotifier, element: HTMLButtonElement) {
	return setThreadLock(arn, element, false)
}

// Set thread locked state
async function setThreadLock(arn: AnimeNotifier, element: HTMLButtonElement, state: boolean) {
	let verb = state ? "lock" : "unlock"

	if(!confirm(`Are you sure you want to ${verb} this Thread?`)) {
		return
	}

	let endpoint = arn.findAPIEndpoint(element)

	try {
		await arn.post(`${endpoint}/${verb}`)
		await arn.reloadContent()
	} catch(err) {
		arn.statusMessage.showError(err)
	}
}