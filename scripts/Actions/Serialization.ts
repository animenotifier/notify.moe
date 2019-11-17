import AnimeNotifier from "../AnimeNotifier"
import { applyTheme } from "./Theme"

// Save new data from an input field
export async function save(arn: AnimeNotifier, input: HTMLElement) {
	if(!input.dataset.field) {
		console.error("Input element missing data-field:", input)
		return
	}

	const obj = {}
	const isContentEditable = input.isContentEditable
	let value = isContentEditable ? input.textContent : (input as HTMLInputElement).value

	if(value === undefined || value === null) {
		return
	}

	// Trim value
	value = value.trim()

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

	const apiEndpoint = arn.findAPIEndpoint(input)

	try {
		await arn.post(apiEndpoint, obj)

		if(apiEndpoint.startsWith("/api/user/") && input.dataset.field === "Nick") {
			// Update nickname based links on the page
			return arn.reloadPage()
		} else if(apiEndpoint.startsWith("/api/settings/") && input.dataset.field === "Theme") {
			// Apply theme instantly
			applyTheme((input as HTMLInputElement).value)

			// Reload to update theme settings
			return arn.reloadContent()
		} else {
			return arn.reloadContent()
		}
	} catch(err) {
		arn.reloadContent()
		arn.statusMessage.showError(err)
	} finally {
		if(isContentEditable) {
			input.contentEditable = "true"
		} else {
			(input as HTMLInputElement).disabled = false
		}
	}
}

// Enable (bool field)
export async function enable(arn: AnimeNotifier, button: HTMLButtonElement) {
	if(!button.dataset.field) {
		console.error("Button missing data-field:", button)
		return
	}

	const obj = {}
	const apiEndpoint = arn.findAPIEndpoint(button)

	obj[button.dataset.field] = true
	button.disabled = true

	try {
		// Update boolean value
		await arn.post(apiEndpoint, obj)

		// Reload content
		arn.reloadContent()

		arn.statusMessage.showInfo("Enabled: " + button.title)
	} catch(err) {
		arn.statusMessage.showError(err)
	} finally {
		button.disabled = false
	}
}

// Disable (bool field)
export async function disable(arn: AnimeNotifier, button: HTMLButtonElement) {
	if(!button.dataset.field) {
		console.error("Button missing data-field:", button)
		return
	}

	const obj = {}
	const apiEndpoint = arn.findAPIEndpoint(button)

	obj[button.dataset.field] = false
	button.disabled = true

	try {
		// Update boolean value
		await arn.post(apiEndpoint, obj)

		// Reload content
		arn.reloadContent()

		arn.statusMessage.showInfo("Disabled: " + button.title)
	} catch(err) {
		arn.statusMessage.showError(err)
	} finally {
		button.disabled = false
	}
}

// Append new element to array
export async function arrayAppend(arn: AnimeNotifier, element: HTMLElement) {
	const field = element.dataset.field
	const object = element.dataset.object || ""
	const apiEndpoint = arn.findAPIEndpoint(element)

	try {
		await arn.post(apiEndpoint + "/field/" + field + "/append", object)
		await arn.reloadContent()
	} catch(err) {
		arn.statusMessage.showError(err)
	}
}

// Remove element from array
export async function arrayRemove(arn: AnimeNotifier, element: HTMLElement) {
	if(!confirm("Are you sure you want to remove this element?")) {
		return
	}

	const field = element.dataset.field
	const index = element.dataset.index
	const apiEndpoint = arn.findAPIEndpoint(element)

	try {
		await arn.post(apiEndpoint + "/field/" + field + "/remove/" + index)
		await arn.reloadContent()
	} catch(err) {
		arn.statusMessage.showError(err)
	}
}

// Increase episode
export function increaseEpisode(arn: AnimeNotifier, element: HTMLElement) {
	if(arn.isLoading) {
		return
	}

	const prev = element.previousSibling

	if(prev === null || !(prev instanceof HTMLElement) || prev.textContent === null) {
		console.error("Previous sibling is invalid:", element)
		return
	}

	const episodes = parseInt(prev.textContent)
	prev.textContent = String(episodes + 1)
	return save(arn, prev)
}

// Add number
export function addNumber(arn: AnimeNotifier, element: HTMLElement) {
	if(arn.isLoading) {
		return
	}

	if(!element.dataset.id || !element.dataset.add) {
		console.error("Element is missing the data-id or data-add attribute:", element)
		return
	}

	const input = document.getElementById(element.dataset.id) as HTMLInputElement
	const add = parseInt(element.dataset.add)
	const num = parseInt(input.value)
	const newValue = num + add
	const min = parseInt(input.min)
	const max = parseInt(input.max)

	if(newValue > max) {
		arn.statusMessage.showError("Maximum: " + max)
		return
	}

	if(newValue < min) {
		arn.statusMessage.showError("Minimum: " + min)
		return
	}

	input.value = newValue.toString()
	return save(arn, input)
}
