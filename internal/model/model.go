package model

type Technology struct {
	Name string `json:"name"`
}

// LanguageCycle struct for standardizing the data
type LanguageCycle struct {
	Cycle           string `json:"cycle"`
	ReleaseDate     string `json:"releaseDate"`
	Eol             string `json:"eol"`
	Latest          string `json:"latest"`
	Link            string `json:"link"`
	Lts             string `json:"lts"`
	Support         string `json:"support"`
	Discontinued    string `json:"discontinued"`
	ExtendedSupport string `json:"extendedSupport"`
}

var TechnologiesCycle []LanguageCycle

func init() {
	TechnologiesCycle = make([]LanguageCycle, 0)
}
