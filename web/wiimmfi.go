package web

import (
	"app.rz-public.xyz/wiimmfi-room-watcher/utils"
	"app.rz-public.xyz/wiimmfi-room-watcher/utils/log"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"
)

func StartParseRoom() {
	var (
		loggingAvoider bool
		err            error
		data           = utils.RoomData{Status: utils.JsonStatus{
			Result: "offline",
			Reason: "",
		}}
		curRoomId string
	)
	// Initialize JSONByte
	JSONByte, err = json.Marshal(data)
	if err != nil {
		log.Logger.Error().Err(err).Send()
	}

	for {
		// Need to init in every loop
		data = utils.RoomData{Status: utils.JsonStatus{Result: "offline"}}
		room, b, err := InitParseRoom()
		if err != nil {
			log.Logger.Error().Err(err).Msg("could not read the room statistics.")
			data.Status = utils.JsonStatus{
				Result: "error",
				Reason: "wiimmfi-room-watcher failed to parse room statistics due to the software-side problem or the website's down.",
			}
			JSONByte, _ = json.Marshal(data)
			return
		}
		if !b {
			if !loggingAvoider {
				log.Logger.Info().Msg("Room not found. Seems the player is offline?")
				curRoomId = ""
			}
			JSONByte, err = json.Marshal(data)
			if err != nil {
				log.Logger.Error().Err(err).Send()
			}
			loggingAvoider = true
			time.Sleep(time.Duration(utils.LoadedConfig.Interval) * time.Second)
			continue
		}

		/// === Start to output room details === ///
		loggingAvoider = false
		/// Room name
		data.Id = room.RoomName

		/// Mode-related
		gameMode := utils.CheckGameMode(room.RaceMode)
		data.Setting.GameMode = gameMode
		// Store game mode but text for non-coder
		if _, b := utils.GAMEMODE[gameMode]; b {
			data.Setting.GameModeText = utils.GAMEMODE[gameMode]
		}
		if gameMode == utils.ModePrivateVS || gameMode == utils.ModeVS {
			if room.Engine >= 0 && room.Engine < len(utils.ENGINE) {
				data.Setting.EngineText = utils.ENGINE[room.Engine]
			} else {
				data.Setting.EngineText = "Unknown"
			}
			data.Setting.Engine = room.Engine
		}
		// Store amount of races in a room
		data.Setting.RaceCount = room.RaceCount

		/// Current track/arena
		data.Setting.Course = room.Track[1].(string)
		data.Setting.CourseId = int(room.Track[0].(float64))

		// Branch if Nintendo track or similar
		if v, ok := NinImagePath[data.Setting.CourseId]; ok {
			data.Setting.ThumbnailUrl = "https://mario.wiki.gallery/images/thumb/" + v
		} else {
			data.Setting.ThumbnailUrl = "https://ct.wiimm.de/img/start/" + strconv.Itoa(data.Setting.CourseId)
		}

		/// Players
		for _, player := range room.Members {
			member := utils.RoomMember{
				Pid:          player.Pid,
				FriendCode:   player.Fc,
				Name:         player.Name[0][0],
				RaceRating:   player.Ev,
				BattleRating: player.Eb,
				Status:       player.OlRole,
				FinishTimes:  []int{player.Time[0]},
				Course: utils.MemberCourse{
					Name:     player.Track[1].(string),
					Id:       int(player.Track[0].(float64)),
					Category: player.Track[4].(string),
				},
				DelayTime: player.Delay[1],
				ConnFail:  player.ConnFail,
			}
			chara := utils.CharacterId(player.Driver[0])
			vehicle := utils.VehicleId(player.Vehicle[0])
			member.Combos = []utils.Combo{{
				Character: utils.ComboChild{Id: int(chara), Name: chara.String()},
				Vehicle:   utils.ComboChild{Id: int(vehicle), Name: vehicle.String()},
			}}

			// If guest exists then store guest data
			if player.PlayerLen > 1 {
				member.GuestName = player.Name[1][0]
				member.FinishTimes = append(member.FinishTimes, player.Time[1])
				chara := utils.CharacterId(player.Driver[1])
				vehicle := utils.VehicleId(player.Vehicle[1])
				member.Combos = append(member.Combos, utils.Combo{
					Character: utils.ComboChild{Id: int(chara), Name: chara.String()},
					Vehicle:   utils.ComboChild{Id: int(vehicle), Name: vehicle.String()},
				})
			}
			data.Members = append(data.Members, member)
			data.MemberLen = room.Players
			data.Status = utils.JsonStatus{Result: "success"}
		}

		// Input encoded data into JSONByte and finally is readable via browser and websocket
		JSONByte, err = json.Marshal(data)
		if err != nil {
			log.Logger.Error().Err(err).Send()
		}

		if curRoomId != data.Id {
			log.Logger.Info().Msgf("Detected room joined: %s", data.Id)
			curRoomId = data.Id
		}

		log.Logger.Debug().
			Str("course", data.Setting.Course).
			Int("course_id", data.Setting.CourseId).
			Str("engine", data.Setting.EngineText).
			Str("game_mode", data.Setting.GameModeText).
			Send()
		log.Logger.Debug().Array("members", data.Members).Send()

		time.Sleep(time.Duration(utils.LoadedConfig.Interval) * time.Second)
	}
}

type SourceParse struct {
	RoomId    int                 `json:"room_id"`
	RoomName  string              `json:"room_name"`
	RaceMode  int                 `json:"race_mode"`
	Players   int                 `json:"n_players"`
	RaceCount int                 `json:"n_races"`
	Engine    int                 `json:"engine"`
	Track     [5]interface{}      `json:"track"`
	Members   []SourceMemberParse `json:"members"`
}

type SourceMemberParse struct {
	Pid       int            `json:"pid"`
	Fc        string         `json:"fc"`
	OlRole    string         `json:"ol_role"`
	Ev        int            `json:"ev"`
	Eb        int            `json:"eb"`
	PlayerLen int            `json:"n_players"`
	Name      [][]string     `json:"name"`
	Track     [5]interface{} `json:"track"`
	Driver    []int          `json:"driver"`
	Vehicle   []int          `json:"vehicle"`
	Delay     []int          `json:"delay"`
	Time      []int          `json:"time"`
	ConnFail  float32        `json:"conn_fail"`
}

func InitParseRoom() (*SourceParse, bool, error) {
	var base []json.RawMessage
	var result SourceParse

	res, err := http.Get("https://wiimmfi.de/stats/mkwx/room/p" + utils.LoadedConfig.Pid + "?m=json")
	if err != nil {
		return nil, false, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, false, err
	}
	//fmt.Printf("%v\n\n", string(body))
	if err := res.Body.Close(); err != nil {
		log.Logger.Debug().Err(err).Send()
	}

	// Unmarshal original output
	if err := json.Unmarshal(body, &base); err != nil {
		return nil, false, err
	}
	if len(base) < 4 {
		return nil, false, nil
	}
	// Unmarshal former parsed output
	if err := json.Unmarshal(base[2], &result); err != nil {
		return nil, false, err
	}

	return &result, true, nil
}
