package arn

// Icons
var forumIcons = map[string]string{
	"general":    "paperclip",
	"news":       "newspaper-o",
	"anime":      "television",
	"update":     "cubes",
	"suggestion": "lightbulb-o",
	"bug":        "bug",
}

// GetForumIcon returns the unprefixed icon class name for the forum.
func GetForumIcon(category string) string {
	icon, exists := forumIcons[category]

	if exists {
		return icon
	}

	return "comments"
}
