window.confirmMatch = function(provider, providerId) {
	aero.postJSON('/api/matches/confirm', {
		provider: provider,
		providerId: providerId
	}).then(aero.content.reload);
};