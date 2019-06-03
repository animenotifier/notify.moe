package arn

// Register a list of gaming servers.
func init() {
	DataLists["ffxiv-servers"] = []*Option{
		{"", ""},
		{"Adamantoise", "Adamantoise"},
		{"Aegis", "Aegis"},
		{"Alexander", "Alexander"},
		{"Anima", "Anima"},
		{"Asura", "Asura"},
		{"Atomos", "Atomos"},
		{"Bahamut", "Bahamut"},
		{"Balmung", "Balmung"},
		{"Behemoth", "Behemoth"},
		{"Belias", "Belias"},
		{"Brynhildr", "Brynhildr"},
		{"Cactuar", "Cactuar"},
		{"Carbuncle", "Carbuncle"},
		{"Cerberus", "Cerberus"},
		{"Chocobo", "Chocobo"},
		{"Coeurl", "Coeurl"},
		{"Diabolos", "Diabolos"},
		{"Durandal", "Durandal"},
		{"Excalibur", "Excalibur"},
		{"Exodus", "Exodus"},
		{"Faerie", "Faerie"},
		{"Famfrit", "Famfrit"},
		{"Fenrir", "Fenrir"},
		{"Garuda", "Garuda"},
		{"Gilgamesh", "Gilgamesh"},
		{"Goblin", "Goblin"},
		{"Gungnir", "Gungnir"},
		{"Hades", "Hades"},
		{"Hyperion", "Hyperion"},
		{"Ifrit", "Ifrit"},
		{"Ixion", "Ixion"},
		{"Jenova", "Jenova"},
		{"Kujata", "Kujata"},
		{"Lamia", "Lamia"},
		{"Leviathan", "Leviathan"},
		{"Lich", "Lich"},
		{"Louisoix", "Louisoix"},
		{"Malboro", "Malboro"},
		{"Mandragora", "Mandragora"},
		{"Masamune", "Masamune"},
		{"Mateus", "Mateus"},
		{"Midgardsormr", "Midgardsormr"},
		{"Moogle", "Moogle"},
		{"Odin", "Odin"},
		{"Omega", "Omega"},
		{"Pandaemonium", "Pandaemonium"},
		{"Phoenix", "Phoenix"},
		{"Ragnarok", "Ragnarok"},
		{"Ramuh", "Ramuh"},
		{"Ridill", "Ridill"},
		{"Sargatanas", "Sargatanas"},
		{"Shinryu", "Shinryu"},
		{"Shiva", "Shiva"},
		{"Siren", "Siren"},
		{"Tiamat", "Tiamat"},
		{"Titan", "Titan"},
		{"Tonberry", "Tonberry"},
		{"Typhon", "Typhon"},
		{"Ultima", "Ultima"},
		{"Ultros", "Ultros"},
		{"Unicorn", "Unicorn"},
		{"Valefor", "Valefor"},
		{"Yojimbo", "Yojimbo"},
		{"Zalera", "Zalera"},
		{"Zeromus", "Zeromus"},
		{"Zodiark", "Zodiark"},
	}
}

// UserAccounts represents a user's accounts on external services.
type UserAccounts struct {
	Facebook struct {
		ID string `json:"id" private:"true"`
	} `json:"facebook"`

	Google struct {
		ID string `json:"id" private:"true"`
	} `json:"google"`

	Twitter struct {
		ID   string `json:"id" private:"true"`
		Nick string `json:"nick" private:"true"`
	} `json:"twitter"`

	Discord struct {
		Nick     string `json:"nick" editable:"true"`
		Verified bool   `json:"verified"`
	} `json:"discord"`

	Osu struct {
		Nick     string  `json:"nick" editable:"true"`
		PP       float64 `json:"pp"`
		Accuracy float64 `json:"accuracy"`
		Level    float64 `json:"level"`
	} `json:"osu"`

	Overwatch struct {
		BattleTag   string `json:"battleTag" editable:"true"`
		SkillRating int    `json:"skillRating"`
		Tier        string `json:"tier"`
	} `json:"overwatch"`

	FinalFantasyXIV struct {
		Nick      string `json:"nick" editable:"true"`
		Server    string `json:"server" editable:"true" datalist:"ffxiv-servers"`
		Class     string `json:"class"`
		Level     int    `json:"level"`
		ItemLevel int    `json:"itemLevel"`
	} `json:"ffxiv"`

	AniList struct {
		Nick string `json:"nick" editable:"true"`
	} `json:"anilist"`

	AnimePlanet struct {
		Nick string `json:"nick" editable:"true"`
	} `json:"animeplanet"`

	MyAnimeList struct {
		Nick string `json:"nick" editable:"true"`
	} `json:"myanimelist"`

	Kitsu struct {
		Nick string `json:"nick" editable:"true"`
	} `json:"kitsu"`
}
