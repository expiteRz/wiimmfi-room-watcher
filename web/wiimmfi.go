package web

import (
	"app.rz-public.xyz/wiimmfi-room-watcher/utils"
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

var RoomData utils.RoomData // expose room data variable to use in tui

func StartParseRoom() {
	var (
		loggingAvoider bool
		err            error
	)
	checkSelf := func(i, j int) string {
		if i == j {
			return ">"
		}
		return " "
	}
	// Initialize JSONByte
	JSONByte, err = json.Marshal(RoomData)
	if err != nil {
		log.SetPrefix("[Wiimmfi] ")
		log.Println(err)
	}

	for {
		RoomData = utils.RoomData{Status: "offline"}
		res, err := http.Get(fmt.Sprintf("https://wiimmfi.de/stats/mkwx/room/p%d?m=json", utils.LoadedConfig.Pid))
		if err != nil {
			log.SetPrefix("[Wiimmfi] ")
			log.Fatalf("Failed to get connection with wiimmfi. Please report it to the program owner: %v\n", err)
			return
		}

		body, _ := io.ReadAll(res.Body)
		//fmt.Printf("%v\n\n", string(body))
		res.Body.Close()

		var roomData interface{}
		if err := json.Unmarshal(body, &roomData); err != nil {
			log.SetPrefix("[Wiimmfi] ")
			log.Printf("Error occurred: %v", err)
			time.Sleep(time.Duration(utils.LoadedConfig.Interval) * time.Second)
			continue
		}

		if len(roomData.([]interface{})) < 4 {
			if !loggingAvoider {
				fmt.Println("Room not found. Seems the player is offline?")
			}
			JSONByte, err = json.Marshal(RoomData)
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
		fmt.Printf("=== Room: %s ===\n", roomData.([]interface{})[2].(map[string]interface{})["room_name"])
		RoomData.Id = roomData.([]interface{})[2].(map[string]interface{})["room_name"].(string)

		/// Mode-related
		gameMode := utils.CheckGameMode(cast.ToInt(roomData.([]interface{})[2].(map[string]interface{})["ol_status"].([]interface{})[0]))
		RoomData.Setting.GameMode = gameMode
		switch gameMode {
		case utils.ModePrivateVS, utils.ModeVS:
			fmt.Printf("Engine: %s\n", utils.ENGINE[int(roomData.([]interface{})[2].(map[string]interface{})["engine"].(float64))])
			RoomData.Setting.Engine = int(roomData.([]interface{})[2].(map[string]interface{})["engine"].(float64))
		case utils.ModePrivateBalloonBattle, utils.ModeBalloonBattle:
			fmt.Println("Balloon Battle")
		case utils.ModePrivateCoinBattle, utils.ModeCoinBattle:
			fmt.Println("Coin Battle")
		}

		/// Current track/arena
		fmt.Printf("Track: %s\n", roomData.([]interface{})[2].(map[string]interface{})["track"].([]interface{})[1])
		RoomData.Setting.Course = roomData.([]interface{})[2].(map[string]interface{})["track"].([]interface{})[1].(string)
		RoomData.Setting.CourseId = cast.ToInt(roomData.([]interface{})[2].(map[string]interface{})["track"].([]interface{})[0])
		/// Players
		players := roomData.([]interface{})[2].(map[string]interface{})["members"].([]interface{})
		for _, player := range players {
			fmt.Printf(
				"%s %-15s   %-20s   %4sVR\n",
				checkSelf(utils.LoadedConfig.Pid, cast.ToInt(player.(map[string]interface{})["pid"])),
				player.(map[string]interface{})["fc"],
				player.(map[string]interface{})["name"].([]interface{})[0].([]interface{})[0],
				strconv.FormatFloat(player.(map[string]interface{})["ev"].(float64), 'f', 0, 64),
			)
			member := utils.RoomMember{
				Pid:          cast.ToInt(player.(map[string]interface{})["pid"]),
				FriendCode:   player.(map[string]interface{})["fc"].(string),
				Name:         player.(map[string]interface{})["name"].([]interface{})[0].([]interface{})[0].(string),
				RaceRating:   cast.ToInt(player.(map[string]interface{})["ev"]),
				BattleRating: cast.ToInt(player.(map[string]interface{})["eb"]),
				Status:       player.(map[string]interface{})["ol_role"].(string),
				FinishTimes:  []int{cast.ToInt(player.(map[string]interface{})["time"].([]interface{})[0])},
			}
			// If guest exists then print guest name
			if player.(map[string]interface{})["name"].([]interface{})[1].([]interface{})[0] != nil {
				fmt.Printf(
					"%s %-15s   %-20s\n", checkSelf(utils.LoadedConfig.Pid, cast.ToInt(player.(map[string]interface{})["pid"])), "", player.(map[string]interface{})["name"].([]interface{})[1].([]interface{})[0],
				)
				member.GuestName = player.(map[string]interface{})["name"].([]interface{})[1].([]interface{})[0].(string)
				member.FinishTimes = append(member.FinishTimes, cast.ToInt(player.(map[string]interface{})["time"].([]interface{})[1]))
			}
			RoomData.Members = append(RoomData.Members, member)
			RoomData.Status = "success"
			RoomData.MemberLen = cast.ToInt(roomData.([]interface{})[2].(map[string]interface{})["n_players"])

			JSONByte, err = json.Marshal(RoomData)
			if err != nil {
				log.SetPrefix("[Wiimmfi] ")
				log.Println(err)
			}
		}
		fmt.Println("")

		time.Sleep(time.Duration(utils.LoadedConfig.Interval) * time.Second)
	}
}
