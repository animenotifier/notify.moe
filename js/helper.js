// Set object (store objects in localStorage)
Storage.prototype.setObject = function(key, value) {
	this.setItem(key, JSON.stringify(value));
}

// Get object (retrieve objects from localStorage)
Storage.prototype.getObject = function(key) {
	var value = this.getItem(key);
	return value && JSON.parse(value);
}