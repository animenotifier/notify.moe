declare module anime {
	export interface Title {
		canonical: string;
		romaji: string;
		english: string;
		japanese: string;
		hiragana: string;
		synonyms: any[];
	}

	export interface AverageColor {
		hue: number;
		saturation: number;
		lightness: number;
	}

	export interface Image {
		extension: string;
		width: number;
		height: number;
		averageColor: AverageColor;
		lastModified: number;
	}

	export interface Count {
		overall: number;
		story: number;
		visuals: number;
		soundtrack: number;
	}

	export interface Rating {
		overall: number;
		story: number;
		visuals: number;
		soundtrack: number;
		count: Count;
	}

	export interface Popularity {
		watching: number;
		completed: number;
		planned: number;
		hold: number;
		dropped: number;
	}

	export interface Trailer {
		service: string;
		serviceId: string;
	}

	export interface Mapping {
		service: string;
		serviceId: string;
	}
}

export interface Anime {
	id: string;
	type: string;
	title: anime.Title;
	summary: string;
	status: string;
	genres: string[];
	startDate: string;
	endDate: string;
	episodeCount: number;
	episodeLength: number;
	source: string;
	image: anime.Image;
	firstChannel: string;
	rating: anime.Rating;
	popularity: anime.Popularity;
	trailers: anime.Trailer[];
	mappings: anime.Mapping[];
	studios: string[];
	producers: string[];
	licensors: string[];
	links?: any;
	created: string;
	createdBy: string;
	edited: string;
	editedBy: string;
}
