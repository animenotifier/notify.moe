var $container = $('#container');
var $headerContainer = $('#header-container');
var $footer = $('#footer');
var $contentContainer = $('#content-container');

function recalculateContentSize() {
	var minContentHeight =
		$container.outerHeight(true)
		- $headerContainer.outerHeight(true);

	$contentContainer.css('min-height', parseInt(minContentHeight) + 'px');
	$footer.show();
}

$(window).resize(function() {
	// Calculate content size
	recalculateContentSize();
});

recalculateContentSize();