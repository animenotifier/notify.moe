export function swapElements(a: Node, b: Node) {
	let parent = b.parentNode
	let bNext = b.nextSibling

	// Special case for when a is the next sibling of b
	if(bNext === a) {
		// Just put a before b
		parent.insertBefore(a, b)
	} else {
		// Insert b right before a
		a.parentNode.insertBefore(b, a)

		// Now insert a where b was
		if(bNext) {
			// If there was an element after b, then insert a right before that
			parent.insertBefore(a, bNext)
		} else {
			// Otherwise just append it as the last child
			parent.appendChild(a)
		}
	}
}