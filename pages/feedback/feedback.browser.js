var feedback = document.getElementById('feedback');
feedback.focus();

window.sendFeedback = function() {
	if(!feedback.value)
		return;

	kaze.postJSON('/api/feedback', {
		text: feedback.value
	}).then(function(response) {
		if(response === 'OK') {
			alert('Thank you, your feedback has been sent!');
			feedback.value = '';
			feedback.focus();
		}
	});
};