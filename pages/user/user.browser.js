window.save = function(e) {
	var item = e.srcElement;

	if(document.saving)
		return;

	document.saving = true;

	var key = item.id;
	var value = item.value;

	item.classList.add('saving');

	kaze.postJSON('/api/users/me', {
		function: 'save',
		key: key,
		value: value
	}, function(response) {
		if(response) {
			kaze.onResponse(response);
			return;
		}

		window.postSave(key, value);

		kaze.get('/_' + location.pathname, function(newPageCode) {
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

			document.saving = false;
		})
	});
};

window.postSave = function(key, value) {
	switch(key) {
		case 'nick':
			var oldPath = window.location.pathname;
			var newPath = '/+' + value;

			window.history.pushState('', document.title, newPath);

			var links = document.querySelectorAll('a');
			for(var l = 0; l < links.length; ++l) {
				var link = links[l];
				if(link.href.endsWith(oldPath))
					link.href = newPath;
			}

			break;
	}
};

var myNodeList = document.querySelectorAll('.save-on-change');

for(var i = 0; i < myNodeList.length; ++i) {
	var element = myNodeList[i];
	element.onchange = window.save;
}