import AnimeNotifier from "../AnimeNotifier"

// Edit post
export function editPost(_: AnimeNotifier, element: HTMLElement) {
	const postId = element.dataset.id

	if(!postId) {
		console.error("Post missing post ID:", postId)
		return
	}

	const render = document.getElementById("render-" + postId) as HTMLElement
	const source = document.getElementById("source-" + postId) as HTMLElement
	const edit = document.getElementById("edit-toolbar-" + postId) as HTMLElement

	render.classList.toggle("hidden")
	source.classList.toggle("hidden")
	edit.classList.toggle("hidden")

	const title = document.getElementById("title-" + postId)

	if(title) {
		title.classList.toggle("hidden")
	}
}

// Save post
export async function savePost(arn: AnimeNotifier, element: HTMLElement) {
	const postId = element.dataset.id
	const source = document.getElementById("source-" + postId) as HTMLTextAreaElement
	const title = document.getElementById("title-" + postId) as HTMLInputElement
	const text = source.value

	const updates: any = {
		Text: text,
	}

	// Add title for threads only
	if(title) {
		updates.Title = title.value
	}

	const apiEndpoint = arn.findAPIEndpoint(element)

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

	const endpoint = arn.findAPIEndpoint(element)

	try {
		await arn.post(endpoint + "/delete")
		arn.reloadContent()
	} catch(err) {
		arn.statusMessage.showError(err)
	}
}

// Create post
export async function createPost(arn: AnimeNotifier, element: HTMLElement) {
	const textarea = document.getElementById("new-post-text") as HTMLTextAreaElement
	const {parentId, parentType} = element.dataset

	const post = {
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
	const title = document.getElementById("title") as HTMLInputElement
	const text = document.getElementById("text") as HTMLTextAreaElement
	const category = document.getElementById("tag") as HTMLInputElement

	const thread = {
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
	const apiEndpoint = arn.findAPIEndpoint(element)
	const repliesId = `replies-${element.dataset.postId}`
	const replies = document.getElementById(repliesId)

	if(!replies) {
		console.error("Missing replies container:", element)
		return
	}

	// Delete old reply area
	const oldReplyArea = document.getElementById("new-post")

	if(oldReplyArea) {
		oldReplyArea.remove()
	}

	// Delete old reply button
	const oldPostActions = document.getElementsByClassName("new-post-actions")[0]

	if(oldPostActions) {
		oldPostActions.remove()
	}

	// Fetch new reply UI
	try {
		const response = await fetch(`${apiEndpoint}/reply/ui`)
		const html = await response.text()
		replies.innerHTML = html + replies.innerHTML
		arn.onNewContent(replies)
		arn.assignActions()
	} catch(err) {
		arn.statusMessage.showError(err)
	}
}

// Cancel replying to a post
export function cancelReply(arn: AnimeNotifier, _: HTMLElement) {
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
	const verb = state ? "lock" : "unlock"

	if(!confirm(`Are you sure you want to ${verb} this Thread?`)) {
		return
	}

	const endpoint = arn.findAPIEndpoint(element)

	try {
		await arn.post(`${endpoint}/${verb}`)
		await arn.reloadContent()
	} catch(err) {
		arn.statusMessage.showError(err)
	}
}
