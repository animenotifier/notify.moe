var keystone = require('keystone'),
	Types = keystone.Field.Types;

/**
 * Anime Model
 * ==========
 */

var Anime = new keystone.List('Anime', {
	map: {
		name: 'romajiName'
	},
	track: true,
	defaultSort: '-createdAt',
	plural: 'Anime',
	label: 'Anime',
	path: 'anime',
	schema: {
		collection: 'anime'
	}
});

Anime.add({
	romajiName: {
		type: String,
		initial: true,
		required: true,
		index: true
	},
	japaneseName: {
		type: String,
		initial: true,
		index: true
	},
	synonyms: {
		type: Types.TextArray
	}
});

/**
 * Registration
 */

Anime.defaultColumns = 'romajiName, japaneseName';
Anime.register();