// Save new data from an input field
// export function save(arn: AnimeNotifier, input: HTMLInputElement | HTMLTextAreaElement) {
	// let apiObject: HTMLElement
	// let parent = input as HTMLElement

	// while(parent = parent.parentElement) {
	// 	if(parent.classList.contains("api-object")) {
	// 		apiObject = parent
	// 		break
	// 	}
	// }

	// if(!apiObject) {
	// 	throw "API object not found"
	// }

	// let request = apiObject["api-fetch"]

	// request.then(obj => {
	// 	obj[input.id] = input.value
	// 	console.log(obj)
	// })
// }

// updateAPIObjects() {
// 	for(let element of findAll(".api-object")) {
// 		let apiObject = element

// 		apiObject["api-fetch"] = fetch(element.dataset.api).then(response => response.json())
// 	}
// }