<script>
	<?php include("scripts/webfont.js"); ?>
</script>

<script>
  (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
  (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
  m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
  })(window,document,'script','//www.google-analytics.com/analytics.js','ga');

  ga('create', 'UA-58621284-1', 'auto');
  ga('require', 'displayfeatures');
  ga('send', 'pageview');
</script>

<script>
	var $container = $('#container');
	var $headerContainer = $('#header-container');
	var $footer = $('#footer');
	var $contentContainer = $('#content-container');

	// Recalculate size
	function recalculateContentSize() {
		var minContentHeight =
			$container.outerHeight(true)
			- $headerContainer.outerHeight(true);
		
		$contentContainer.css("min-height", parseInt(minContentHeight) + "px");
		$footer.show();
	}

	// Page handler
	setPageHandler(function(pageId) {
		recalculateContentSize();
	})

	// Resize
	$(window).resize(function() {
		// Calculate content size
		recalculateContentSize();
	});

	recalculateContentSize();

	// Font
	WebFont.load({
		google: {
			families: ['Open Sans', 'Open Sans:bold']
		},
		active: function() {
			$("#container").animate({
				'opacity' : '1'
			}, 290, 'linear');

			$("#title").animate({
				'opacity' : '1',
				'letter-spacing' : '6px'
			}, 290, 'linear');
		}
	});
</script>