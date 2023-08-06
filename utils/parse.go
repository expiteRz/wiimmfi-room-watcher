package utils

func CheckGameMode(ol int) (mode int) {
	//fmt.Printf("%b\n", ol&0b11000110)

	switch ol & 0b10000011000110 {
	case 0b01000010:
		mode = ModePrivateVS
	case 0b10000010:
		mode = ModePrivateCoinBattle
	case 0b10000010000010:
		mode = ModePrivateBalloonBattle
	case 0b01000100:
		mode = ModeVS
	case 0b10000010000100:
		mode = ModeCoinBattle
	default:
		mode = ModeNone
	}

	return
}
