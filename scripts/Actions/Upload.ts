import AnimeNotifier from "../AnimeNotifier"
import { bytesHumanReadable, uploadWithProgress } from "../Utils"

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
	input.value = null

	input.onchange = async () => {
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
			let previews = document.getElementsByClassName(button.id + "-preview")
			let dataURL = await readImageAsDataURL(file)
			let img = await loadImage(dataURL)

			switch(endpoint) {
				case "/api/upload/avatar":
					if(img.naturalWidth <= 280 || img.naturalHeight < 280) {
						arn.statusMessage.showError(`Your image has a resolution of ${img.naturalWidth} x ${img.naturalHeight} pixels which is too small. Recommended: 560 x 560. Minimum: 280 x 280.`, 8000)
						return
					}
					break

				case "/api/upload/cover":
					if(img.naturalWidth <= 960 || img.naturalHeight < 225) {
						arn.statusMessage.showError(`Your image has a resolution of ${img.naturalWidth} x ${img.naturalHeight} pixels which is too small. Recommended: 1920 x 450. Minimum: 960 x 225.`, 8000)
						return
					}
					break
			}

			previewImage(dataURL, endpoint, previews)
		}

		uploadFile(file, fileType, endpoint, arn)
	}

	input.click()
}

// Upload file
function uploadFile(file: File, fileType: string, endpoint: string, arn: AnimeNotifier) {
	let reader = new FileReader()

	reader.onloadend = async () => {
		let result = reader.result as ArrayBuffer
		let fileSize = result.byteLength

		if(fileSize === 0) {
			arn.statusMessage.showError("File is empty")
			return
		}

		arn.statusMessage.showInfo(`Preparing to upload ${fileType} (${bytesHumanReadable(fileSize)})`, -1)

		try {
			let responseText = await uploadWithProgress(endpoint, {
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

// Read image as data URL
function readImageAsDataURL(file: File): Promise<string> {
	return new Promise((resolve, reject) => {
		let reader = new FileReader()

		reader.onloadend = () => {
			let dataURL = reader.result as string
			resolve(dataURL)
		}

		reader.onerror = event => {
			reader.abort()
			reject(event)
		}

		reader.readAsDataURL(file)
	})
}

// Load image and resolve when loading has finished
function loadImage(url: string): Promise<HTMLImageElement> {
	return new Promise((resolve, reject) => {
		let img = new Image()
		img.src = url

		img.onload = () => {
			resolve(img)
		}

		img.onerror = error => {
			reject(error)
		}
	})
}

// Preview image
function previewImage(dataURL: string, endpoint: string, previews: HTMLCollectionOf<Element>) {
	if(endpoint === "/api/upload/avatar") {
		let svgPreview = document.getElementById("avatar-input-preview-svg") as HTMLImageElement

		if(svgPreview) {
			svgPreview.classList.add("hidden")
		}
	}

	for(let preview of previews) {
		let img = preview as HTMLImageElement
		img.classList.remove("hidden")

		// Make not found images visible again
		if(img.classList.contains("lazy")) {
			img.classList.remove("element-not-found")
			img.classList.add("element-found")
		}

		img.src = dataURL
	}
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