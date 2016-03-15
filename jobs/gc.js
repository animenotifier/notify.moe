'use strict'

arn.repeatedly(5 * minutes, () => {
	Object.keys(require.cache).forEach(key => delete require.cache[key])
	
	if(global.gc)
		global.gc()
})