var keystone = require('keystone'),
	Types = keystone.Field.Types;

/**
 * Anime Model
 * ==========
 */

var Anime = new keystone.List('Anime');

Anime.add({
	
});

/**
 * Registration
 */

Anime.defaultColumns = 'titles';
Anime.register();