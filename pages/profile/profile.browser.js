// Save
window.save = function(e) {
	var item = e.srcElement;

	if(item.saving)
		return;

	item.saving = true;

	var key = item.id;
	var value = item.value;

	item.classList.add('saving');

	console.log(key, value);
	kaze.postJSON('/api/users/me', {
		function: 'save',
		key: key,
		value: value
	}, function(response) {
		//kaze.onResponse(response);
		console.log('Saved.')
	});
};

var myNodeList = document.querySelectorAll('.save-on-change');

for(var i = 0; i < myNodeList.length; ++i) {
	var element = myNodeList[i];
	element.onchange = window.save;
}