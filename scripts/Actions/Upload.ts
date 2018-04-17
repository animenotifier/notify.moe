import AnimeNotifier from "../AnimeNotifier"
import StatusMessage from "../StatusMessage"
import { bytesHumanReadable, fetchWithProgress } from "../Utils"

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
		let fileSize = reader.result.byteLength

		if(fileSize === 0) {
			arn.statusMessage.showError("File is empty")
			return
		}

		arn.statusMessage.showInfo(`Preparing to upload ${fileType} (${bytesHumanReadable(fileSize)})`, -1)

		try {
			let responseText = await fetchWithProgress(endpoint, {
				method: "POST",
				credentials: "include",
				headers: {
					"Content-Type": "application/octet-stream"
				},
				body: reader.result
			}, e => {
				let progress = e.loaded / (e.lengthComputable ? e.total : fileSize) * 100
				arn.statusMessage.showInfo(`Uploading ${fileType}...${progress.toFixed(1)}%`, -1)
			})

			arn.statusMessage.showInfo(`Successfully uploaded your new ${fileType}.`)

			if(endpoint === "/api/upload/avatar") {
				// We received the new avatar URL
				updateSideBarAvatar(responseText)
			}
		} catch(err) {
			arn.statusMessage.showError(`Failed uploading your new ${fileType}.`)
			console.error(err)
		}
	}

	arn.statusMessage.showInfo(`Reading ${fileType} from disk...`, -1)
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