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
			"ArrowLeft",
			"ArrowRight",
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
			"ArrowLeft",
			"ArrowRight",
		},
	},

	{
		Name:    "MinimalDots-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/m/minimal-dots/v2/",
		Width:   320,
		Height:  480,

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
		},
	},

	{
		Name:    "BurninRubber-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/b/burnin-rubber/11-bc4df/",
		Width:   320,
		Height:  480,

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
			"Space",
		},
	},

	{
		Name:    "OnetConnectClassic-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/o/onet-connect-classic/v160/",
		Width:   480,
		Height:  320,
	},

	{
		Name:    "DontCrash-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/d/dont-crash/v2/",
		Width:   480,
		Height:  320,
	},

	{
		Name:    "SmartyBubbles-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/s/smarty-bubbles/v030/",
		Width:   320,
		Height:  480,
	},

	{
		Name:    "StreetPursuit-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/s/street-pursuit/7-8cd52b/",
		Width:   320,
		Height:  497,

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
		},
	},

	{
		Name:    "GoldMinerTom-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/g/gold-miner-tom/v120/",
		Width:   480,
		Height:  320,
	},

	{
		Name:    "UfoRun-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/u/ufo-run/6f9ac9a0/",
		Width:   480,
		Height:  321,

		KeyWhitelist: []string{
			"Space",
			"ArrowUp",
		},
	},

	{
		Name:    "BirdyRush-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/b/birdy-rush/v040/",
		Width:   320,
		Height:  533,

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
		},
	},

	{
		Name:    "HexBlitz-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/h/hex-blitz/v270/",
		Width:   480,
		Height:  320,
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
