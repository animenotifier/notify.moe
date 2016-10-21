package main

// Anime ...
type Anime struct {
	ID          int        `as:"id"`
	Title       AnimeTitle `as:"title"`
	Description string     `as:"description"`
}

// AnimeTitle ...
type AnimeTitle struct {
	Romaji   string `as:"romaji"`
	English  string `as:"english"`
	Japanese string `as:"japanese"`
}
