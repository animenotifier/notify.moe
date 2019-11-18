import bytesHumanReadable from "scripts/Utils/bytesHumanReadable"
import uploadWithProgress from "scripts/Utils/uploadWithProgress"
import AnimeNotifier from "../AnimeNotifier"

// Select file
export function selectFile(arn: AnimeNotifier, button: HTMLButtonElement) {
	const fileType = button.dataset.type
	const endpoint = button.dataset.endpoint

	if(endpoint === "/api/upload/user/cover" && arn.user && !arn.user.IsPro()) {
		alert("Please buy a PRO account to use this feature.")
		return
	}

	// Click on virtual file input element
	const input = document.createElement("input")
	input.setAttribute("type", "file")
	input.value = ""

	input.onchange = async () => {
		if(!fileType || !endpoint) {
			console.error("Missing data-type or data-endpoint:", button)
			return
		}

		if(!input.files || input.files.length === 0) {
			return
		}

		const file = input.files[0]

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
			const previews = document.getElementsByClassName(button.id + "-preview")
			const dataURL = await readImageAsDataURL(file)
			const img = await loadImage(dataURL)

			switch(endpoint) {
				case "/api/upload/user/image":
					if(img.naturalWidth <= 280 || img.naturalHeight < 280) {
						arn.statusMessage.showError(`Your image has a resolution of ${img.naturalWidth} x ${img.naturalHeight} pixels which is too small. Recommended: 560 x 560. Minimum: 280 x 280.`, 8000)
						return
					}
					break

				case "/api/upload/user/cover":
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
	const reader = new FileReader()

	reader.onloadend = async () => {
		const result = reader.result as ArrayBuffer
		const fileSize = result.byteLength

		if(fileSize === 0) {
			arn.statusMessage.showError("File is empty")
			return
		}

		arn.statusMessage.showInfo(`Preparing to upload ${fileType} (${bytesHumanReadable(fileSize)})`, -1)

		try {
			const responseText = await uploadWithProgress(endpoint, {
				method: "POST",
				credentials: "include",
				headers: {
					"Content-Type": "application/octet-stream"
				},
				body: reader.result
			}, e => {
				const progress = e.loaded / (e.lengthComputable ? e.total : fileSize) * 100
				arn.statusMessage.showInfo(`Uploading ${fileType}...${progress.toFixed(1)}%`, -1)
			})

			arn.statusMessage.showInfo(`Successfully uploaded your new ${fileType}.`)

			if(endpoint === "/api/upload/user/image") {
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
		const reader = new FileReader()

		reader.onloadend = () => {
			const dataURL = reader.result as string
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
		const img = new Image()
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
	if(endpoint === "/api/upload/user/image") {
		const svgPreview = document.getElementById("avatar-input-preview-svg") as HTMLImageElement

		if(svgPreview) {
			svgPreview.classList.add("hidden")
		}
	}

	for(const preview of previews) {
		const img = preview as HTMLImageElement
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
	const sidebar = document.getElementById("sidebar") as HTMLElement
	const userImage = sidebar.getElementsByClassName("user-image")[0] as HTMLImageElement
	const lazyLoad = userImage["became visible"]

	if(lazyLoad) {
		userImage.dataset.src = url
		lazyLoad()
	} else {
		location.reload()
	}
}
