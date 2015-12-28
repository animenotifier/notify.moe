window.onSave = function(key, value) {
	switch(key) {
		case 'nick':
			var oldPath = window.location.pathname;
			var newPath = '/profile/' + value;

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

var myNodeList = document.querySelectorAll('.editable');

for(var i = 0; i < myNodeList.length; ++i) {
	var item = myNodeList[i];

	item.onclick = function() {
		if(item.editing)
			return;

		item.editing = true;

		var value = item.innerText;
		item.innerHTML = '';

		var input = document.createElement('input');
		input.setAttribute('type', 'text');
		input.setAttribute('value', value);

		var acceptText = function() {
			if(!item.editing)
				return;

			item.editing = false;

			var newValue = input.value;
			item.innerHTML = '';
			item.innerText = newValue;
			item.classList.add('saving');

			kaze.postJSON('/api/users/', {
				key: item.dataset.binding,
				value: newValue
			}, function(response) {
				window.onSave(item.dataset.binding, newValue);
				item.classList.remove('saving');
			});
		};

		input.onkeydown = function(e) {
			if(!e)
				e = window.event;

			var keyCode = e.keyCode || e.which;

			// Escape
			if(keyCode === 27) {
				item.editing = false;
				item.innerHTML = '';
				item.innerText = value;
				return;
			}

			// Return
			if(keyCode !== 13) // && keyCode !== 9)
				return;

			acceptText();
		};

		input.onblur = acceptText;

		item.appendChild(input);
		input.focus();
		input.select();
	};
}