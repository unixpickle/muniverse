package muniverse

// An EnvSpec contains meta-data about an environment.
type EnvSpec struct {
	Name    string
	BaseURL string
	Width   int
	Height  int

	KeyWhitelist []string

	VariantOf   string
	VariantOpts string
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
		Name:    "BurninRubber-v1",
		BaseURL: "http://games.cdn.famobi.com/html5games/b/burnin-rubber/11-bc4df/",
		Width:   320,
		Height:  480,

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
			"Space",
		},

		VariantOf:   "BurninRubber-v0",
		VariantOpts: "{\"car\":0,\"weapon\":1}",
	},

	{
		Name:    "BurninRubber-v2",
		BaseURL: "http://games.cdn.famobi.com/html5games/b/burnin-rubber/11-bc4df/",
		Width:   320,
		Height:  480,

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
			"Space",
		},

		VariantOf:   "BurninRubber-v0",
		VariantOpts: "{\"car\":1,\"weapon\":0}",
	},

	{
		Name:    "BurninRubber-v3",
		BaseURL: "http://games.cdn.famobi.com/html5games/b/burnin-rubber/11-bc4df/",
		Width:   320,
		Height:  480,

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
			"Space",
		},

		VariantOf:   "BurninRubber-v0",
		VariantOpts: "{\"car\":1,\"weapon\":1}",
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
		Name:    "GoldMinerTom-v1",
		BaseURL: "http://games.cdn.famobi.com/html5games/g/gold-miner-tom/v120/",
		Width:   320,
		Height:  480,

		VariantOf:   "GoldMinerTom-v0",
		VariantOpts: "{}",
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

	{
		Name:    "RedHead-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/r/red-head/v1/",
		Width:   480,
		Height:  320,

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
		},
	},

	{
		Name:    "NutRush-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/n/nut-rush/v030/",
		Width:   480,
		Height:  320,
	},

	{
		Name:    "NutRush-v1",
		BaseURL: "http://games.cdn.famobi.com/html5games/n/nut-rush/v030/",
		Width:   480,
		Height:  320,

		VariantOf:   "NutRush-v0",
		VariantOpts: "{\"level\":2}",
	},

	{
		Name:    "NutRush-v2",
		BaseURL: "http://games.cdn.famobi.com/html5games/n/nut-rush/v030/",
		Width:   480,
		Height:  320,

		VariantOf:   "NutRush-v0",
		VariantOpts: "{\"level\":3}",
	},

	{
		Name:    "KibaKumbaShadowRun-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/k/kk-shadow-run/v3/",
		Width:   560,
		Height:  320,
	},

	{
		Name:    "KibaKumbaShadowRun-v1",
		BaseURL: "http://games.cdn.famobi.com/html5games/k/kk-shadow-run/v3/",
		Width:   560,
		Height:  320,

		VariantOf:   "KibaKumbaShadowRun-v0",
		VariantOpts: "{\"character\":1}",
	},

	{
		Name:    "Knightower-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/k/knightower/3-c82d31/",
		Width:   320,
		Height:  480,

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
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
