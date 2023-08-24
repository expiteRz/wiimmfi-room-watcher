package utils

type RoomData struct {
	Status    string       `json:"status"`
	Id        string       `json:"id"`
	Setting   RoomSetting  `json:"setting"`
	MemberLen int          `json:"player_amount"`
	Members   []RoomMember `json:"members"`
}

type RoomSetting struct {
	GameMode int    `json:"game_mode"`
	Engine   int    `json:"engine"`
	Course   string `json:"course"`
	CourseId int    `json:"course_id"`
}

type RoomMember struct {
	Pid          int    `json:"pid"`
	FriendCode   string `json:"friend_code"`
	Name         string `json:"name"`
	GuestName    string `json:"guest_name"`
	RaceRating   int    `json:"vr"`
	BattleRating int    `json:"br"`
	Status       string `json:"role"`
	FinishTimes  []int  `json:"finish_times"`
}

type MemberCourse struct {
	Name    string `json:"name"`
	Id      int    `json:"id"`
	Allowed string `json:"allowed"`
}
