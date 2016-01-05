window.save = function(e) {
	var item = e.target;

	if(document.saving)
		return;

	document.saving = true;

	var key = item.id;
	var value = item.value;

	item.classList.add('saving');
	kaze.content.style.cursor = 'wait';

	kaze.postJSON('/api/users/me', {
		function: 'save',
		key: key,
		value: value
	}, function(error, response) {
		if(error)
			console.log(error);

		window.postSave(key, response);

		kaze.get('/_' + location.pathname, function(error, newPageCode) {
			if(error)
				console.log(error);

			var focusedElementId = document.activeElement.id;
			var focusedElementValue = document.activeElement.value;

			kaze.onResponse(newPageCode);

			// Re-focus previously selected element
			if(focusedElementId) {
				var focusedElement = document.getElementById(focusedElementId);

				if(focusedElement) {
					focusedElement.value = focusedElementValue;

					if(focusedElement.select)
						focusedElement.select();
					else if(focusedElement.focus)
						focusedElement.focus();
				}
			}

			kaze.content.style.cursor = 'auto';
			document.saving = false;
		})
	});
};

window.postSave = function(key, value) {
	switch(key) {
		case 'nick':
			value = value.replace(/\s+/g, '');
			var oldPath = '/+' + userName;
			var newPath = '/+' + value;

			//window.history.pushState('', document.title, newPath);

			var links = document.querySelectorAll('a');
			for(var l = 0; l < links.length; ++l) {
				var link = links[l];
				if(link.href.endsWith(oldPath))
					link.href = newPath;
			}

			break;
	}
};

var userName = document.getElementById('nick').value;
var myNodeList = document.querySelectorAll('.save-on-change');

for(var i = 0; i < myNodeList.length; ++i) {
	var element = myNodeList[i];
	element.onchange = window.save;
}

// Push notifications
var pushButton = document.querySelector('.push-button');

if(pushButton) {
	pushButton.addEventListener('click', function() {
		if(isPushEnabled) {
			unsubscribe();
		} else {
			subscribe();
		}
	});

	// Check that service workers are supported, if so, progressively
	// enhance and add push messaging support, otherwise continue without it.
	if('serviceWorker' in navigator) {
		console.log('Registering service worker...');
		navigator.serviceWorker.register('/web/service-worker.js').then(initialiseState);
	} else {
		console.warn('Service workers aren\'t supported in this browser.');
	}
} else {
	console.warn('Push button not found.');
}