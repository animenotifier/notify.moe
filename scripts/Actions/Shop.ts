import AnimeNotifier from "../AnimeNotifier"

// Charge up
export function chargeUp(arn: AnimeNotifier, button: HTMLElement) {
	let amount = button.dataset.amount

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

		console.log(payment)
		let link = payment.links.find(link => link.rel === "approval_url")

		if(!link) {
			throw "Error finding PayPal payment link"
		}

		arn.statusMessage.showInfo("Redirecting to PayPal...", 5000)

		let url = link.href
		window.location.href = url
	})
	.catch(err => arn.statusMessage.showError(err))
	.then(() => arn.loading(false))
}

// Buy item
export function buyItem(arn: AnimeNotifier, button: HTMLElement) {
	let itemId = button.dataset.itemId
	let itemName = button.dataset.itemName
	let price = button.dataset.price

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
		if(body !== "ok") {
			throw body
		}

		return arn.reloadContent()
	})
	.then(() => arn.statusMessage.showInfo(`You bought ${itemName} for ${price} gems. Check out your inventory to confirm the purchase.`, 4000))
	.catch(err => arn.statusMessage.showError(err))
	.then(() => arn.loading(false))
}