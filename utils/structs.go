package utils

type RoomData struct {
	Status    string       `json:"status"`
	Id        string       `json:"id"`
	Setting   RoomSetting  `json:"setting"`
	MemberLen int          `json:"player_amount"`
	Members   []RoomMember `json:"members"`
}

type RoomSetting struct {
	GameMode     int    `json:"game_mode"`
	Engine       int    `json:"engine"`
	Course       string `json:"course"`
	CourseId     int    `json:"course_id"`
	ThumbnailUrl string `json:"thumbnail_url"`
}

type RoomMember struct {
	Pid          int          `json:"pid"`
	FriendCode   string       `json:"friend_code"`
	Name         string       `json:"name"`
	GuestName    string       `json:"guest_name"`
	RaceRating   int          `json:"vr"`
	BattleRating int          `json:"br"`
	Status       string       `json:"role"`
	FinishTimes  []int        `json:"finish_times"` // FinishTimes[0] = Primary players time, FinishTimes[1] = Guest players time
	Course       MemberCourse `json:"course"`
	Combos       []Combo      `json:"combos"` // Combo[0] = Primary players combo, Combo[1] = Guest players combo
}

type MemberCourse struct {
	Name    string `json:"name"`
	Id      int    `json:"id"`
	Allowed string `json:"allowed"`
}

type Combo struct {
	Character ComboChild `json:"character"`
	Vehicle   ComboChild `json:"vehicle"`
}

type ComboChild struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
