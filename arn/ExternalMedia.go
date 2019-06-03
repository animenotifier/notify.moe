package arn

// Register a list of supported media services.
func init() {
	DataLists["media-services"] = []*Option{
		{"Youtube", "Youtube"},
		{"SoundCloud", "SoundCloud"},
		{"DailyMotion", "DailyMotion"},
	}
}

// ExternalMedia ...
type ExternalMedia struct {
	Service   string `json:"service" editable:"true" datalist:"media-services"`
	ServiceID string `json:"serviceId" editable:"true"`
}

// EmbedLink returns the embed link used in iframes for the given media.
func (media *ExternalMedia) EmbedLink() string {
	switch media.Service {
	case "SoundCloud":
		return "//w.soundcloud.com/player/?url=https://api.soundcloud.com/tracks/" + media.ServiceID + "?auto_play=false&hide_related=true&show_comments=false&show_user=false&show_reposts=false&visual=true"
	case "Youtube":
		return "//youtube.com/embed/" + media.ServiceID + "?showinfo=0"
	case "DailyMotion":
		return "//www.dailymotion.com/embed/video/" + media.ServiceID
	case "NicoVideo":
		return "//ext.nicovideo.jp/thumb/" + media.ServiceID
	default:
		return ""
	}
}
