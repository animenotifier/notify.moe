'use strict'

arn.repeatedly(5 * minutes, () => {
	if(global.gc)
		global.gc()
})