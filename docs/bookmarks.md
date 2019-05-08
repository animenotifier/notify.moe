# Bookmarks

## Live/Dev switcher

Create a bookmark in your browser and set this code as the URL:

```js
javascript:(() => {
	location = location.href.indexOf('://beta.') === -1 ?
	location.href.replace('://', '://beta.') : location.href.replace('://beta.', '://');
})();
```

Clicking this bookmark will let you switch between `notify.moe` (live) and `beta.notify.moe` (development).