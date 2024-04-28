package utils

var (
	ENGINE = []string{
		"50cc",
		"100cc",
		"150cc",
		"Mirror",
	}
	GAMEMODE = map[int]string{
		ModeNone:                 "Unknown game mode",
		ModeVS:                   "Public VS",
		ModeCoinBattle:           "Public Coin Battle",
		ModeBalloonBattle:        "Public Balloon Battle",
		ModePrivateVS:            "Private Room VS",
		ModePrivateCoinBattle:    "Private Room Coin Battle",
		ModePrivateBalloonBattle: "Private Room Balloon Battle",
	}
)

const (
	ModeNone = iota - 1
	ModeVS
	ModeCoinBattle
	ModeBalloonBattle
	ModePrivateVS = iota - 1 + 0x10
	ModePrivateCoinBattle
	ModePrivateBalloonBattle
)

// https://github.com/riidefi/mkw/blob/master/source/game/system/ResourceManager.hpp#L25
type CharacterId int

const (
	Mario CharacterId = iota
	BabyPeach
	Waluigi
	Bowser
	BabyDaisy
	DryBones
	BabyMario
	Luigi
	Toad
	DK
	Yoshi
	Wario
	BabyLuigi
	Toadette
	Nokonoko
	Daisy
	Peach
	Birdo
	Diddy
	Boo
	BowserJr
	DryBowser
	Funky
	Rosalina
	MiiSAMale
	MiiSAFemale
	MiiSBMale
	MiiSBFemale
	MiiSCMale
	MiiSCFemale
	MiiMAMale
	MiiMAFemale
	MiiMBMale
	MiiMBFemale
	MiiMCMale
	MiiMCFemale
	MiiLAMale
	MiiLAFemale
	MiiLBMale
	MiiLBFemale
	MiiLCMale
	MiiLCFemale
	MiiM
	MiiS
	MiiL
	PeachBiker
	DaisyBiker
	RosalinaBiker
)

func (c CharacterId) String() (s string) {
	switch c {
	case Mario:
		s = "Mario"
	case BabyPeach:
		s = "Baby Peach"
	case Waluigi:
		s = "Waluigi"
	case Bowser:
		s = "Bowser"
	case BabyDaisy:
		s = "Baby Daisy"
	case DryBones:
		s = "Dry Bones"
	case BabyMario:
		s = "Baby Mario"
	case Luigi:
		s = "Luigi"
	case Toad:
		s = "Toad"
	case DK:
		s = "Donkey Kong"
	case Yoshi:
		s = "Yoshi"
	case Wario:
		s = "Wario"
	case BabyLuigi:
		s = "Baby Luigi"
	case Toadette:
		s = "Toadette"
	case Nokonoko:
		s = "Koopa Troopa"
	case Daisy, DaisyBiker:
		s = "Daisy"
	case Peach, PeachBiker:
		s = "Peach"
	case Birdo:
		s = "Birdo"
	case Diddy:
		s = "Diddy Kong"
	case Boo:
		s = "King Boo"
	case BowserJr:
		s = "Bowser Jr."
	case DryBowser:
		s = "Dry Bowser"
	case Funky:
		s = "Funky Kong"
	case Rosalina, RosalinaBiker:
		s = "Rosalina"
	case MiiSAMale, MiiSAFemale:
		s = "Mii S A"
	case MiiSBMale, MiiSBFemale:
		s = "Mii S B"
	case MiiMAMale, MiiMAFemale:
		s = "Mii M A"
	case MiiMBMale, MiiMBFemale:
		s = "Mii M B"
	case MiiLAMale, MiiLAFemale:
		s = "Mii L A"
	case MiiLBMale, MiiLBFemale:
		s = "Mii L B"
	default:
		s = "Unknown"
	}
	return
}

// https://github.com/riidefi/mkw/blob/master/source/game/system/ResourceManager.hpp#L77
type VehicleId int

const (
	StandardKartS VehicleId = iota
	StandardKartM
	StandardKartL
	BabyBooster
	ClassicDragster
	Offroader
	MiniBeast
	WildWing
	FlameFlyer
	CheepCharger
	SuperBlooper
	PiranhaProwler
	RallyRomper
	RoyalRacer
	Jetsetter
	BlueFalcon
	Sprinter
	Honeycoupe
	StandardBikeS
	StandardBikeM
	StandardBikeL
	BulletBike
	MachBike
	BowserBike
	BitBike
	BonBon
	WarioBike
	Quacker
	Rapide
	ShootingStar
	Magikruiser
	Nitrocycle
	Spear
	JetBubble
	DolphinDasher
	Phantom
)

func (v VehicleId) String() (s string) {
	// NTSC-U Format
	switch v {
	case StandardKartS:
		s = "Standard Kart S"
	case StandardKartM:
		s = "Standard Kart M"
	case StandardKartL:
		s = "Standard Kart L"
	case BabyBooster:
		s = "Booster Seat"
	case ClassicDragster:
		s = "Classic Dragster"
	case Offroader:
		s = "Offroader"
	case MiniBeast:
		s = "Mini Beast"
	case WildWing:
		s = "Wild Wing"
	case FlameFlyer:
		s = "Flame Flyer"
	case CheepCharger:
		s = "Cheep Charger"
	case SuperBlooper:
		s = "Super Blooper"
	case PiranhaProwler:
		s = "Piranha Prowler"
	case RallyRomper:
		s = "Tiny Titan"
	case RoyalRacer:
		s = "Daytripper"
	case Jetsetter:
		s = "Jetsetter"
	case BlueFalcon:
		s = "Blue Falcon"
	case Sprinter:
		s = "Sprinter"
	case Honeycoupe:
		s = "Honeycoupe"
	case StandardBikeS:
		s = "Standard Bike S"
	case StandardBikeM:
		s = "Standard Bike M"
	case StandardBikeL:
		s = "Standard Bike L"
	case BulletBike:
		s = "Bullet Bike"
	case MachBike:
		s = "Mach Bike"
	case BowserBike:
		s = "Flame Runner"
	case BitBike:
		s = "Bit Bike"
	case BonBon:
		s = "Sugarscoot"
	case WarioBike:
		s = "Wario Bike"
	case Quacker:
		s = "Quacker"
	case Rapide:
		s = "Zip Zip"
	case ShootingStar:
		s = "Shooting Star"
	case Magikruiser:
		s = "Magikruiser"
	case Nitrocycle:
		s = "Sneakster"
	case Spear:
		s = "Spear"
	case JetBubble:
		s = "Jet Bubble"
	case DolphinDasher:
		s = "Dolphin Dasher"
	case Phantom:
		s = "Phantom"
	default:
		s = "Unknown"
	}
	return
}
