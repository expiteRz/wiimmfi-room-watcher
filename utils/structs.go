package utils

import "github.com/rs/zerolog"

type RoomData struct {
	Status    JsonStatus  `json:"status"`
	Id        string      `json:"id"`
	Setting   RoomSetting `json:"setting"`
	MemberLen int         `json:"player_amount"`
	Members   RoomMembers `json:"members"`
}

type JsonStatus struct {
	Result string `json:"result"`
	Reason string `json:"reason"`
}

type RoomSetting struct {
	GameMode     int8   `json:"game_mode"`
	GameModeText string `json:"game_mode_text"`
	RaceCount    int    `json:"race_count"`
	Engine       int    `json:"engine"`
	EngineText   string `json:"engine_text"`
	CourseId     int    `json:"course_id"`
	Course       string `json:"course"`
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
	Combos       []Combo      `json:"combos"`     // Combo[0] = Primary players combo, Combo[1] = Guest players combo
	DelayTime    int          `json:"delay_time"` // milliseconds
	ConnFail     float32      `json:"conn_fail"`
}

type RoomMembers []RoomMember

func (mm RoomMembers) MarshalZerologArray(a *zerolog.Array) {
	for _, m := range mm {
		a.Str(m.Name)
		if m.GuestName != "" {
			a.Str(m.GuestName)
		}
	}
}

type MemberCourse struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

type Combo struct {
	Character ComboChild `json:"character"`
	Vehicle   ComboChild `json:"vehicle"`
}

type ComboChild struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
