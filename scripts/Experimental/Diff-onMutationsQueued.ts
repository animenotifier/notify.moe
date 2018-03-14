// let diffStart: number
// let diffMountableCount: number

// Diff.onStart = () => {
// 	diffMountableCount = 0
// }

// Diff.onMutationsQueued = elem => {
// 	Diff.mutations.queue(() => {
// 		// Mountable
// 		if(elem.classList.contains("mountable")) {
// 			diffMountableCount++

// 			let now = Date.now()

// 			if(diffMountableCount === 1) {
// 				diffStart = now
// 			}

// 			const mountTime = diffMountableCount * 20 + diffStart

// 			if(mountTime > now && mountTime < now + 800) {
// 				let mount = () => {
// 					if(Date.now() >= mountTime) {
// 						elem.classList.add("mounted")
// 					} else {
// 						Diff.mutations.queue(mount)
// 					}
// 				}

// 				Diff.mutations.queue(mount)
// 			} else {
// 				elem.classList.add("mounted")
// 			}
// 		} else if(elem.classList.contains("lazy")) {
// 			this.lazyLoadElement(elem)
// 		}
// 	})
// }