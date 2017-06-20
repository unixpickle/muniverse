package muniverse

// An EnvSpec contains meta-data about an environment.
type EnvSpec struct {
	Name    string
	BaseURL string
	Width   int
	Height  int
	Options string

	KeyWhitelist []string

	VariantOf string
}

var EnvSpecs = []*EnvSpec{
	{
		Name:    "KatanaFruits-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/k/katana-fruits/v1/",
		Width:   320,
		Height:  427,
		Options: "{}",
	},

	{
		Name:    "MiniRaceRush-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/m/mini-race-rush/v1/",
		Width:   320,
		Height:  497,
		Options: "{}",

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
		Options: "{}",
	},

	{
		Name:    "BubblesShooter-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/b/bubbles-shooter/11-2be95/",
		Width:   522,
		Height:  348,
		Options: "{}",

		KeyWhitelist: []string{
			"Space",
		},
	},

	{
		Name:    "PenguinSkip-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/p/penguin-skip/72b65ff0/",
		Width:   480,
		Height:  320,
		Options: "{}",

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
		Options: "{}",

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
		Options: "{\"car\":0,\"weapon\":0}",

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
		Options: "{\"car\":0,\"weapon\":1}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
			"Space",
		},

		VariantOf: "BurninRubber-v0",
	},

	{
		Name:    "BurninRubber-v2",
		BaseURL: "http://games.cdn.famobi.com/html5games/b/burnin-rubber/11-bc4df/",
		Width:   320,
		Height:  480,
		Options: "{\"car\":1,\"weapon\":0}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
			"Space",
		},

		VariantOf: "BurninRubber-v0",
	},

	{
		Name:    "BurninRubber-v3",
		BaseURL: "http://games.cdn.famobi.com/html5games/b/burnin-rubber/11-bc4df/",
		Width:   320,
		Height:  480,
		Options: "{\"car\":1,\"weapon\":1}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
			"Space",
		},

		VariantOf: "BurninRubber-v0",
	},

	{
		Name:    "OnetConnectClassic-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/o/onet-connect-classic/v160/",
		Width:   480,
		Height:  320,
		Options: "{}",
	},

	{
		Name:    "DontCrash-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/d/dont-crash/v2/",
		Width:   480,
		Height:  320,
		Options: "{}",
	},

	{
		Name:    "SmartyBubbles-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/s/smarty-bubbles/v030/",
		Width:   320,
		Height:  480,
		Options: "{}",
	},

	{
		Name:    "StreetPursuit-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/s/street-pursuit/7-8cd52b/",
		Width:   320,
		Height:  497,
		Options: "{}",

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
		Options: "{}",
	},

	{
		Name:    "GoldMinerTom-v1",
		BaseURL: "http://games.cdn.famobi.com/html5games/g/gold-miner-tom/v120/",
		Width:   320,
		Height:  480,
		Options: "{}",

		VariantOf: "GoldMinerTom-v0",
	},

	{
		Name:    "UfoRun-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/u/ufo-run/6f9ac9a0/",
		Width:   480,
		Height:  321,
		Options: "{}",

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
		Options: "{}",

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
		Options: "{}",
	},

	{
		Name:    "RedHead-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/r/red-head/v1/",
		Width:   480,
		Height:  320,
		Options: "{}",

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
		Options: "{\"level\":1}",
	},

	{
		Name:    "NutRush-v1",
		BaseURL: "http://games.cdn.famobi.com/html5games/n/nut-rush/v030/",
		Width:   480,
		Height:  320,
		Options: "{\"level\":2}",

		VariantOf: "NutRush-v0",
	},

	{
		Name:    "NutRush-v2",
		BaseURL: "http://games.cdn.famobi.com/html5games/n/nut-rush/v030/",
		Width:   480,
		Height:  320,
		Options: "{\"level\":3}",

		VariantOf: "NutRush-v0",
	},

	{
		Name:    "KibaKumbaShadowRun-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/k/kk-shadow-run/v3/",
		Width:   560,
		Height:  320,
		Options: "{\"character\":0}",
	},

	{
		Name:    "KibaKumbaShadowRun-v1",
		BaseURL: "http://games.cdn.famobi.com/html5games/k/kk-shadow-run/v3/",
		Width:   560,
		Height:  320,
		Options: "{\"character\":1}",

		VariantOf: "KibaKumbaShadowRun-v0",
	},

	{
		Name:    "Knightower-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/k/knightower/3-c82d31/",
		Width:   320,
		Height:  480,
		Options: "{}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
		},
	},

	{
		Name:    "BurgerMaker-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/b/burger-maker/v1/",
		Width:   560,
		Height:  420,
		Options: "{}",
	},

	{
		Name:    "PaperPlaneFlight-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/p/paper-plane-flight/v3/",
		Width:   320,
		Height:  480,
		Options: "{}",
	},

	{
		Name:    "PaperPlaneFlight-v1",
		BaseURL: "http://games.cdn.famobi.com/html5games/p/paper-plane-flight/v3/",
		Width:   320,
		Height:  480,
		Options: "{\"planeType\":\"yellow\"}",

		VariantOf: "PaperPlaneFlight-v0",
	},

	{
		Name:    "PaperPlaneFlight-v2",
		BaseURL: "http://games.cdn.famobi.com/html5games/p/paper-plane-flight/v3/",
		Width:   320,
		Height:  480,
		Options: "{\"planeType\":\"violetSecond\"}",

		VariantOf: "PaperPlaneFlight-v0",
	},

	{
		Name:    "OfficeLove-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/o/office-love/v1/",
		Width:   320,
		Height:  510,
		Options: "{}",
	},

	{
		Name:    "RabbitPunch-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/r/rabbit-punch/v1/",
		Width:   480,
		Height:  270,
		Options: "{}",
	},

	{
		Name:    "KumbaKarate-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/k/kumba-karate/v2/",
		Width:   480,
		Height:  380,
		Options: "{}",

		KeyWhitelist: []string{
			"ArrowUp",
			"ArrowDown",
			"ArrowRight",
			"ArrowLeft",
		},
	},

	{
		Name:    "Twins-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/t/twins/v2/",
		Width:   320,
		Height:  480,
		Options: "{}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
		},
	},

	{
		Name:    "Babel-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/b/babel/v040/",
		Width:   270,
		Height:  480,
		Options: "{}",
	},

	{
		Name:    "SpringPanda-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/s/spring-panda/v040/",
		Width:   270,
		Height:  480,
		Options: "{}",
	},

	{
		Name:    "Traffic-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/t/traffic/v050/",
		Width:   320,
		Height:  486,
		Options: "{}",
	},

	{
		Name:    "CartoonFlight-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/c/cartoon-flight/v3/",
		Width:   321,
		Height:  480,
		Options: "{\"scoreMode\":\"distance\"}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
		},
	},

	{
		Name:    "CartoonFlight-v1",
		BaseURL: "http://games.cdn.famobi.com/html5games/c/cartoon-flight/v3/",
		Width:   321,
		Height:  480,
		Options: "{\"scoreMode\":\"stars\"}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
		},

		VariantOf: "CartoonFlight-v0",
	},

	{
		Name:    "GoldMine-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/g/gold-mine/v130/",
		Width:   480,
		Height:  320,
		Options: "{}",
	},

	{
		Name:    "Basketball-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/b/basketball/v070/",
		Width:   320,
		Height:  480,
		Options: "{}",
	},

	{
		Name:    "ProtectThePlanet-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/p/protect-the-planet/v2/",
		Width:   1024,
		Height:  640,
		Options: "{}",
	},

	{
		Name:    "FruitBreak-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/f/fruit-break/v150/",
		Width:   480,
		Height:  270,
		Options: "{}",
	},

	{
		Name:    "PizzaNinja3-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/p/pizza-ninja-3/v090/",
		Width:   480,
		Height:  288,
		Options: "{}",
	},

	{
		Name:    "TimberMan-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/t/timber-man/v1/",
		Width:   320,
		Height:  480,
		Options: "{}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
		},
	},

	{
		Name:    "Shards-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/s/shards/v220/",
		Width:   320,
		Height:  512,
		Options: "{\"level\":0}",
	},

	{
		Name:    "Shards-v1",
		BaseURL: "http://games.cdn.famobi.com/html5games/s/shards/v220/",
		Width:   320,
		Height:  512,
		Options: "{\"level\":1}",

		VariantOf: "Shards-v0",
	},

	{
		Name:    "Shards-v2",
		BaseURL: "http://games.cdn.famobi.com/html5games/s/shards/v220/",
		Width:   320,
		Height:  512,
		Options: "{\"level\":2}",

		VariantOf: "Shards-v0",
	},

	{
		Name:    "Shards-v3",
		BaseURL: "http://games.cdn.famobi.com/html5games/s/shards/v220/",
		Width:   320,
		Height:  512,
		Options: "{\"level\":3}",

		VariantOf: "Shards-v0",
	},

	{
		Name:    "Zop-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/z/zop/v070/",
		Width:   320,
		Height:  480,
		Options: "{}",
	},

	{
		Name:    "Zop-v1",
		BaseURL: "http://games.cdn.famobi.com/html5games/z/zop/v070/",
		Width:   480,
		Height:  320,
		Options: "{}",

		VariantOf: "Zop-v0",
	},

	{
		Name:    "DotsMania-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/d/dots-mania/19-4aace/",
		Width:   320,
		Height:  480,
		Options: "{\"bonus\":false,\"mode\":\"time\"}",
	},

	{
		Name:    "DotsMania-v1",
		BaseURL: "http://games.cdn.famobi.com/html5games/d/dots-mania/19-4aace/",
		Width:   320,
		Height:  480,
		Options: "{\"bonus\":false,\"mode\":\"moves\"}",

		VariantOf: "DotsMania-v0",
	},

	{
		Name:    "DotsMania-v2",
		BaseURL: "http://games.cdn.famobi.com/html5games/d/dots-mania/19-4aace/",
		Width:   320,
		Height:  480,
		Options: "{\"bonus\":true,\"mode\":\"time\"}",

		VariantOf: "DotsMania-v0",
	},

	{
		Name:    "DotsMania-v3",
		BaseURL: "http://games.cdn.famobi.com/html5games/d/dots-mania/19-4aace/",
		Width:   320,
		Height:  480,
		Options: "{\"bonus\":true,\"mode\":\"moves\"}",

		VariantOf: "DotsMania-v0",
	},

	{
		Name:    "FitItQuick-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/f/fit-it-quick/v010/",
		Width:   430,
		Height:  320,
		Options: "{\"level\":1}",
	},

	{
		Name:    "FitItQuick-v1",
		BaseURL: "http://games.cdn.famobi.com/html5games/f/fit-it-quick/v010/",
		Width:   430,
		Height:  320,
		Options: "{\"level\":2}",

		VariantOf: "FitItQuick-v0",
	},

	{
		Name:    "FitItQuick-v2",
		BaseURL: "http://games.cdn.famobi.com/html5games/f/fit-it-quick/v010/",
		Width:   430,
		Height:  320,
		Options: "{\"level\":3}",

		VariantOf: "FitItQuick-v0",
	},

	{
		Name:    "FitItQuick-v3",
		BaseURL: "http://games.cdn.famobi.com/html5games/f/fit-it-quick/v010/",
		Width:   430,
		Height:  320,
		Options: "{\"level\":4}",

		VariantOf: "FitItQuick-v0",
	},

	{
		Name:    "FitItQuick-v4",
		BaseURL: "http://games.cdn.famobi.com/html5games/f/fit-it-quick/v010/",
		Width:   430,
		Height:  320,
		Options: "{\"level\":5}",

		VariantOf: "FitItQuick-v0",
	},

	{
		Name:    "StackTowerClassic-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/s/stack-tower-classic/v110/",
		Width:   320,
		Height:  480,
		Options: "{}",
	},

	{
		Name:    "StackTowerClassic-v1",
		BaseURL: "http://games.cdn.famobi.com/html5games/s/stack-tower-classic/v110/",
		Width:   390,
		Height:  390,
		Options: "{}",

		VariantOf: "StackTowerClassic-v0",
	},

	{
		Name:    "TowerMania-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/t/tower-mania/v3/",
		Width:   320,
		Height:  455,
		Options: "{\"level\":0}",
	},

	{
		Name:    "TowerMania-v1",
		BaseURL: "http://games.cdn.famobi.com/html5games/t/tower-mania/v3/",
		Width:   320,
		Height:  455,
		Options: "{\"level\":1}",

		VariantOf: "TowerMania-v0",
	},

	{
		Name:    "TowerMania-v2",
		BaseURL: "http://games.cdn.famobi.com/html5games/t/tower-mania/v3/",
		Width:   320,
		Height:  455,
		Options: "{\"level\":2}",

		VariantOf: "TowerMania-v0",
	},

	{
		Name:    "TowerMania-v3",
		BaseURL: "http://games.cdn.famobi.com/html5games/t/tower-mania/v3/",
		Width:   320,
		Height:  455,
		Options: "{\"level\":3}",

		VariantOf: "TowerMania-v0",
	},

	{
		Name:    "GroovySki-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/g/groovy-ski/8-52acd9/",
		Width:   320,
		Height:  544,
		Options: "{}",
	},

	{
		Name:    "ColorCircles-v0",
		BaseURL: "http://games.cdn.famobi.com/html5games/c/color-circles/v070/",
		Width:   320,
		Height:  480,
		Options: "{}",

		KeyWhitelist: []string{
			"Space",
		},
	},

	{
		Name:    "ColorCircles-v1",
		BaseURL: "http://games.cdn.famobi.com/html5games/c/color-circles/v070/",
		Width:   390,
		Height:  390,
		Options: "{}",

		KeyWhitelist: []string{
			"Space",
		},

		VariantOf: "ColorCircles-v0",
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
