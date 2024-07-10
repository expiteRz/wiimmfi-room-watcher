package utils

func CheckGameMode(ol int) (mode int) {
	switch ol & 0b10000011000110 {
	case 0b01000010:
		mode = ModePrivateVS
	case 0b10000010:
		mode = ModePrivateCoinBattle
	case 0b10000010000010:
		mode = ModePrivateBalloonBattle
	case 0b01000100:
		mode = ModeVS
	case 0b10000100:
		mode = ModeCoinBattle
	case 0b10000010000100:
		mode = ModeBalloonBattle
	default:
		mode = ModeNone
	}
	return
}

func CheckGameModeTEST(raceMode int) int8 {
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
