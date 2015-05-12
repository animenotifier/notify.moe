var keystone = require('keystone'),
	gravatar = require('gravatar'),
	Types = keystone.Field.Types;

/**
 * User Model
 * ==========
 */

var User = new keystone.List('User', {
	map: {
		name: 'nick'
	},
	track: true,
	defaultSort: '-createdAt'
});

User.add({
	nick: {
		type: String,
		initial: true,
		required: true,
		index: true
	},
	name: {
		type: Types.Name,
		initial: true,
		index: true
	},
	email: {
		type: Types.Email,
		initial: true,
		required: true,
		index: true
	},
	password: {
		type: Types.Password,
		initial: true,
		required: true
	}
}, 'Permissions', {
	isAdmin: {
		type: Boolean,
		index: true
	}
}, 'Dates', {
	birthDay: {
		type: Types.Date
	}
}, 'Payments', {
	balance: {
		type: Types.Money,
		format: '$0,0.00',
		default: 0
	}
});

// Helper functions
var getAge = function(d1, d2) {
	d2 = d2 || new Date();
	var diff = d2.getTime() - d1.getTime();
	return Math.floor(diff / (1000 * 60 * 60 * 24 * 365.25));
};

// Provide access to Keystone
User.schema.virtual('canAccessKeystone').get(function() {
	return this.isAdmin;
});

User.schema.virtual('age').get(function() {
	return (this.birthDay !== undefined) ? getAge(this.birthDay) : 0;
});

User.schema.virtual('avatarUrl').get(function() {
	return (this.email !== undefined) ? gravatar.url(this.email, {}) : '';
});

/**
 * Registration
 */

User.defaultColumns = 'nick, name, email, isAdmin';
User.register();