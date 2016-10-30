export function auth(req, res, role: string): boolean {
	if(!req.user) {
		res.end('Not logged in!')
		return false
	}

	if(req.user.role !== 'admin' && req.user.role !== role) {
		res.end('Not authorized to view this page!')
		return false
	}

	return true
}