import { AnimeNotifier } from "../AnimeNotifier"

// Select file
export function selectFile(arn: AnimeNotifier, button: HTMLButtonElement) {
	let input = document.createElement("input")
	let preview = document.getElementById(button.dataset.previewImageId) as HTMLImageElement
	input.setAttribute("type", "file")

	input.onchange = () => {
		previewImage(input, preview)
		uploadImage(input, preview)
	}

	input.click()
}

// Preview image
function previewImage(input: HTMLInputElement, preview: HTMLImageElement) {
	let file = input.files[0]
	let reader = new FileReader()

	console.log(file.name, file.size, file.type)

	reader.onloadend = () => {
		preview.classList.remove("hidden")
		preview.src = reader.result
	}

	if(file) {
		reader.readAsDataURL(file)
	} else {
		preview.src = ""
	}
}

// Upload image
function uploadImage(input: HTMLInputElement, preview: HTMLImageElement) {

}