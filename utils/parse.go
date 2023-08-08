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
