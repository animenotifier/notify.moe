var timeSince = function(start, date) {
	var seconds = Math.floor((start - date) / 1000);

	var interval = Math.floor(seconds / 31536000);
	if(interval > 1)
	    return interval + ' years';

	interval = Math.floor(seconds / 2592000);
	if(interval > 1)
	    return interval + ' months';

	interval = Math.floor(seconds / 86400);
	if(interval > 1)
	    return interval + ' days';

	interval = Math.floor(seconds / 3600);
	if(interval > 1)
	    return interval + ' hours';

	interval = Math.floor(seconds / 60);
	if(interval > 1)
	    return interval + ' minutes';

	return Math.floor(seconds) + ' seconds';
};

kaze.getJSON('https://api.github.com/users/animenotifier/events?clientid=e8fe5e8bcaf6b7ebe0534a93976dca8bdc320ee4&clientsecret=eae6fea79ebe2c919770e0c5e2e38d64d70453d5', function(error, data) {
	var now = new Date();

	document.getElementById('github-events').innerHTML = '<ul>' +
		data
		.filter(function(e) {
			return e.type === 'PushEvent'
		})
		.map(function(e) {
			return e.payload.commits.map(function(commit) {
				return '<li class="commit"><a href="https://github.com/' + e.repo.name + '/commit/' + commit.sha + '" target="_blank" title="' + e.repo.name.substring(e.repo.name.indexOf('/') + 1) + '">' + commit.message.split('\n')[0] + '</a>'
					+ '<span class="datetime"> (' + timeSince(now, new Date(e.created_at)) + ' ago)</span>'
					+ '</li>';
			}).join('');
		}).join('') + '</ul>';
});