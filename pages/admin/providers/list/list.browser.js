window.confirmMatch = function(provider, providerId) {
	kaze.postJSON('/api/matches/confirm', {
		provider: provider,
		providerId: providerId
	}).then(kaze.content.reload);
};