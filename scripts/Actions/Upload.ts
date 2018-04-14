import AnimeNotifier from "../AnimeNotifier"
import StatusMessage from "../StatusMessage"

// Select file
export function selectFile(arn: AnimeNotifier, button: HTMLButtonElement) {
	if(button.dataset.endpoint === "/api/upload/cover" && arn.user.dataset.pro !== "true") {
		alert("Please buy a PRO account to use this feature.")
		return
	}

	let fileType = button.dataset.type
	let endpoint = button.dataset.endpoint

	// Click on virtual file input element
	let input = document.createElement("input")
	input.setAttribute("type", "file")

	input.onchange = () => {
		let file = input.files[0]

		if(!file) {
			return
		}

		// Check mime type for images
		if(fileType === "image" && !file.type.startsWith("image/")) {
			arn.statusMessage.showError(file.name + " is not an image file!")
			return
		}

		// Check mime type for videos
		if(fileType === "video" && !file.type.startsWith("video/webm")) {
			arn.statusMessage.showError(file.name + " is not a WebM video file!")
			return
		}

		// Preview image
		if(fileType === "image") {
			let preview = document.getElementById(button.id + "-preview") as HTMLImageElement

			if(preview) {
				previewImage(file, endpoint, preview)
			}
		}

		uploadFile(file, fileType, endpoint, arn)
	}

	input.click()
}

// Upload file
function uploadFile(file: File, fileType: string, endpoint: string, arn: AnimeNotifier) {
	let reader = new FileReader()

	reader.onloadend = async () => {
		arn.statusMessage.showInfo(`Uploading ${fileType}...`, 60000)

		let response = await fetch(endpoint, {
			method: "POST",
			credentials: "include",
			headers: {
				"Content-Type": "application/octet-stream"
			},
			body: reader.result
		})

		if(endpoint === "/api/upload/avatar") {
			let newURL = await response.text()
			updateSideBarAvatar(newURL)
		}

		if(response.ok) {
			arn.statusMessage.showInfo(`Successfully uploaded your new ${fileType}.`)
		} else {
			arn.statusMessage.showError(`Failed uploading your new ${fileType}.`)
		}
	}

	reader.readAsArrayBuffer(file)
}

// Preview image
function previewImage(file: File, endpoint: string, preview: HTMLImageElement) {
	let reader = new FileReader()

	reader.onloadend = () => {
		if(endpoint === "/api/upload/avatar") {
			let svgPreview = document.getElementById("avatar-input-preview-svg") as HTMLImageElement

			if(svgPreview) {
				svgPreview.classList.add("hidden")
			}
		}

		preview.classList.remove("hidden")
		preview.src = reader.result
	}

	reader.readAsDataURL(file)
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