function makeSaveable(apiEndpoint: string, postSaveCallback?: (string, any) => void) {
	$.save = function(e) {
		var item = e.target;

		if($.saving)
			return;

		$.saving = true;

		var key = item.id;
		var value = item.value ? item.value : '';
		var old = item.dataset.old ? item.dataset.old : '';

		item.classList.add('saving');
		$.content.style.cursor = 'wait';

		$.post(apiEndpoint, {
			function: 'save',
			key: key,
			value: value,
			old: old
		}).then(function(response) {
			if(postSaveCallback)
				postSaveCallback(key, response);
		}).catch(function(error) {
			console.error(error);
		}).then(function() {
			$.get('/_' + location.pathname).then(function(newPageCode) {
				var focusedElementId = document.activeElement.id;
				var focusedElementValue = (<HTMLInputElement> document.activeElement).value;

				$.setContent(newPageCode);

				// Re-focus previously selected element
				if(focusedElementId) {
					var focusedElement = $(focusedElementId);

					if(focusedElement) {
						focusedElement.value = focusedElementValue;

						if(focusedElement.select)
							focusedElement.select();
						else if(focusedElement.focus)
							focusedElement.focus();
					}
				}

				$.content.style.cursor = 'auto';
				$.saving = false;
			});
		});
	};

	var myNodeList = document.querySelectorAll('.save-on-change');

	for(var i = 0; i < myNodeList.length; ++i) {
		var element = <HTMLInputElement> myNodeList[i];
		element.onchange = $.save;
		element.dataset['old'] = element.value;
	}
}