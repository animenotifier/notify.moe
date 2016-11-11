document.addEventListener('keydown', e => {
	// Alt + A = Staff info
	if(e.keyCode === 65 && e.altKey) {
		let staffInfo = $('staff-info')

		if(staffInfo.style.display !== 'block') {
			staffInfo.style.display = 'block'
		} else {
			staffInfo.style.display = 'none'
		}
	}
})