export default function bytesHumanReadable(fileSize: number): string {
	let unit = "bytes"

	if(fileSize >= 1024) {
		fileSize /= 1024
		unit = "KB"

		if(fileSize >= 1024) {
			fileSize /= 1024
			unit = "MB"
		}
	}

	return `${fileSize.toFixed(0)} ${unit}`
}
