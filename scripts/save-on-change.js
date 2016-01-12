function makeSaveable(apiEndpoint, postSaveCallback) {
	window.save = function(e) {
		var item = e.target;

		if(document.saving)
			return;

		document.saving = true;

		var key = item.id;
		var value = item.value;
		var old = item.dataset.old;

		item.classList.add('saving');
		kaze.content.style.cursor = 'wait';

		kaze.postJSON(apiEndpoint, {
			function: 'save',
			key: key,
			value: value,
			old: old
		}, function(error, response) {
			if(error)
				console.error(error.stack);

			if(postSaveCallback)
				postSaveCallback(key, response);

			kaze.get('/_' + location.pathname, function(error, newPageCode) {
				if(error)
					console.error(error.stack);

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

	var myNodeList = document.querySelectorAll('.save-on-change');

	for(var i = 0; i < myNodeList.length; ++i) {
		var element = myNodeList[i];
		element.onchange = window.save;
		element.dataset.old = element.value;
	}
}