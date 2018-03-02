import { AnimeNotifier } from "../AnimeNotifier"

// Select file
export function selectFile(arn: AnimeNotifier, button: HTMLButtonElement) {
	let input = document.createElement("input")
	let preview = document.getElementById(button.dataset.previewImageId) as HTMLImageElement
	input.setAttribute("type", "file")

	input.onchange = () => {
		let file = input.files[0]

		previewImage(file, preview)
		uploadFile(file, "/api/upload/avatar")
	}

	input.click()
}

// Preview image
function previewImage(file: File, preview: HTMLImageElement) {
	let reader = new FileReader()

	reader.onloadend = () => {
		preview.classList.remove("hidden")
		preview.src = reader.result
	}

	if(file) {
		reader.readAsDataURL(file)
	} else {
		preview.classList.add("hidden")
	}
}

// Upload file
function uploadFile(file: File, endpoint: string) {
	let reader = new FileReader()

	reader.onloadend = async () => {
		await fetch(endpoint, {
			method: "POST",
			credentials: "include",
			headers: {
				"Content-Type": "application/octet-stream"
			},
			body: reader.result
		})
	}

	reader.readAsArrayBuffer(file)
}