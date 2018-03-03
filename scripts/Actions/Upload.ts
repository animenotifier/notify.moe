import { AnimeNotifier } from "../AnimeNotifier"
import { StatusMessage } from "../StatusMessage"

// Select file
export function selectFile(arn: AnimeNotifier, button: HTMLButtonElement) {
	let input = document.createElement("input")
	let preview = document.getElementById(button.dataset.previewImageId) as HTMLImageElement
	input.setAttribute("type", "file")

	input.onchange = () => {
		let file = input.files[0]

		if(!file) {
			return
		}

		if(!file.type.startsWith("image/")) {
			arn.statusMessage.showError(file.name + " is not an image file!")
			return
		}

		previewImage(file, preview)
		uploadFile(file, "/api/upload/avatar", arn)
	}

	input.click()
}

// Preview image
function previewImage(file: File, preview: HTMLImageElement) {
	let reader = new FileReader()

	reader.onloadend = () => {
		let svgPreview = document.getElementById("avatar-input-preview-svg") as HTMLImageElement

		if(svgPreview) {
			svgPreview.classList.add("hidden")
		}

		preview.classList.remove("hidden")
		preview.src = reader.result
	}

	reader.readAsDataURL(file)
}

// Upload file
function uploadFile(file: File, endpoint: string, arn: AnimeNotifier) {
	let reader = new FileReader()

	reader.onloadend = async () => {
		arn.statusMessage.showInfo("Uploading avatar...", 60000)

		let response = await fetch(endpoint, {
			method: "POST",
			credentials: "include",
			headers: {
				"Content-Type": "application/octet-stream"
			},
			body: reader.result
		})

		let newURL = await response.text()
		updateSideBarAvatar(newURL)

		if(response.ok) {
			arn.statusMessage.showInfo("Successfully uploaded your new avatar.")
		} else {
			arn.statusMessage.showError("Failed uploading your new avatar.")
		}
	}

	reader.readAsArrayBuffer(file)
}

// Update sidebar avatar
function updateSideBarAvatar(url: string) {
	let sidebar = document.getElementById("sidebar")
	let userImage = sidebar.getElementsByClassName("user-image")[0] as HTMLImageElement
	let lazyLoad = userImage["became visible"]

	if(lazyLoad) {
		userImage.dataset.src = url
		lazyLoad()
	} else {
		location.reload()
	}
}