import { AnimeNotifier } from "../AnimeNotifier"

// Search
export function search(arn: AnimeNotifier, search: HTMLInputElement, e: KeyboardEvent) {
	if(e.ctrlKey || e.altKey) {
		return
	}

	let term = search.value

	if(!term || term.length < 1) {
		arn.app.content.innerHTML = "Please enter at least 1 character to start searching."
		return
	}

	arn.diff("/search/" + term)
}

// Search database
export function searchDB(arn: AnimeNotifier, input: HTMLInputElement, e: KeyboardEvent) {
	if(e.ctrlKey || e.altKey) {
		return
	}

	let dataType = (arn.app.find("data-type") as HTMLInputElement).value || "+"
	let field = (arn.app.find("field") as HTMLInputElement).value || "+"
	let fieldValue = (arn.app.find("field-value") as HTMLInputElement).value || "+"
	let records = arn.app.find("records")

	arn.loading(true)

	fetch(`/api/select/${dataType}/where/${field}/is/${fieldValue}`)
	.then(response => {
		if(response.status !== 200) {
			throw response
		}

		return response
	})
	.then(response => response.json())
	.then(data => {
		records.innerHTML = ""
		let count = 0

		if(data.results.length === 0) {
			records.innerText = "No results."
			return
		}

		for(let record of data.results) {
			count++

			let container = document.createElement("div")
			container.classList.add("record")

			let id = document.createElement("div")
			id.innerText = record.id
			id.classList.add("record-id")
			container.appendChild(id)

			let link = document.createElement("a")
			link.classList.add("record-view")
			link.innerText = "Open " + dataType.toLowerCase()

			if(dataType === "User") {
				link.href = "/+" + record.nick
			} else {
				link.href = "/" + dataType.toLowerCase() + "/" + record.id
			}
			
			link.target = "_blank"
			container.appendChild(link)

			let apiLink = document.createElement("a")
			apiLink.classList.add("record-view-api")
			apiLink.innerText = "JSON data"
			apiLink.href = "/api/" + dataType.toLowerCase() + "/" + record.id
			apiLink.target = "_blank"
			container.appendChild(apiLink)

			let recordCount = document.createElement("div")
			recordCount.innerText = count + "/" + data.results.length
			recordCount.classList.add("record-count")
			container.appendChild(recordCount)

			records.appendChild(container)
		}
	})
	.catch(response => {
		response.text().then(text => {
			arn.statusMessage.showError(text)
		})
	})
	.then(() => arn.loading(false))
}