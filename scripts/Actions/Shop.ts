import AnimeNotifier from "../AnimeNotifier"

// Charge up
export function chargeUp(arn: AnimeNotifier, button: HTMLElement) {
	const amount = button.dataset.amount

	arn.loading(true)
	arn.statusMessage.showInfo("Creating PayPal transaction... This might take a few seconds.")

	fetch("/api/paypal/payment/create", {
		method: "POST",
		body: amount,
		credentials: "same-origin"
	})
	.then(response => response.json())
	.then(payment => {
		if(!payment || !payment.links) {
			throw "Error creating PayPal payment"
		}

		const link = payment.links.find(link => link.rel === "approval_url")

		if(!link) {
			throw "Error finding PayPal payment link"
		}

		arn.statusMessage.showInfo("Redirecting to PayPal...", 5000)

		const url = link.href
		window.location.href = url
	})
	.catch(err => arn.statusMessage.showError(err))
	.then(() => arn.loading(false))
}

// Toggle fade
export function toggleFade(_: AnimeNotifier, button: HTMLElement) {
	const elementId = button.dataset.elementId

	if(!elementId) {
		console.error("Missing element ID:", elementId)
		return
	}

	const element = document.getElementById(elementId)

	if(!element) {
		console.error("Invalid element ID:", elementId)
		return
	}

	if(element.classList.contains("fade-out")) {
		element.classList.remove("fade-out")
	} else {
		element.classList.add("fade-out")
	}
}

// Buy item
export function buyItem(arn: AnimeNotifier, button: HTMLElement) {
	const itemId = button.dataset.itemId
	const itemName = button.dataset.itemName
	const price = button.dataset.price

	if(!confirm(`Would you like to buy ${itemName} for ${price} gems?`)) {
		return
	}

	arn.loading(true)

	fetch(`/api/shop/buy/${itemId}/1`, {
		method: "POST",
		credentials: "same-origin"
	})
	.then(response => response.text())
	.then(body => {
		if(body !== "") {
			throw body
		}

		return arn.reloadContent()
	})
	.then(() => arn.statusMessage.showInfo(`You bought ${itemName} for ${price} gems. Check out your inventory to confirm the purchase.`, 4000))
	.catch(err => arn.statusMessage.showError(err))
	.then(() => arn.loading(false))
}
