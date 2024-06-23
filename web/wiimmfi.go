package web

import (
	"app.rz-public.xyz/wiimmfi-room-watcher/utils"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

func StartParseRoom() {
	var (
		loggingAvoider bool
		err            error
		data           utils.RoomData
	)
	// Initialize JSONByte
	JSONByte, err = json.Marshal(data)
	if err != nil {
		log.SetPrefix("[Wiimmfi] ")
		log.Println(err)
	}

	for {
		data = utils.RoomData{Status: "offline"}
		room, b, err := InitParseRoom()
		if err != nil {
			log.SetPrefix("[Wiimmfi] ")
			log.Fatalf("Failed to get connection with wiimmfi. Please report it to the program owner: %v\n", err)
			return
		}
		if !b {
			if !loggingAvoider {
				fmt.Println("Room not found. Seems the player is offline?")
			}
			JSONByte, err = json.Marshal(data)
			if err != nil {
				log.SetPrefix("[Wiimmfi] ")
				log.Println(err)
			}
			loggingAvoider = true
			time.Sleep(time.Duration(utils.LoadedConfig.Interval) * time.Second)
			continue
		}

		/// === Start to output room details === ///
		loggingAvoider = false
		/// Room name
		fmt.Printf("=== Room: %s ===\n", room.RoomName)
		data.Id = room.RoomName

		/// Mode-related
		gameMode := utils.CheckGameMode(int(room.OlStatus[0].(float64)))
		data.Setting.GameMode = gameMode
		// Store game mode but text for non-coder
		if _, b := utils.GAMEMODE[gameMode]; b {
			data.Setting.GameModeText = utils.GAMEMODE[gameMode]
		}
		switch gameMode {
		case utils.ModePrivateVS, utils.ModeVS:
			fmt.Println("Engine:", utils.ENGINE[room.Engine])
			data.Setting.EngineText = utils.ENGINE[room.Engine]
			data.Setting.Engine = room.Engine
		case utils.ModePrivateBalloonBattle, utils.ModeBalloonBattle:
			fmt.Println("Balloon Battle")
		case utils.ModePrivateCoinBattle, utils.ModeCoinBattle:
			fmt.Println("Coin Battle")
		default:
			break
		}
		// Store amount of races in a room
		data.Setting.RaceCount = room.RaceCount

		/// Current track/arena
		fmt.Println("Track:", room.Track[1].(string))
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
			fmt.Printf(
				"%s %-15s   %-20s   %4sVR\n",
				checkSelf(utils.LoadedConfig.Pid, player.Pid),
				player.Fc,
				player.Name[0][0],
				strconv.FormatInt(int64(player.Ev), 10),
			)
			member := utils.RoomMember{
				Pid:          player.Pid,
				FriendCode:   player.Fc,
				Name:         player.Name[0][0],
				RaceRating:   player.Ev,
				BattleRating: player.Eb,
				Status:       player.OlRole,
				FinishTimes:  []int{player.Time[0]},
				Course: utils.MemberCourse{
					Name:    player.Track[1].(string),
					Id:      int(player.Track[0].(float64)),
					Allowed: player.Track[4].(string),
				},
			}
			chara := utils.CharacterId(player.Driver[0])
			vehicle := utils.VehicleId(player.Vehicle[0])
			member.Combos = []utils.Combo{{
				Character: utils.ComboChild{Id: int(chara), Name: chara.String()},
				Vehicle:   utils.ComboChild{Id: int(vehicle), Name: vehicle.String()},
			}}

			// If guest exists then store guest data
			if player.PlayerLen > 1 {
				fmt.Printf(
					"%s %-15s   %-20s\n", checkSelf(utils.LoadedConfig.Pid, player.Pid), "", player.Name[1][0],
				)
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
			data.Status = "success"
		}

		// Input encoded data into JSONByte and finally is readable via browser and websocket
		JSONByte, err = json.Marshal(data)
		if err != nil {
			log.SetPrefix("[Wiimmfi] ")
			log.Println(err)
		}
		fmt.Println("")

		time.Sleep(time.Duration(utils.LoadedConfig.Interval) * time.Second)
		fmt.Println("\033[H\033[2J")
	}
}

func checkSelf(i string, j int) string {
	x, err := strconv.Atoi(i)
	if err != nil {
		log.Println(err)
		return " "
	}
	if x == j {
		return ">"
	}
	return " "
}

type SourceParse struct {
	RoomId    int                 `json:"room_id"`
	RoomName  string              `json:"room_name"`
	OlStatus  []interface{}       `json:"ol_status"`
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
	Time      []int          `json:"time"`
}

func InitParseRoom() (*SourceParse, bool, error) {
	var base []json.RawMessage
	var result SourceParse

	res, err := http.Get("https://wiimmfi.de/stats/mkwx/room/p" + utils.LoadedConfig.Pid + "?m=json")
	if err != nil {
		return nil, false, err
	}
	body, _ := io.ReadAll(res.Body)
	//fmt.Printf("%v\n\n", string(body))
	res.Body.Close()

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
