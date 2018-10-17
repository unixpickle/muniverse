// This file was auto-generated.

package muniverse

// An EnvSpec contains meta-data about an environment.
type EnvSpec struct {
	Name      string `bson:"Name"`
	BaseURL   string `bson:"BaseURL"`
	Width     int    `bson:"Width"`
	Height    int    `bson:"Height"`
	AllCanvas bool   `bson:"AllCanvas"`
	Options   string `bson:"Options"`

	KeyWhitelist []string `bson:"KeyWhitelist,omitempty"`

	VariantOf string `bson:"VariantOf"`
}

var EnvSpecs = []*EnvSpec{
	{
		Name:      "KatanaFruits-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/k/katana-fruits/v1/",
		Width:     320,
		Height:    427,
		AllCanvas: true,
		Options:   "{}",
	},

	{
		Name:      "MiniRaceRush-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/m/mini-race-rush/v1/",
		Width:     320,
		Height:    497,
		AllCanvas: true,
		Options:   "{}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
		},
	},

	{
		Name:      "Jewelish-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/j/jewelish/10-4bc4e/",
		Width:     320,
		Height:    433,
		AllCanvas: false,
		Options:   "{}",
	},

	{
		Name:      "BubblesShooter-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/b/bubbles-shooter/11-2be95/",
		Width:     522,
		Height:    348,
		AllCanvas: false,
		Options:   "{}",

		KeyWhitelist: []string{
			"Space",
		},
	},

	{
		Name:      "PenguinSkip-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/p/penguin-skip/72b65ff0/",
		Width:     480,
		Height:    320,
		AllCanvas: true,
		Options:   "{}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
		},
	},

	{
		Name:      "MinimalDots-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/m/minimal-dots/v2/",
		Width:     320,
		Height:    480,
		AllCanvas: true,
		Options:   "{}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
		},
	},

	{
		Name:      "BurninRubber-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/b/burnin-rubber/11-bc4df/",
		Width:     320,
		Height:    480,
		AllCanvas: true,
		Options:   "{\"car\":0,\"weapon\":0}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
			"Space",
		},
	},

	{
		Name:      "BurninRubber-v1",
		BaseURL:   "http://games.cdn.famobi.com/html5games/b/burnin-rubber/11-bc4df/",
		Width:     320,
		Height:    480,
		AllCanvas: true,
		Options:   "{\"car\":0,\"weapon\":1}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
			"Space",
		},

		VariantOf: "BurninRubber-v0",
	},

	{
		Name:      "BurninRubber-v2",
		BaseURL:   "http://games.cdn.famobi.com/html5games/b/burnin-rubber/11-bc4df/",
		Width:     320,
		Height:    480,
		AllCanvas: true,
		Options:   "{\"car\":1,\"weapon\":0}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
			"Space",
		},

		VariantOf: "BurninRubber-v0",
	},

	{
		Name:      "BurninRubber-v3",
		BaseURL:   "http://games.cdn.famobi.com/html5games/b/burnin-rubber/11-bc4df/",
		Width:     320,
		Height:    480,
		AllCanvas: true,
		Options:   "{\"car\":1,\"weapon\":1}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
			"Space",
		},

		VariantOf: "BurninRubber-v0",
	},

	{
		Name:      "OnetConnectClassic-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/o/onet-connect-classic/v160/",
		Width:     480,
		Height:    320,
		AllCanvas: true,
		Options:   "{}",
	},

	{
		Name:      "DontCrash-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/d/dont-crash/v2/",
		Width:     480,
		Height:    320,
		AllCanvas: true,
		Options:   "{}",
	},

	{
		Name:      "SmartyBubbles-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/s/smarty-bubbles/v030/",
		Width:     320,
		Height:    480,
		AllCanvas: false,
		Options:   "{}",
	},

	{
		Name:      "StreetPursuit-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/s/street-pursuit/7-8cd52b/",
		Width:     320,
		Height:    497,
		AllCanvas: true,
		Options:   "{}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
		},
	},

	{
		Name:      "GoldMinerTom-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/g/gold-miner-tom/v120/",
		Width:     480,
		Height:    320,
		AllCanvas: true,
		Options:   "{}",
	},

	{
		Name:      "GoldMinerTom-v1",
		BaseURL:   "http://games.cdn.famobi.com/html5games/g/gold-miner-tom/v120/",
		Width:     320,
		Height:    480,
		AllCanvas: true,
		Options:   "{}",

		VariantOf: "GoldMinerTom-v0",
	},

	{
		Name:      "UfoRun-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/u/ufo-run/6f9ac9a0/",
		Width:     480,
		Height:    321,
		AllCanvas: false,
		Options:   "{}",

		KeyWhitelist: []string{
			"Space",
			"ArrowUp",
		},
	},

	{
		Name:      "BirdyRush-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/b/birdy-rush/v040/",
		Width:     320,
		Height:    533,
		AllCanvas: true,
		Options:   "{}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
		},
	},

	{
		Name:      "HexBlitz-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/h/hex-blitz/v270/",
		Width:     480,
		Height:    320,
		AllCanvas: true,
		Options:   "{}",
	},

	{
		Name:      "RedHead-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/r/red-head/v1/",
		Width:     480,
		Height:    320,
		AllCanvas: true,
		Options:   "{}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
		},
	},

	{
		Name:      "NutRush-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/n/nut-rush/v030/",
		Width:     480,
		Height:    320,
		AllCanvas: true,
		Options:   "{\"level\":1}",
	},

	{
		Name:      "NutRush-v1",
		BaseURL:   "http://games.cdn.famobi.com/html5games/n/nut-rush/v030/",
		Width:     480,
		Height:    320,
		AllCanvas: true,
		Options:   "{\"level\":2}",

		VariantOf: "NutRush-v0",
	},

	{
		Name:      "NutRush-v2",
		BaseURL:   "http://games.cdn.famobi.com/html5games/n/nut-rush/v030/",
		Width:     480,
		Height:    320,
		AllCanvas: true,
		Options:   "{\"level\":3}",

		VariantOf: "NutRush-v0",
	},

	{
		Name:      "KibaKumbaShadowRun-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/k/kk-shadow-run/v3/",
		Width:     560,
		Height:    320,
		AllCanvas: true,
		Options:   "{\"character\":0}",
	},

	{
		Name:      "KibaKumbaShadowRun-v1",
		BaseURL:   "http://games.cdn.famobi.com/html5games/k/kk-shadow-run/v3/",
		Width:     560,
		Height:    320,
		AllCanvas: true,
		Options:   "{\"character\":1}",

		VariantOf: "KibaKumbaShadowRun-v0",
	},

	{
		Name:      "Knightower-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/k/knightower/3-c82d31/",
		Width:     320,
		Height:    480,
		AllCanvas: true,
		Options:   "{}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
		},
	},

	{
		Name:      "BurgerMaker-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/b/burger-maker/v1/",
		Width:     560,
		Height:    420,
		AllCanvas: false,
		Options:   "{}",
	},

	{
		Name:      "PaperPlaneFlight-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/p/paper-plane-flight/v3/",
		Width:     320,
		Height:    480,
		AllCanvas: true,
		Options:   "{}",
	},

	{
		Name:      "PaperPlaneFlight-v1",
		BaseURL:   "http://games.cdn.famobi.com/html5games/p/paper-plane-flight/v3/",
		Width:     320,
		Height:    480,
		AllCanvas: true,
		Options:   "{\"planeType\":\"yellow\"}",

		VariantOf: "PaperPlaneFlight-v0",
	},

	{
		Name:      "PaperPlaneFlight-v2",
		BaseURL:   "http://games.cdn.famobi.com/html5games/p/paper-plane-flight/v3/",
		Width:     320,
		Height:    480,
		AllCanvas: true,
		Options:   "{\"planeType\":\"violetSecond\"}",

		VariantOf: "PaperPlaneFlight-v0",
	},

	{
		Name:      "OfficeLove-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/o/office-love/v1/",
		Width:     320,
		Height:    510,
		AllCanvas: false,
		Options:   "{}",
	},

	{
		Name:      "RabbitPunch-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/r/rabbit-punch/v1/",
		Width:     480,
		Height:    270,
		AllCanvas: true,
		Options:   "{}",
	},

	{
		Name:      "KumbaKarate-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/k/kumba-karate/v2/",
		Width:     480,
		Height:    380,
		AllCanvas: true,
		Options:   "{}",

		KeyWhitelist: []string{
			"ArrowUp",
			"ArrowDown",
			"ArrowRight",
			"ArrowLeft",
		},
	},

	{
		Name:      "Twins-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/t/twins/v2/",
		Width:     320,
		Height:    480,
		AllCanvas: true,
		Options:   "{}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
		},
	},

	{
		Name:      "Babel-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/b/babel/v040/",
		Width:     270,
		Height:    480,
		AllCanvas: true,
		Options:   "{}",
	},

	{
		Name:      "SpringPanda-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/s/spring-panda/v040/",
		Width:     270,
		Height:    480,
		AllCanvas: true,
		Options:   "{}",
	},

	{
		Name:      "Traffic-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/t/traffic/v050/",
		Width:     320,
		Height:    486,
		AllCanvas: true,
		Options:   "{}",
	},

	{
		Name:      "CartoonFlight-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/c/cartoon-flight/v3/",
		Width:     321,
		Height:    480,
		AllCanvas: false,
		Options:   "{\"scoreMode\":\"distance\"}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
		},
	},

	{
		Name:      "CartoonFlight-v1",
		BaseURL:   "http://games.cdn.famobi.com/html5games/c/cartoon-flight/v3/",
		Width:     321,
		Height:    480,
		AllCanvas: false,
		Options:   "{\"scoreMode\":\"stars\"}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
		},

		VariantOf: "CartoonFlight-v0",
	},

	{
		Name:      "GoldMine-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/g/gold-mine/v130/",
		Width:     480,
		Height:    320,
		AllCanvas: true,
		Options:   "{}",
	},

	{
		Name:      "Basketball-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/b/basketball/v070/",
		Width:     320,
		Height:    480,
		AllCanvas: true,
		Options:   "{}",
	},

	{
		Name:      "ProtectThePlanet-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/p/protect-the-planet/v2/",
		Width:     1024,
		Height:    640,
		AllCanvas: false,
		Options:   "{}",
	},

	{
		Name:      "FruitBreak-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/f/fruit-break/v150/",
		Width:     480,
		Height:    270,
		AllCanvas: true,
		Options:   "{}",
	},

	{
		Name:      "PizzaNinja3-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/p/pizza-ninja-3/v090/",
		Width:     480,
		Height:    288,
		AllCanvas: false,
		Options:   "{}",
	},

	{
		Name:      "TimberMan-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/t/timber-man/v1/",
		Width:     320,
		Height:    480,
		AllCanvas: true,
		Options:   "{}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
		},
	},

	{
		Name:      "Shards-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/s/shards/v220/",
		Width:     320,
		Height:    512,
		AllCanvas: false,
		Options:   "{\"level\":0}",
	},

	{
		Name:      "Shards-v1",
		BaseURL:   "http://games.cdn.famobi.com/html5games/s/shards/v220/",
		Width:     320,
		Height:    512,
		AllCanvas: false,
		Options:   "{\"level\":1}",

		VariantOf: "Shards-v0",
	},

	{
		Name:      "Shards-v2",
		BaseURL:   "http://games.cdn.famobi.com/html5games/s/shards/v220/",
		Width:     320,
		Height:    512,
		AllCanvas: false,
		Options:   "{\"level\":2}",

		VariantOf: "Shards-v0",
	},

	{
		Name:      "Shards-v3",
		BaseURL:   "http://games.cdn.famobi.com/html5games/s/shards/v220/",
		Width:     320,
		Height:    512,
		AllCanvas: false,
		Options:   "{\"level\":3}",

		VariantOf: "Shards-v0",
	},

	{
		Name:      "Zop-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/z/zop/v070/",
		Width:     320,
		Height:    480,
		AllCanvas: false,
		Options:   "{}",
	},

	{
		Name:      "Zop-v1",
		BaseURL:   "http://games.cdn.famobi.com/html5games/z/zop/v070/",
		Width:     480,
		Height:    320,
		AllCanvas: false,
		Options:   "{}",

		VariantOf: "Zop-v0",
	},

	{
		Name:      "DotsMania-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/d/dots-mania/19-4aace/",
		Width:     320,
		Height:    480,
		AllCanvas: false,
		Options:   "{\"bonus\":false,\"mode\":\"time\"}",
	},

	{
		Name:      "DotsMania-v1",
		BaseURL:   "http://games.cdn.famobi.com/html5games/d/dots-mania/19-4aace/",
		Width:     320,
		Height:    480,
		AllCanvas: false,
		Options:   "{\"bonus\":false,\"mode\":\"moves\"}",

		VariantOf: "DotsMania-v0",
	},

	{
		Name:      "DotsMania-v2",
		BaseURL:   "http://games.cdn.famobi.com/html5games/d/dots-mania/19-4aace/",
		Width:     320,
		Height:    480,
		AllCanvas: false,
		Options:   "{\"bonus\":true,\"mode\":\"time\"}",

		VariantOf: "DotsMania-v0",
	},

	{
		Name:      "DotsMania-v3",
		BaseURL:   "http://games.cdn.famobi.com/html5games/d/dots-mania/19-4aace/",
		Width:     320,
		Height:    480,
		AllCanvas: false,
		Options:   "{\"bonus\":true,\"mode\":\"moves\"}",

		VariantOf: "DotsMania-v0",
	},

	{
		Name:      "FitItQuick-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/f/fit-it-quick/v010/",
		Width:     430,
		Height:    320,
		AllCanvas: true,
		Options:   "{\"level\":1}",
	},

	{
		Name:      "FitItQuick-v1",
		BaseURL:   "http://games.cdn.famobi.com/html5games/f/fit-it-quick/v010/",
		Width:     430,
		Height:    320,
		AllCanvas: true,
		Options:   "{\"level\":2}",

		VariantOf: "FitItQuick-v0",
	},

	{
		Name:      "FitItQuick-v2",
		BaseURL:   "http://games.cdn.famobi.com/html5games/f/fit-it-quick/v010/",
		Width:     430,
		Height:    320,
		AllCanvas: true,
		Options:   "{\"level\":3}",

		VariantOf: "FitItQuick-v0",
	},

	{
		Name:      "FitItQuick-v3",
		BaseURL:   "http://games.cdn.famobi.com/html5games/f/fit-it-quick/v010/",
		Width:     430,
		Height:    320,
		AllCanvas: true,
		Options:   "{\"level\":4}",

		VariantOf: "FitItQuick-v0",
	},

	{
		Name:      "FitItQuick-v4",
		BaseURL:   "http://games.cdn.famobi.com/html5games/f/fit-it-quick/v010/",
		Width:     430,
		Height:    320,
		AllCanvas: true,
		Options:   "{\"level\":5}",

		VariantOf: "FitItQuick-v0",
	},

	{
		Name:      "StackTowerClassic-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/s/stack-tower-classic/v110/",
		Width:     320,
		Height:    480,
		AllCanvas: true,
		Options:   "{}",
	},

	{
		Name:      "StackTowerClassic-v1",
		BaseURL:   "http://games.cdn.famobi.com/html5games/s/stack-tower-classic/v110/",
		Width:     390,
		Height:    390,
		AllCanvas: true,
		Options:   "{}",

		VariantOf: "StackTowerClassic-v0",
	},

	{
		Name:      "TowerMania-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/t/tower-mania/v3/",
		Width:     320,
		Height:    455,
		AllCanvas: false,
		Options:   "{\"level\":0}",
	},

	{
		Name:      "TowerMania-v1",
		BaseURL:   "http://games.cdn.famobi.com/html5games/t/tower-mania/v3/",
		Width:     320,
		Height:    455,
		AllCanvas: false,
		Options:   "{\"level\":1}",

		VariantOf: "TowerMania-v0",
	},

	{
		Name:      "TowerMania-v2",
		BaseURL:   "http://games.cdn.famobi.com/html5games/t/tower-mania/v3/",
		Width:     320,
		Height:    455,
		AllCanvas: false,
		Options:   "{\"level\":2}",

		VariantOf: "TowerMania-v0",
	},

	{
		Name:      "TowerMania-v3",
		BaseURL:   "http://games.cdn.famobi.com/html5games/t/tower-mania/v3/",
		Width:     320,
		Height:    455,
		AllCanvas: false,
		Options:   "{\"level\":3}",

		VariantOf: "TowerMania-v0",
	},

	{
		Name:      "GroovySki-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/g/groovy-ski/8-52acd9/",
		Width:     320,
		Height:    544,
		AllCanvas: true,
		Options:   "{}",
	},

	{
		Name:      "ColorCircles-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/c/color-circles/v070/",
		Width:     320,
		Height:    480,
		AllCanvas: true,
		Options:   "{}",

		KeyWhitelist: []string{
			"Space",
		},
	},

	{
		Name:      "ColorCircles-v1",
		BaseURL:   "http://games.cdn.famobi.com/html5games/c/color-circles/v070/",
		Width:     390,
		Height:    390,
		AllCanvas: true,
		Options:   "{}",

		KeyWhitelist: []string{
			"Space",
		},

		VariantOf: "ColorCircles-v0",
	},

	{
		Name:      "SnowSmasher-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/s/snow-smasher/v4/",
		Width:     320,
		Height:    550,
		AllCanvas: true,
		Options:   "{}",
	},

	{
		Name:      "Multisquare-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/m/multisquare/17-687f3/",
		Width:     320,
		Height:    487,
		AllCanvas: true,
		Options:   "{}",
	},

	{
		Name:      "Colorpop-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/c/colorpop/v120/",
		Width:     320,
		Height:    480,
		AllCanvas: true,
		Options:   "{}",
	},

	{
		Name:      "FruitPulp-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/f/fruit-pulp/v080/",
		Width:     320,
		Height:    480,
		AllCanvas: false,
		Options:   "{\"level\":1}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
			"ArrowUp",
			"ArrowDown",
		},
	},

	{
		Name:      "FruitPulp-v1",
		BaseURL:   "http://games.cdn.famobi.com/html5games/f/fruit-pulp/v080/",
		Width:     320,
		Height:    480,
		AllCanvas: false,
		Options:   "{\"level\":2}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
			"ArrowUp",
			"ArrowDown",
		},

		VariantOf: "FruitPulp-v0",
	},

	{
		Name:      "FruitPulp-v2",
		BaseURL:   "http://games.cdn.famobi.com/html5games/f/fruit-pulp/v080/",
		Width:     320,
		Height:    480,
		AllCanvas: false,
		Options:   "{\"level\":3}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
			"ArrowUp",
			"ArrowDown",
		},

		VariantOf: "FruitPulp-v0",
	},

	{
		Name:      "FruitPulp-v3",
		BaseURL:   "http://games.cdn.famobi.com/html5games/f/fruit-pulp/v080/",
		Width:     320,
		Height:    480,
		AllCanvas: false,
		Options:   "{\"level\":4}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
			"ArrowUp",
			"ArrowDown",
		},

		VariantOf: "FruitPulp-v0",
	},

	{
		Name:      "FruitPulp-v4",
		BaseURL:   "http://games.cdn.famobi.com/html5games/f/fruit-pulp/v080/",
		Width:     320,
		Height:    480,
		AllCanvas: false,
		Options:   "{\"level\":5}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
			"ArrowUp",
			"ArrowDown",
		},

		VariantOf: "FruitPulp-v0",
	},

	{
		Name:      "JewelExplode-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/j/jewel-explode/v080/",
		Width:     320,
		Height:    533,
		AllCanvas: false,
		Options:   "{\"level\":0}",
	},

	{
		Name:      "JewelExplode-v1",
		BaseURL:   "http://games.cdn.famobi.com/html5games/j/jewel-explode/v080/",
		Width:     320,
		Height:    533,
		AllCanvas: false,
		Options:   "{\"level\":1}",

		VariantOf: "JewelExplode-v0",
	},

	{
		Name:      "JewelExplode-v2",
		BaseURL:   "http://games.cdn.famobi.com/html5games/j/jewel-explode/v080/",
		Width:     320,
		Height:    533,
		AllCanvas: false,
		Options:   "{\"level\":2}",

		VariantOf: "JewelExplode-v0",
	},

	{
		Name:      "JewelExplode-v3",
		BaseURL:   "http://games.cdn.famobi.com/html5games/j/jewel-explode/v080/",
		Width:     320,
		Height:    533,
		AllCanvas: false,
		Options:   "{\"level\":3}",

		VariantOf: "JewelExplode-v0",
	},

	{
		Name:      "JewelExplode-v4",
		BaseURL:   "http://games.cdn.famobi.com/html5games/j/jewel-explode/v080/",
		Width:     320,
		Height:    533,
		AllCanvas: false,
		Options:   "{\"level\":4}",

		VariantOf: "JewelExplode-v0",
	},

	{
		Name:      "ZooPinball-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/z/zoo-pinball/5-1454c9/",
		Width:     320,
		Height:    544,
		AllCanvas: true,
		Options:   "{}",

		KeyWhitelist: []string{
			"Space",
			"ArrowLeft",
			"ArrowRight",
		},
	},

	{
		Name:      "BubbleHamsters-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/b/bubble-hamsters/5-1bbcb7/",
		Width:     320,
		Height:    460,
		AllCanvas: true,
		Options:   "{}",
	},

	{
		Name:      "FlyingSchool-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/f/flying-school/8-066c1e/",
		Width:     480,
		Height:    320,
		AllCanvas: true,
		Options:   "{\"scoreMode\":\"distance\"}",
	},

	{
		Name:      "FlyingSchool-v1",
		BaseURL:   "http://games.cdn.famobi.com/html5games/f/flying-school/8-066c1e/",
		Width:     480,
		Height:    320,
		AllCanvas: true,
		Options:   "{\"scoreMode\":\"stars\"}",

		VariantOf: "FlyingSchool-v0",
	},

	{
		Name:      "CarCrossing-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/c/car-crossing/v130/",
		Width:     320,
		Height:    570,
		AllCanvas: false,
		Options:   "{}",
	},

	{
		Name:      "DoggyDive-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/d/doggy-dive/5-671a84/",
		Width:     320,
		Height:    480,
		AllCanvas: true,
		Options:   "{}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
		},
	},

	{
		Name:      "StickFreak-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/s/stick-freak/89f33144/",
		Width:     320,
		Height:    533,
		AllCanvas: false,
		Options:   "{}",
	},

	{
		Name:      "Lectro-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/l/lectro/v2/",
		Width:     320,
		Height:    480,
		AllCanvas: true,
		Options:   "{}",
	},

	{
		Name:      "Lectro-v1",
		BaseURL:   "http://games.cdn.famobi.com/html5games/l/lectro/v2/",
		Width:     390,
		Height:    390,
		AllCanvas: true,
		Options:   "{}",

		VariantOf: "Lectro-v0",
	},

	{
		Name:      "StreetBallStar-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/s/street-ball-star/7-7489fa/",
		Width:     320,
		Height:    480,
		AllCanvas: true,
		Options:   "{\"mode\":\"time\"}",
	},

	{
		Name:      "StreetBallStar-v1",
		BaseURL:   "http://games.cdn.famobi.com/html5games/s/street-ball-star/7-7489fa/",
		Width:     480,
		Height:    360,
		AllCanvas: true,
		Options:   "{\"mode\":\"time\"}",

		VariantOf: "StreetBallStar-v0",
	},

	{
		Name:      "StreetBallStar-v2",
		BaseURL:   "http://games.cdn.famobi.com/html5games/s/street-ball-star/7-7489fa/",
		Width:     320,
		Height:    480,
		AllCanvas: true,
		Options:   "{\"mode\":\"challenge\"}",

		VariantOf: "StreetBallStar-v0",
	},

	{
		Name:      "StreetBallStar-v3",
		BaseURL:   "http://games.cdn.famobi.com/html5games/s/street-ball-star/7-7489fa/",
		Width:     480,
		Height:    360,
		AllCanvas: true,
		Options:   "{\"mode\":\"challenge\"}",

		VariantOf: "StreetBallStar-v0",
	},

	{
		Name:      "BaseballPro-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/b/baseball-pro/v050/",
		Width:     320,
		Height:    480,
		AllCanvas: false,
		Options:   "{}",
	},

	{
		Name:      "SoccerGirl-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/s/soccer-girl/v050/",
		Width:     320,
		Height:    646,
		AllCanvas: true,
		Options:   "{}",
	},

	{
		Name:      "SoccerGirl-v1",
		BaseURL:   "http://games.cdn.famobi.com/html5games/s/soccer-girl/v050/",
		Width:     426,
		Height:    320,
		AllCanvas: true,
		Options:   "{}",

		VariantOf: "SoccerGirl-v0",
	},

	{
		Name:      "TRex-v0",
		BaseURL:   "https://chromedino.com/",
		Width:     600,
		Height:    150,
		AllCanvas: false,
		Options:   "{}",

		KeyWhitelist: []string{
			"ArrowUp",
			"ArrowDown",
			"Space",
		},
	},

	{
		Name:      "HopDontStop-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/h/hop-dont-stop/v040/",
		Width:     320,
		Height:    498,
		AllCanvas: true,
		Options:   "{\"scoreMode\":\"distance\"}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
			"ArrowUp",
			"ArrowDown",
		},
	},

	{
		Name:      "HopDontStop-v1",
		BaseURL:   "http://games.cdn.famobi.com/html5games/h/hop-dont-stop/v040/",
		Width:     320,
		Height:    498,
		AllCanvas: true,
		Options:   "{\"scoreMode\":\"totalGems\"}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
			"ArrowUp",
			"ArrowDown",
		},

		VariantOf: "HopDontStop-v0",
	},

	{
		Name:      "FitzColor-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/f/fitz-color/v2/",
		Width:     320,
		Height:    480,
		AllCanvas: true,
		Options:   "{\"difficulty\":\"normal\"}",
	},

	{
		Name:      "FitzColor-v1",
		BaseURL:   "http://games.cdn.famobi.com/html5games/f/fitz-color/v2/",
		Width:     320,
		Height:    480,
		AllCanvas: true,
		Options:   "{\"difficulty\":\"hard\"}",

		VariantOf: "FitzColor-v0",
	},

	{
		Name:      "SushiNinjaDash-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/s/sushi-ninja-dash/v3/",
		Width:     320,
		Height:    470,
		AllCanvas: true,
		Options:   "{}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
			"ArrowUp",
		},
	},

	{
		Name:      "RainbowStarPinball-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/r/rainbow-star-pinball/v2/",
		Width:     320,
		Height:    480,
		AllCanvas: true,
		Options:   "{}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
			"ArrowUp",
			"ArrowDown",
			"KeyW",
			"KeyA",
			"KeyS",
			"KeyD",
		},
	},

	{
		Name:      "100GolfBalls-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/0/100-golf-balls/3-30f6bb/",
		Width:     320,
		Height:    480,
		AllCanvas: false,
		Options:   "{}",
	},

	{
		Name:      "KittyBubbles-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/k/kitty-bubbles/v100/",
		Width:     320,
		Height:    480,
		AllCanvas: false,
		Options:   "{}",
	},

	{
		Name:      "JungleRun-v0",
		BaseURL:   "http://games.cdn.famobi.com/html5games/j/jungle-run/v120/",
		Width:     540,
		Height:    360,
		AllCanvas: false,
		Options:   "{\"player\":\"Kiba\"}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
			"ArrowUp",
			"ArrowDown",
		},
	},

	{
		Name:      "JungleRun-v1",
		BaseURL:   "http://games.cdn.famobi.com/html5games/j/jungle-run/v120/",
		Width:     540,
		Height:    360,
		AllCanvas: false,
		Options:   "{\"player\":\"Kumba\"}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
			"ArrowUp",
			"ArrowDown",
		},

		VariantOf: "JungleRun-v0",
	},

	{
		Name:      "S135-v0",
		BaseURL:   "https://game265272.konggames.com/gamez/0026/5272/live/",
		Width:     320,
		Height:    570,
		AllCanvas: false,
		Options:   "{}",

		KeyWhitelist: []string{
			"KeyZ",
		},
	},

	{
		Name:      "Pixlaser-v0",
		BaseURL:   "https://game266022.konggames.com/gamez/0026/6022/live/",
		Width:     550,
		Height:    550,
		AllCanvas: true,
		Options:   "{\"colors\":2}",

		KeyWhitelist: []string{
			"KeyA",
			"KeyD",
		},
	},

	{
		Name:      "Pixlaser-v1",
		BaseURL:   "https://game266022.konggames.com/gamez/0026/6022/live/",
		Width:     550,
		Height:    550,
		AllCanvas: true,
		Options:   "{\"colors\":3}",

		KeyWhitelist: []string{
			"KeyA",
			"KeyD",
		},

		VariantOf: "Pixlaser-v0",
	},

	{
		Name:      "Pixlaser-v2",
		BaseURL:   "https://game266022.konggames.com/gamez/0026/6022/live/",
		Width:     550,
		Height:    550,
		AllCanvas: true,
		Options:   "{\"colors\":4}",

		KeyWhitelist: []string{
			"KeyA",
			"KeyD",
		},

		VariantOf: "Pixlaser-v0",
	},

	{
		Name:      "DoodleJump-v0",
		BaseURL:   "https://cdn.cloudgames.com/games/doodle-jump-new-en-s-iga-cloud/",
		Width:     320,
		Height:    480,
		AllCanvas: true,
		Options:   "{}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
			"ArrowUp",
		},
	},

	{
		Name:      "Valto-v0",
		BaseURL:   "https://www.yiv.com/games/Valto/",
		Width:     320,
		Height:    480,
		AllCanvas: true,
		Options:   "{}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
		},
	},

	{
		Name:      "CatchTheClown-v0",
		BaseURL:   "http://game205756.konggames.com/gamez/0020/5756/live/",
		Width:     640,
		Height:    512,
		AllCanvas: true,
		Options:   "{\"difficulty\":\"easy\"}",
	},

	{
		Name:      "CatchTheClown-v1",
		BaseURL:   "http://game205756.konggames.com/gamez/0020/5756/live/",
		Width:     640,
		Height:    512,
		AllCanvas: true,
		Options:   "{\"difficulty\":\"medium\"}",

		VariantOf: "CatchTheClown-v0",
	},

	{
		Name:      "CatchTheClown-v2",
		BaseURL:   "http://game205756.konggames.com/gamez/0020/5756/live/",
		Width:     640,
		Height:    512,
		AllCanvas: true,
		Options:   "{\"difficulty\":\"hard\"}",

		VariantOf: "CatchTheClown-v0",
	},

	{
		Name:      "MeatBoyClicker-v0",
		BaseURL:   "http://flashok.ru/files/game/meat-boy-clicker/",
		Width:     320,
		Height:    525,
		AllCanvas: false,
		Options:   "{}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
		},
	},

	{
		Name:      "PopPop-v0",
		BaseURL:   "https://www.yiv.com/games/Pop-Pop/",
		Width:     320,
		Height:    568,
		AllCanvas: true,
		Options:   "{}",
	},

	{
		Name:      "FrontInvaders-v0",
		BaseURL:   "http://end3r.com/games/frontinvaders/",
		Width:     700,
		Height:    400,
		AllCanvas: false,
		Options:   "{\"level\":1}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
			"ArrowUp",
		},
	},

	{
		Name:      "FrontInvaders-v1",
		BaseURL:   "http://end3r.com/games/frontinvaders/",
		Width:     700,
		Height:    400,
		AllCanvas: false,
		Options:   "{\"level\":2}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
			"ArrowUp",
		},

		VariantOf: "FrontInvaders-v0",
	},

	{
		Name:      "FrontInvaders-v2",
		BaseURL:   "http://end3r.com/games/frontinvaders/",
		Width:     700,
		Height:    400,
		AllCanvas: false,
		Options:   "{\"level\":3}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
			"ArrowUp",
		},

		VariantOf: "FrontInvaders-v0",
	},

	{
		Name:      "FrontInvaders-v3",
		BaseURL:   "http://end3r.com/games/frontinvaders/",
		Width:     700,
		Height:    400,
		AllCanvas: false,
		Options:   "{\"level\":4}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
			"ArrowUp",
		},

		VariantOf: "FrontInvaders-v0",
	},

	{
		Name:      "FrontInvaders-v4",
		BaseURL:   "http://end3r.com/games/frontinvaders/",
		Width:     700,
		Height:    400,
		AllCanvas: false,
		Options:   "{\"level\":5}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
			"ArrowUp",
		},

		VariantOf: "FrontInvaders-v0",
	},

	{
		Name:      "FrontInvaders-v5",
		BaseURL:   "http://end3r.com/games/frontinvaders/",
		Width:     700,
		Height:    400,
		AllCanvas: false,
		Options:   "{\"level\":6}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
			"ArrowUp",
		},

		VariantOf: "FrontInvaders-v0",
	},

	{
		Name:      "FrontInvaders-v6",
		BaseURL:   "http://end3r.com/games/frontinvaders/",
		Width:     700,
		Height:    400,
		AllCanvas: false,
		Options:   "{\"level\":7}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
			"ArrowUp",
		},

		VariantOf: "FrontInvaders-v0",
	},

	{
		Name:      "FrontInvaders-v7",
		BaseURL:   "http://end3r.com/games/frontinvaders/",
		Width:     700,
		Height:    400,
		AllCanvas: false,
		Options:   "{\"level\":8}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
			"ArrowUp",
		},

		VariantOf: "FrontInvaders-v0",
	},

	{
		Name:      "FrontInvaders-v8",
		BaseURL:   "http://end3r.com/games/frontinvaders/",
		Width:     700,
		Height:    400,
		AllCanvas: false,
		Options:   "{\"level\":9}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
			"ArrowUp",
		},

		VariantOf: "FrontInvaders-v0",
	},

	{
		Name:      "FrontInvaders-v9",
		BaseURL:   "http://end3r.com/games/frontinvaders/",
		Width:     700,
		Height:    400,
		AllCanvas: false,
		Options:   "{\"level\":10}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
			"ArrowUp",
		},

		VariantOf: "FrontInvaders-v0",
	},

	{
		Name:      "FrontInvaders-v10",
		BaseURL:   "http://end3r.com/games/frontinvaders/",
		Width:     700,
		Height:    400,
		AllCanvas: false,
		Options:   "{\"level\":11}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
			"ArrowUp",
		},

		VariantOf: "FrontInvaders-v0",
	},

	{
		Name:      "FrontInvaders-v11",
		BaseURL:   "http://end3r.com/games/frontinvaders/",
		Width:     700,
		Height:    400,
		AllCanvas: false,
		Options:   "{\"level\":12}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
			"ArrowUp",
		},

		VariantOf: "FrontInvaders-v0",
	},

	{
		Name:      "ClickThemAll-v0",
		BaseURL:   "http://game205682.konggames.com/gamez/0020/5682/live/",
		Width:     480,
		Height:    360,
		AllCanvas: true,
		Options:   "{}",
	},

	{
		Name:      "100BallsOnline-v0",
		BaseURL:   "https://h5.4j.com/games/100-Balls-Online/",
		Width:     320,
		Height:    533,
		AllCanvas: false,
		Options:   "{}",
	},

	{
		Name:      "ChristmasAdventure-v0",
		BaseURL:   "https://www.yiv.com/games/Christmas-Adventure/",
		Width:     533,
		Height:    320,
		AllCanvas: true,
		Options:   "{\"level\":1}",

		KeyWhitelist: []string{
			"KeyA",
			"ArrowRight",
			"ArrowLeft",
			"ArrowUp",
		},
	},

	{
		Name:      "ChirstmasAdventure-v1",
		BaseURL:   "https://www.yiv.com/games/Christmas-Adventure/",
		Width:     533,
		Height:    320,
		AllCanvas: true,
		Options:   "{\"level\":2}",

		KeyWhitelist: []string{
			"KeyA",
			"ArrowRight",
			"ArrowLeft",
			"ArrowUp",
		},

		VariantOf: "ChristmasAdventure-v0",
	},

	{
		Name:      "ChirstmasAdventure-v2",
		BaseURL:   "https://www.yiv.com/games/Christmas-Adventure/",
		Width:     533,
		Height:    320,
		AllCanvas: true,
		Options:   "{\"level\":3}",

		KeyWhitelist: []string{
			"KeyA",
			"ArrowRight",
			"ArrowLeft",
			"ArrowUp",
		},

		VariantOf: "ChristmasAdventure-v0",
	},

	{
		Name:      "ChirstmasAdventure-v3",
		BaseURL:   "https://www.yiv.com/games/Christmas-Adventure/",
		Width:     533,
		Height:    320,
		AllCanvas: true,
		Options:   "{\"level\":4}",

		KeyWhitelist: []string{
			"KeyA",
			"ArrowRight",
			"ArrowLeft",
			"ArrowUp",
		},

		VariantOf: "ChristmasAdventure-v0",
	},

	{
		Name:      "ChirstmasAdventure-v4",
		BaseURL:   "https://www.yiv.com/games/Christmas-Adventure/",
		Width:     533,
		Height:    320,
		AllCanvas: true,
		Options:   "{\"level\":5}",

		KeyWhitelist: []string{
			"KeyA",
			"ArrowRight",
			"ArrowLeft",
			"ArrowUp",
		},

		VariantOf: "ChristmasAdventure-v0",
	},

	{
		Name:      "ChirstmasAdventure-v5",
		BaseURL:   "https://www.yiv.com/games/Christmas-Adventure/",
		Width:     533,
		Height:    320,
		AllCanvas: true,
		Options:   "{\"level\":6}",

		KeyWhitelist: []string{
			"KeyA",
			"ArrowRight",
			"ArrowLeft",
			"ArrowUp",
		},

		VariantOf: "ChristmasAdventure-v0",
	},

	{
		Name:      "ChirstmasAdventure-v6",
		BaseURL:   "https://www.yiv.com/games/Christmas-Adventure/",
		Width:     533,
		Height:    320,
		AllCanvas: true,
		Options:   "{\"level\":7}",

		KeyWhitelist: []string{
			"KeyA",
			"ArrowRight",
			"ArrowLeft",
			"ArrowUp",
		},

		VariantOf: "ChristmasAdventure-v0",
	},

	{
		Name:      "ChirstmasAdventure-v7",
		BaseURL:   "https://www.yiv.com/games/Christmas-Adventure/",
		Width:     533,
		Height:    320,
		AllCanvas: true,
		Options:   "{\"level\":8}",

		KeyWhitelist: []string{
			"KeyA",
			"ArrowRight",
			"ArrowLeft",
			"ArrowUp",
		},

		VariantOf: "ChristmasAdventure-v0",
	},

	{
		Name:      "ChirstmasAdventure-v8",
		BaseURL:   "https://www.yiv.com/games/Christmas-Adventure/",
		Width:     533,
		Height:    320,
		AllCanvas: true,
		Options:   "{\"level\":9}",

		KeyWhitelist: []string{
			"KeyA",
			"ArrowRight",
			"ArrowLeft",
			"ArrowUp",
		},

		VariantOf: "ChristmasAdventure-v0",
	},

	{
		Name:      "ChirstmasAdventure-v10",
		BaseURL:   "https://www.yiv.com/games/Christmas-Adventure/",
		Width:     533,
		Height:    320,
		AllCanvas: true,
		Options:   "{\"level\":11}",

		KeyWhitelist: []string{
			"KeyA",
			"ArrowRight",
			"ArrowLeft",
			"ArrowUp",
		},

		VariantOf: "ChristmasAdventure-v0",
	},

	{
		Name:      "ChirstmasAdventure-v11",
		BaseURL:   "https://www.yiv.com/games/Christmas-Adventure/",
		Width:     533,
		Height:    320,
		AllCanvas: true,
		Options:   "{\"level\":12}",

		KeyWhitelist: []string{
			"KeyA",
			"ArrowRight",
			"ArrowLeft",
			"ArrowUp",
		},

		VariantOf: "ChristmasAdventure-v0",
	},

	{
		Name:      "ChirstmasAdventure-v12",
		BaseURL:   "https://www.yiv.com/games/Christmas-Adventure/",
		Width:     533,
		Height:    320,
		AllCanvas: true,
		Options:   "{\"level\":13}",

		KeyWhitelist: []string{
			"KeyA",
			"ArrowRight",
			"ArrowLeft",
			"ArrowUp",
		},

		VariantOf: "ChristmasAdventure-v0",
	},

	{
		Name:      "ChirstmasAdventure-v13",
		BaseURL:   "https://www.yiv.com/games/Christmas-Adventure/",
		Width:     533,
		Height:    320,
		AllCanvas: true,
		Options:   "{\"level\":14}",

		KeyWhitelist: []string{
			"KeyA",
			"ArrowRight",
			"ArrowLeft",
			"ArrowUp",
		},

		VariantOf: "ChristmasAdventure-v0",
	},

	{
		Name:      "ChirstmasAdventure-v14",
		BaseURL:   "https://www.yiv.com/games/Christmas-Adventure/",
		Width:     533,
		Height:    320,
		AllCanvas: true,
		Options:   "{\"level\":15}",

		KeyWhitelist: []string{
			"KeyA",
			"ArrowRight",
			"ArrowLeft",
			"ArrowUp",
		},

		VariantOf: "ChristmasAdventure-v0",
	},

	{
		Name:      "ChirstmasAdventure-v15",
		BaseURL:   "https://www.yiv.com/games/Christmas-Adventure/",
		Width:     533,
		Height:    320,
		AllCanvas: true,
		Options:   "{\"level\":16}",

		KeyWhitelist: []string{
			"KeyA",
			"ArrowRight",
			"ArrowLeft",
			"ArrowUp",
		},

		VariantOf: "ChristmasAdventure-v0",
	},

	{
		Name:      "ChirstmasAdventure-v16",
		BaseURL:   "https://www.yiv.com/games/Christmas-Adventure/",
		Width:     533,
		Height:    320,
		AllCanvas: true,
		Options:   "{\"level\":17}",

		KeyWhitelist: []string{
			"KeyA",
			"ArrowRight",
			"ArrowLeft",
			"ArrowUp",
		},

		VariantOf: "ChristmasAdventure-v0",
	},

	{
		Name:      "ChirstmasAdventure-v17",
		BaseURL:   "https://www.yiv.com/games/Christmas-Adventure/",
		Width:     533,
		Height:    320,
		AllCanvas: true,
		Options:   "{\"level\":18}",

		KeyWhitelist: []string{
			"KeyA",
			"ArrowRight",
			"ArrowLeft",
			"ArrowUp",
		},

		VariantOf: "ChristmasAdventure-v0",
	},

	{
		Name:      "ChirstmasAdventure-v18",
		BaseURL:   "https://www.yiv.com/games/Christmas-Adventure/",
		Width:     533,
		Height:    320,
		AllCanvas: true,
		Options:   "{\"level\":19}",

		KeyWhitelist: []string{
			"KeyA",
			"ArrowRight",
			"ArrowLeft",
			"ArrowUp",
		},

		VariantOf: "ChristmasAdventure-v0",
	},

	{
		Name:      "CandySuperLines-v0",
		BaseURL:   "https://www.yiv.com/games/Candy-Super-Lines-Match3/",
		Width:     568,
		Height:    320,
		AllCanvas: true,
		Options:   "{}",
	},

	{
		Name:      "MiniRacer-v0",
		BaseURL:   "https://h5.4j.com/games/Mini-Racer/",
		Width:     320,
		Height:    512,
		AllCanvas: true,
		Options:   "{\"exit\":[\"7020333949850896\",\"7439249625248278\",\"2135357787493767\",\"8480831502092825\"]}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
			"ArrowUp",
			"ArrowDown",
			"Space",
		},
	},

	{
		Name:      "ChaseRacingCars-v0",
		BaseURL:   "https://h5.4j.com/games/Chase-Racing-Cars/",
		Width:     320,
		Height:    426,
		AllCanvas: true,
		Options:   "{}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
			"Space",
		},
	},

	{
		Name:      "PineapplePen-v0",
		BaseURL:   "https://h5.4j.com/games/Pineapple-Pen-Online/",
		Width:     320,
		Height:    533,
		AllCanvas: true,
		Options:   "{}",
	},

	{
		Name:      "TapTapDash-v0",
		BaseURL:   "https://h5.4j.com/games/Tap-Tap-Dash-Online/",
		Width:     320,
		Height:    533,
		AllCanvas: true,
		Options:   "{}",
	},

	{
		Name:      "BeachKiss-v0",
		BaseURL:   "https://www.yiv.com/games/Beach-Kiss/",
		Width:     320,
		Height:    533,
		AllCanvas: true,
		Options:   "{}",
	},

	{
		Name:      "BallsVsBlocks-v0",
		BaseURL:   "http://h5.4j.com/games/Balls-Vs-Blocks-Online/",
		Width:     320,
		Height:    533,
		AllCanvas: true,
		Options:   "{}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
		},
	},

	{
		Name:      "RunIntoDeath-v0",
		BaseURL:   "http://h5.4j.com/games/Run-Into-Death/",
		Width:     480,
		Height:    320,
		AllCanvas: true,
		Options:   "{}",
	},

	{
		Name:      "TrackRacer-v0",
		BaseURL:   "https://h5.4j.com/games/Track-Racer/",
		Width:     320,
		Height:    512,
		AllCanvas: false,
		Options:   "{\"exit\":[\"942555504266250\",\"474343296605419\",\"710812321871081\",\"629012887341791\"]}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
			"ArrowUp",
			"ArrowDown",
			"Space",
		},
	},

	{
		Name:      "Cars-v0",
		BaseURL:   "https://h5.4j.com/games/Cars/",
		Width:     320,
		Height:    512,
		AllCanvas: false,
		Options:   "{\"exit\":[\"8480831502092825\",\"190710987972732\",\"2135357787493767\",\"9835750250435948\"]}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
			"ArrowUp",
			"ArrowDown",
		},
	},

	{
		Name:      "PopUp-v0",
		BaseURL:   "http://h5.4j.com/games/Pop-Up/",
		Width:     320,
		Height:    480,
		AllCanvas: true,
		Options:   "{}",
	},

	{
		Name:      "NinjaRun-v0",
		BaseURL:   "http://h5.4j.com/games/Ninja-Run/",
		Width:     568,
		Height:    320,
		AllCanvas: false,
		Options:   "{\"scoreMode\":\"stars\"}",
	},

	{
		Name:      "NinjaRun-v1",
		BaseURL:   "http://h5.4j.com/games/Ninja-Run/",
		Width:     568,
		Height:    320,
		AllCanvas: false,
		Options:   "{\"scoreMode\":\"distance\"}",

		VariantOf: "NinjaRun-v0",
	},

	{
		Name:      "NinjaRun-v2",
		BaseURL:   "http://h5.4j.com/games/Ninja-Run/",
		Width:     568,
		Height:    320,
		AllCanvas: false,
		Options:   "{\"scoreMode\":\"coins\"}",

		VariantOf: "NinjaRun-v0",
	},

	{
		Name:      "AmazingNinja-v0",
		BaseURL:   "http://h5.4j.com/games/Amazing-Ninja/",
		Width:     320,
		Height:    480,
		AllCanvas: true,
		Options:   "{}",
	},

	{
		Name:      "StickmanArcher-v0",
		BaseURL:   "http://h5.4j.com/games/Stickman-Archer-Online/",
		Width:     569,
		Height:    320,
		AllCanvas: true,
		Options:   "{}",
	},

	{
		Name:      "RiseUp-v0",
		BaseURL:   "https://h5.4j.com/games/Rise-Up/",
		Width:     320,
		Height:    480,
		AllCanvas: true,
		Options:   "{}",
	},

	{
		Name:      "ColorTease-v0",
		BaseURL:   "http://h5.4j.com/games/Color-Tease/",
		Width:     320,
		Height:    569,
		AllCanvas: true,
		Options:   "{}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
		},
	},

	{
		Name:      "TRexRunner-v0",
		BaseURL:   "http://h5.4j.com/games/T-Rex-Runner/",
		Width:     569,
		Height:    320,
		AllCanvas: false,
		Options:   "{}",

		KeyWhitelist: []string{
			"ArrowUp",
			"Space",
		},
	},

	{
		Name:      "TruckTravel-v0",
		BaseURL:   "http://h5.4j.com/games/Truck-Travel/",
		Width:     569,
		Height:    320,
		AllCanvas: true,
		Options:   "{\"character\":0}",

		KeyWhitelist: []string{
			"ArrowUp",
			"ArrowDown",
		},
	},

	{
		Name:      "TruckTravel-v1",
		BaseURL:   "http://h5.4j.com/games/Truck-Travel/",
		Width:     569,
		Height:    320,
		AllCanvas: true,
		Options:   "{\"character\":1}",

		KeyWhitelist: []string{
			"ArrowUp",
			"ArrowDown",
		},

		VariantOf: "TruckTravel-v0",
	},

	{
		Name:      "PanicDrop-v0",
		BaseURL:   "http://h5.4j.com/games/Panic-Drop/",
		Width:     320,
		Height:    569,
		AllCanvas: true,
		Options:   "{\"scoreMode\":\"distance\"}",
	},

	{
		Name:      "PanicDrop-v1",
		BaseURL:   "http://h5.4j.com/games/Panic-Drop/",
		Width:     320,
		Height:    569,
		AllCanvas: true,
		Options:   "{\"scoreMode\":\"coins\"}",

		VariantOf: "PanicDrop-v0",
	},

	{
		Name:      "BlockSnake-v0",
		BaseURL:   "http://h5.4j.com/games/Block-Snake/",
		Width:     320,
		Height:    569,
		AllCanvas: true,
		Options:   "{\"level\":1}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
		},
	},

	{
		Name:      "BlockSnake-v1",
		BaseURL:   "http://h5.4j.com/games/Block-Snake/",
		Width:     320,
		Height:    569,
		AllCanvas: true,
		Options:   "{\"level\":2}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
		},

		VariantOf: "BlockSnake-v0",
	},

	{
		Name:      "BlockSnake-v2",
		BaseURL:   "http://h5.4j.com/games/Block-Snake/",
		Width:     320,
		Height:    569,
		AllCanvas: true,
		Options:   "{\"level\":3}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
		},

		VariantOf: "BlockSnake-v0",
	},

	{
		Name:      "BlockSnake-v3",
		BaseURL:   "http://h5.4j.com/games/Block-Snake/",
		Width:     320,
		Height:    569,
		AllCanvas: true,
		Options:   "{\"level\":4}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
		},

		VariantOf: "BlockSnake-v0",
	},

	{
		Name:      "BlockSnake-v4",
		BaseURL:   "http://h5.4j.com/games/Block-Snake/",
		Width:     320,
		Height:    569,
		AllCanvas: true,
		Options:   "{\"level\":5}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
		},

		VariantOf: "BlockSnake-v0",
	},

	{
		Name:      "BlockSnake-v5",
		BaseURL:   "http://h5.4j.com/games/Block-Snake/",
		Width:     320,
		Height:    569,
		AllCanvas: true,
		Options:   "{\"level\":6}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
		},

		VariantOf: "BlockSnake-v0",
	},

	{
		Name:      "TrafficRacer-v0",
		BaseURL:   "http://h5.4j.com/games/Traffic-Racer/",
		Width:     320,
		Height:    480,
		AllCanvas: true,
		Options:   "{}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
			"ArrowUp",
			"ArrowDown",
		},
	},

	{
		Name:      "FlappyBird-v0",
		BaseURL:   "http://flappybird.io/",
		Width:     320,
		Height:    512,
		AllCanvas: true,
		Options:   "{}",

		KeyWhitelist: []string{
			"Space",
		},
	},

	{
		Name:      "PacMan-v0",
		BaseURL:   "https://fir.sh/projects/jsnes/",
		Width:     256,
		Height:    240,
		AllCanvas: true,
		Options:   "{}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
			"ArrowUp",
			"ArrowDown",
		},
	},

	{
		Name:      "Tetris-v0",
		BaseURL:   "https://fir.sh/projects/jsnes/",
		Width:     256,
		Height:    240,
		AllCanvas: true,
		Options:   "{}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
			"ArrowDown",
			"KeyX",
			"KeyY",
		},
	},

	{
		Name:      "CubeNinja-v0",
		BaseURL:   "http://h5.4j.com/games/Cube-Ninja/",
		Width:     480,
		Height:    320,
		AllCanvas: true,
		Options:   "{}",
	},

	{
		Name:      "SuperCowboyRun-v0",
		BaseURL:   "http://h5.4j.com/games/Super-Cowboy-Run/",
		Width:     569,
		Height:    320,
		AllCanvas: true,
		Options:   "{\"scoreMode\":\"distance\"}",

		KeyWhitelist: []string{
			"ArrowUp",
			"Space",
		},
	},

	{
		Name:      "SuperCowboyRun-v1",
		BaseURL:   "http://h5.4j.com/games/Super-Cowboy-Run/",
		Width:     569,
		Height:    320,
		AllCanvas: true,
		Options:   "{\"scoreMode\":\"money\"}",

		KeyWhitelist: []string{
			"ArrowUp",
			"Space",
		},

		VariantOf: "SuperCowboyRun-v0",
	},

	{
		Name:      "SpikeBird-v0",
		BaseURL:   "http://h5.4j.com/games/Spike-Bird/",
		Width:     288,
		Height:    512,
		AllCanvas: true,
		Options:   "{}",
	},

	{
		Name:      "GeoJump-v0",
		BaseURL:   "http://h5.4j.com/games/Geo-Jump/",
		Width:     288,
		Height:    512,
		AllCanvas: true,
		Options:   "{}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
		},
	},

	{
		Name:      "JumpyShark-v0",
		BaseURL:   "http://h5.4j.com/games/Jumpy-Shark/",
		Width:     320,
		Height:    480,
		AllCanvas: true,
		Options:   "{\"difficulty\":\"easy\"}",
	},

	{
		Name:      "JumpyShark-v1",
		BaseURL:   "http://h5.4j.com/games/Jumpy-Shark/",
		Width:     320,
		Height:    480,
		AllCanvas: true,
		Options:   "{\"difficulty\":\"normal\"}",

		VariantOf: "JumpyShark-v0",
	},

	{
		Name:      "JumpyShark-v2",
		BaseURL:   "http://h5.4j.com/games/Jumpy-Shark/",
		Width:     320,
		Height:    480,
		AllCanvas: true,
		Options:   "{\"difficulty\":\"hard\"}",

		VariantOf: "JumpyShark-v0",
	},

	{
		Name:      "StickSanta-v0",
		BaseURL:   "http://h5.4j.com/games/Stick-Santa/",
		Width:     288,
		Height:    512,
		AllCanvas: true,
		Options:   "{}",
	},

	{
		Name:      "BlockIt-v0",
		BaseURL:   "http://h5.4j.com/games/Block-It-Online/",
		Width:     320,
		Height:    533,
		AllCanvas: true,
		Options:   "{}",
	},

	{
		Name:      "MagicRun-v0",
		BaseURL:   "http://h5.4j.com/games/Magic-Run/",
		Width:     320,
		Height:    569,
		AllCanvas: true,
		Options:   "{}",

		KeyWhitelist: []string{
			"ArrowLeft",
			"ArrowRight",
			"ArrowUp",
		},
	},

	{
		Name:      "HopHero-v0",
		BaseURL:   "http://h5.4j.com/games/Hop-Hero/",
		Width:     320,
		Height:    569,
		AllCanvas: true,
		Options:   "{\"scoreMode\":\"distance\"}",

		KeyWhitelist: []string{
			"Space",
			"ArrowUp",
		},
	},

	{
		Name:      "HopHero-v1",
		BaseURL:   "http://h5.4j.com/games/Hop-Hero/",
		Width:     320,
		Height:    569,
		AllCanvas: true,
		Options:   "{\"scoreMode\":\"coins\"}",

		KeyWhitelist: []string{
			"Space",
			"ArrowUp",
		},

		VariantOf: "HopHero-v0",
	},

	{
		Name:      "Spacecraft-v0",
		BaseURL:   "https://js13kgames.com/games/spacecraft/",
		Width:     320,
		Height:    480,
		AllCanvas: false,
		Options:   "{}",

		KeyWhitelist: []string{
			"Space",
			"ArrowUp",
			"ArrowLeft",
			"ArrowRight",
			"ArrowDown",
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
