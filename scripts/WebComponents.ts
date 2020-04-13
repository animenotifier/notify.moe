import ToolTip from "./Elements/tool-tip/tool-tip"
import SVGIcon from "./Elements/svg-icon/svg-icon"

export function register() {
	if(!("customElements" in window)) {
		console.warn("Web components not supported in your current browser")
		return
	}

	// Custom element names must have a dash in their name
	const elements = new Map<string, Function>([
		["tool-tip", ToolTip],
		["svg-icon", SVGIcon]
	])

	// Register all custom elements
	for(const [tag, definition] of elements.entries()) {
		window.customElements.define(tag, definition as CustomElementConstructor)
	}
}
