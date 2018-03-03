import { AnimeNotifier } from "../AnimeNotifier"

// Save new data from an input field
export function save(arn: AnimeNotifier, input: HTMLElement) {
	let obj = {}
	let isContentEditable = input.isContentEditable
	let value = isContentEditable ? input.innerText : (input as HTMLInputElement).value

	if(value === undefined) {
		return
	}

	if((input as HTMLInputElement).type === "number" || input.dataset.type === "number") {
		if(input.getAttribute("step") === "1" || input.dataset.step === "1") {
			obj[input.dataset.field] = parseInt(value)
		} else {
			obj[input.dataset.field] = parseFloat(value)
		}
	} else {
		obj[input.dataset.field] = value
	}

	if(isContentEditable) {
		input.contentEditable = "false"
	} else {
		(input as HTMLInputElement).disabled = true
	}

	let apiEndpoint = arn.findAPIEndpoint(input)

	arn.post(apiEndpoint, obj)
	.catch(err => arn.statusMessage.showError(err))
	.then(() => {
		if(isContentEditable) {
			input.contentEditable = "true"
		} else {
			(input as HTMLInputElement).disabled = false
		}

		if(apiEndpoint.startsWith("/api/user/") && input.dataset.field === "Nick") {
			return arn.reloadPage()
		} else {
			return arn.reloadContent()
		}
	})
}

// Append new element to array
export function arrayAppend(arn: AnimeNotifier, element: HTMLElement) {
	let field = element.dataset.field
	let object = element.dataset.object || ""
	let apiEndpoint = arn.findAPIEndpoint(element)

	arn.post(apiEndpoint + "/field/" + field + "/append", object)
	.then(() => arn.reloadContent())
	.catch(err => arn.statusMessage.showError(err))
}

// Remove element from array
export function arrayRemove(arn: AnimeNotifier, element: HTMLElement) {
	if(!confirm("Are you sure you want to remove this element?")) {
		return
	}

	let field = element.dataset.field
	let index = element.dataset.index
	let apiEndpoint = arn.findAPIEndpoint(element)

	arn.post(apiEndpoint + "/field/" + field + "/remove/" + index, "")
	.then(() => arn.reloadContent())
	.catch(err => arn.statusMessage.showError(err))
}

// Increase episode
export function increaseEpisode(arn: AnimeNotifier, element: HTMLElement) {
	if(arn.isLoading) {
		return
	}

	let prev = element.previousSibling as HTMLElement
	let episodes = parseInt(prev.innerText)
	prev.innerText = String(episodes + 1)
	save(arn, prev)
}

// Add number
export function addNumber(arn: AnimeNotifier, element: HTMLElement) {
	if(arn.isLoading) {
		return
	}

	let input = arn.app.find(element.dataset.id) as HTMLInputElement
	let add = parseInt(element.dataset.add)
	let num = parseInt(input.value)
	let newValue = num + add
	let min = parseInt(input.min)
	let max = parseInt(input.max)

	if(newValue > max) {
		arn.statusMessage.showError("Maximum: " + max)
		return
	}

	if(newValue < min) {
		arn.statusMessage.showError("Minimum: " + min)
		return
	}

	input.value = newValue.toString()
	save(arn, input)
}