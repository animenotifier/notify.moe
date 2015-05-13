var keystone = require('keystone');

exports = module.exports = function(req, res) {
	var view = new keystone.View(req, res),
		locals = res.locals;

	// locals.section is used to set the currently selected
	// item in the header navigation.
	locals.section = 'statistics';

	User.model
		.find()
		.exec(function(err, users) {
			locals.userCount = users.length;
			view.render('statistics');
		});
};