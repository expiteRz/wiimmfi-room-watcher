package utils

var (
	ENGINE = []string{
		"50cc",
		"100cc",
		"150cc",
		"Mirror",
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
