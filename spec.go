package muniverse

// An EnvSpec contains meta-data about an environment.
type EnvSpec struct {
	Name    string
	BaseURL string
	Width   int
	Height  int

	KeyWhitelist []string
}

var EnvSpecs = []*EnvSpec{
	{
		Name:    "KatanaFruits-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/k/katana-fruits/v1/",
		Width:   320,
		Height:  427,
	},

	{
		Name:    "MiniRaceRush-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/m/mini-race-rush/v1/",
		Width:   320,
		Height:  497,

		KeyWhitelist: []string{
			"KeyLeft",
			"KeyRight",
		},
	},

	{
		Name:    "Jewelish-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/j/jewelish/10-4bc4e/",
		Width:   320,
		Height:  433,
	},

	{
		Name:    "BubblesShooter-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/b/bubbles-shooter/11-2be95/",
		Width:   522,
		Height:  348,

		KeyWhitelist: []string{
			"Space",
		},
	},

	{
		Name:    "PenguinSkip-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/p/penguin-skip/72b65ff0/",
		Width:   480,
		Height:  320,

		KeyWhitelist: []string{
			"KeyLeft",
			"KeyRight",
		},
	},
}

// SpecForName finds the first entry in EnvSpecs with the
// given name.
// If no entry is found, nil is returned.
func SpecForName(name string) *EnvSpec {
	for _, s := range EnvSpecs {
		if s.Name == name {
			return s
		}
	}
	return nil
}
