var kaze = {};

kaze.fadeSpeed = 250;
kaze.$navigation = $('#navigation');
kaze.$container = $('#container');
kaze.$content = $('#content');

kaze.ajaxifyLinks = function() {
	$('.ajax').each(function() {
		$(this).removeClass('ajax');
	}).click(function(e) {
		var url = this.getAttribute('href');

		e.preventDefault();
		e.stopPropagation();

        if(url === window.location.pathname)
            return;

        history.pushState(null, null, url);

        if(kaze.$navigation.offset().top < 0)
            kaze.scrollToElement(kaze.$navigation);

        $.get('/_' + url, function(response) {
            kaze.fadeContent(kaze.$content, response);
        });
	});
};

kaze.fadeContent = function($content, response) {
	$content.stop().fadeOut(kaze.fadeSpeed, function() {
		$content.html(response).fadeIn(kaze.fadeSpeed, function() {
			kaze.ajaxifyLinks();
		});
	});
};

kaze.scrollToElement = function(element, time) {
    time = (time !== undefined) ? time : kaze.fadeSpeed * 2;

    kaze.$container.animate({
        scrollTop: kaze.$container.scrollTop() + element.offset().top
    }, time);
}

kaze.ajaxifyLinks();