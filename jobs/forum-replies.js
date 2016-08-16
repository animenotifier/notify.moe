let checkAniListForumReplies = arn.listProviders.AniList.checkForumReplies.bind(arn.listProviders.AniList)

let checkForumReplies = coroutine(function*() {
	yield arn.listProviders.AniList.authorize()
	checkAniListForumReplies()
})

arn.on('new forum reply', function(link, userName) {
	let webhook = 'https://hooks.slack.com/services/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'

	fetch.post({
		url: webhook,
		body: JSON.stringify({
			text: `<${link}|${userName}>`
		})
	}).then(body => {
		console.log(`Sent slack message about a new forum reply from ${userName}`)
	}).catch(error => {
		console.error('Error sending slack message:', error)
	})
})

arn.repeatedly(3 * minutes, checkForumReplies)