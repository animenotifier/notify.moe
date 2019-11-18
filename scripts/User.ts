export default class User {
	public id: string
	private proExpires: string

	public constructor(id: string) {
		this.id = id
		this.sync()
	}

	public IsPro(): boolean {
		return new Date() < new Date(this.proExpires)
	}

	private async sync() {
		const response = await fetch(`/api/user/${this.id}`)
		const json = await response.json()
		Object.assign(this, json)
	}
}
