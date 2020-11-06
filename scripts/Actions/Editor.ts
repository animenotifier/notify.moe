import AnimeNotifier from "../AnimeNotifier"

// newAnimeDiffIgnore
export async function newAnimeDiffIgnore(arn: AnimeNotifier, button: HTMLButtonElement) {
	if(!confirm("Are you sure you want to permanently ignore this difference?")) {
		return
	}

	const id = button.dataset.id
	const hash = button.dataset.hash

	try {
		await arn.post(`/api/new/ignoreanimedifference`, {
			id,
			hash
		})

		arn.reloadContent()
	} catch(err) {
		arn.statusMessage.showError(err)
	}
}

// Import Kitsu anime
export async function importKitsuAnime(arn: AnimeNotifier, button: HTMLButtonElement) {
	if(!confirm("Are you sure you want to import this anime?")) {
		return
	}

	const newTab = window.open()

	if(!newTab) {
		arn.statusMessage.showError("Error opening new tab")
		return
	}

	const animeId = button.dataset.id
	const response = await fetch(`/api/import/kitsu/anime/${animeId}`, {
		method: "POST",
		credentials: "same-origin"
	})

	if(response.ok) {
		newTab.location.href = `/kitsu/anime/${animeId}`
		arn.reloadContent()
	} else {
		arn.statusMessage.showError(await response.text())
	}
}

// Delete Kitsu anime
export async function deleteKitsuAnime(arn: AnimeNotifier, button: HTMLButtonElement) {
	if(!confirm("Are you sure you want to delete this anime?")) {
		return
	}

	const animeId = button.dataset.id
	await arn.post(`/api/delete/kitsu/anime/${animeId}`)
	arn.reloadContent()
}

// Multi-search anime
export async function multiSearchAnime(arn: AnimeNotifier, textarea: HTMLTextAreaElement) {
	const results = document.getElementById("multi-search-anime") as HTMLDivElement
	const animeTitles = textarea.value.split("\n")

	results.innerHTML = ""

	for(let i = 0; i < animeTitles.length; i++) {
		console.log(animeTitles[i])
		const response = await fetch("/_/anime-search/" + animeTitles[i])
		const html = await response.text()
		results.innerHTML += "<h3>" + animeTitles[i] + "</h3>" + html
	}

	results.classList.remove("hidden")
	arn.onNewContent(results)
}

// Download soundtrack file
export async function downloadSoundTrackFile(arn: AnimeNotifier, button: HTMLButtonElement) {
	const id = button.dataset.id

	try {
		await arn.post(`/api/soundtrack/${id}/download`)
		arn.reloadContent()
	} catch(err) {
		arn.statusMessage.showError(err)
	}
}

// Start background job
export async function startJob(arn: AnimeNotifier, button: HTMLButtonElement) {
	if(button.dataset.running === "true") {
		alert("Job is currently running!")
		return
	}

	const jobName = button.dataset.job

	if(!confirm(`Are you sure you want to start the "${jobName}" job?`)) {
		return
	}

	await arn.post(`/api/job/${jobName}/start`)
	arn.reloadContent()
}

// Sync episodes
export async function syncEpisodes(arn: AnimeNotifier, button: HTMLButtonElement) {
	if(!confirm("Are you sure you want to start the episode sync?")) {
		return
	}

	const animeId = button.dataset.animeId

	arn.statusMessage.showInfo("Started episode sync.", -1)

	try {
		await arn.post(`/api/anime/${animeId}/sync-episodes`)
		arn.statusMessage.showInfo("Finished episode sync.")
	} catch(err) {
		arn.statusMessage.showError(err)
	}
}
