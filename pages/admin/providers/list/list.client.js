window.confirmMatch = function(provider, providerId) {
	$.post('/api/matches/confirm', {
		provider: provider,
		providerId: providerId
	}).then(aero.content.reload);
};