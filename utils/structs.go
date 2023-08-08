package utils

type RoomData struct {
	Status  string       `json:"status"`
	Id      string       `json:"id"`
	Setting RoomSetting  `json:"setting"`
	Members []RoomMember `json:"members"`
}

type RoomSetting struct {
	GameMode int    `json:"game_mode"`
	Engine   int    `json:"engine"`
	Course   string `json:"course"`
}

type RoomMember struct {
	FriendCode   string `json:"friend_code"`
	Name         string `json:"name"`
	GuestName    string `json:"guest_name,omitempty"`
	RaceRating   int    `json:"vr"`
	BattleRating int    `json:"br"`
}
