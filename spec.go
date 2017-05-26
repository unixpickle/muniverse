package muniverse

// An EnvSpec contains meta-data about an environment.
type EnvSpec struct {
	Name    string
	BaseURL string
	Width   int
	Height  int
}

var EnvSpecs = []*EnvSpec{
	{
		Name:    "KatanaFruits-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/k/katana-fruits/v1/",
		Width:   320,
		Height:  427,
	},
}
