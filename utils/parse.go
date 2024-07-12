package utils

func CheckGameMode(raceMode int) int8 {
	switch raceMode {
	case 0x00:
		return ModeNone
	case 0x10:
		return ModeVS
	case 0x11:
		return ModePrivateVS
	case 0x30:
		return ModeBalloonBattle
	case 0x31:
		return ModePrivateBalloonBattle
	case 0x40:
		return ModeCoinBattle
	case 0x41:
		return ModePrivateCoinBattle
	default:
		return ModeUnknown
	}
}
