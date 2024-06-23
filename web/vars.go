package web

import (
	"embed"
	_ "embed"
)

//go:embed assets/templates/* assets/deps/*
var webUiAssets embed.FS

func GetWebAsset(filename string) string {
	file, err := webUiAssets.ReadFile(filename)
	if err != nil {
		return ""
	}
	return string(file)
}

const overlayItemTemplate = `
<div class="overlay-item">
  <div>
    <h2>{NAME}</h2>
    <button class="button open_button" id="{ID}" to="{NAME}">Open folder</button>
  </div>
  <iframe src="{URL}" width="500" height="300" scrolling="no" frameborder="0"></iframe>
</div>
`

const settingItemTemplate = `
<div class="setting-item">
  <div>
    <h3>{NAME}</h3>
    <p>{DESC}</p>
  </div>
  {INPUT}
</div>
`

const settingItemInputTemplate = `
<label class="{TYPE}">
  <input type="{TYPE}" name="{NAME}" id="{ID}" {ADDON} value="{VALUE}" />
</label>
`

const settingSubmitSaveTemplate = `
<button class="button" id="save_settings">Save</button>
`

const internalErrorTemplate = `
<div>500 internal server error</div>
`

var settingItemDict = []map[string]string{
	{"id": "pid", "name": "Target PID", "description": "PID to watch the current room statistic"},
	{"id": "interval", "name": "Interval", "description": "Frequency in seconds to update statistic (10 as default, least 5)"},
	{"id": "serverip", "name": "Host address", "description": "IP address to host the API and overlays (127.0.0.1 as default)"},
	{"id": "serverport", "name": "Host port", "description": "Targeted port number to host the API and overlays (24050 as default)"},
	//{"id": "open_at_startup", "name": "Open dashboard on startup", "description": "Open dashboard in browser on startup"},
}

// All images borrowed from mariowiki
var NinImagePath = map[int]string{
	3445: "8/8b/Luigi_Circuit_MKWii.png/800px-Luigi_Circuit_MKWii.png",
	3450: "f/f4/Moo_Moo_Meadows_MKWii.png/800px-Moo_Moo_Meadows_MKWii.png",
	3454: "8/84/Mushroom_Gorge_MKWii.png/800px-Mushroom_Gorge_MKWii.png",
	3464: "5/54/Toad%27s_Factory_MKWii.png/800px-Toad%27s_Factory_MKWii.png",

	3448: "b/be/MarioCircuit_MKWii.png/800px-MarioCircuit_MKWii.png",
	3424: "7/7d/Coconut_Mall_MKWii.png/800px-Coconut_Mall_MKWii.png",
	3426: "8/89/DK_Summit_MKWii.png/800px-DK_Summit_MKWii.png",
	3466: "2/21/MKW_Wario%27s_Gold_Mine_Course_Overview.png/800px-MKW_Wario%27s_Gold_Mine_Course_Overview.png",

	3431: "8/87/Daisy_circuit1.png/800px-Daisy_circuit1.png",
	3443: "3/37/MKW_Koopa_Cape.png/800px-MKW_Koopa_Cape.png",
	3447: "e/e9/MKW_Maple_Treeway_Overlook.png/800px-MKW_Maple_Treeway_Overlook.png",
	3441: "5/55/Grumble_Volcano.png/800px-Grumble_Volcano.png",

	3433: "d/d9/Dry_Dry_Ruins.png/800px-Dry_Dry_Ruins.png",
	3452: "d/db/Moonviewhighway.png/800px-Moonviewhighway.png",
	3422: "b/b0/BowsersCastleMKW.png/800px-BowsersCastleMKW.png",
	3460: "8/83/MKW_Rainbow_Road_Overview.png/800px-MKW_Rainbow_Road_Overview.png",

	3439: "1/18/Peach_Beach_MKWii.png/800px-Peach_Beach_MKWii.png",
	3430: "e/e6/Yoshi_Falls_MKWii.png/800px-Yoshi_Falls_MKWii.png",
	3462: "e/e1/Ghost_Valley_2_MKWii.png/800px-Ghost_Valley_2_MKWii.png",
	3458: "3/35/Mario_Raceway_64_MKWii.png/800px-Mario_Raceway_64_MKWii.png",

	3459: "4/4b/Sherbet_Land_N64_MKWii.png/800px-Sherbet_Land_N64_MKWii.png",
	3436: "0/06/Shy_Guy_Beach_MKWii.png/800px-Shy_Guy_Beach_MKWii.png",
	3427: "a/ab/Delfino_Square_MKWii.png/800px-Delfino_Square_MKWii.png",
	3440: "6/61/Waluigi_Stadium_MKWii.png/800px-Waluigi_Stadium_MKWii.png",

	3428: "4/44/Desert_Hills.png/800px-Desert_Hills.png",
	3435: "e/e1/MKW_GBA_Bowser_Castle_3_Intro.png/800px-MKW_GBA_Bowser_Castle_3_Intro.png",
	3457: "b/bc/D.K.'s_Jungle_Parkway.png/800px-D.K.'s_Jungle_Parkway.png",
	3438: "9/9e/MKWii_GCNMario.png/800px-MKWii_GCNMario.png",

	3463: "f/f9/MKW_SNES_Mario_Circuit_3_Overview.png/800px-MKW_SNES_Mario_Circuit_3_Overview.png",
	3429: "e/e2/MKWii_Peach_Gardens.png/800px-MKWii_Peach_Gardens.png",
	3437: "c/c2/DK_Mountain.png/800px-DK_Mountain.png",
	3456: "e/e5/N64BowserCastle-MKWii.png/800px-N64BowserCastle-MKWii.png",

	3409: "e/e6/Block_Plaza.png/800px-Block_Plaza.png",
	3413: "1/11/Delfino_Pier.png/800px-Delfino_Pier.png",
	3415: "b/b9/MKW_Funky_Stadium_Overview.png/800px-MKW_Funky_Stadium_Overview.png",
	3410: "d/d3/Chain_Chomp_Roulette.png/800px-Chain_Chomp_Roulette.png",
	3420: "9/9a/Thwomp_Desert.png/800px-Thwomp_Desert.png",

	3419: "2/24/Battle_Course_4_(SNES).png/800px-Battle_Course_4_(SNES).png",
	3416: "0/01/Battle_Course_3_(GBA).png/800px-Battle_Course_3_(GBA).png",
	3418: "3/31/N64Skyscraper-MKWii.png/800px-N64Skyscraper-MKWii.png",
	3417: "9/90/Cookie_Land.png/800px-Cookie_Land.png",
	3412: "f/f0/Twilight_House.png/800px-Twilight_House.png",
}
