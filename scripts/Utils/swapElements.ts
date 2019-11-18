// swapElements assumes that both elements have valid parent nodes.
export default function swapElements(a: Node, b: Node) {
	const bParent = b.parentNode as Node
	const bNext = b.nextSibling

	// Special case for when a is the next sibling of b
	if(bNext === a) {
		// Just put a before b
		bParent.insertBefore(a, b)
	} else {
		// Insert b right before a
		(a.parentNode as Node).insertBefore(b, a)

		// Now insert a where b was
		if(bNext) {
			// If there was an element after b, then insert a right before that
			bParent.insertBefore(a, bNext)
		} else {
			// Otherwise just append it as the last child
			bParent.appendChild(a)
		}
	}
}
