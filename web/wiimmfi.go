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

func StartParseRoom() {
	var (
		loggingAvoider bool
		err            error
		data           utils.RoomData
	)
	checkSelf := func(i, j int) string {
		if i == j {
			return ">"
		}
		return " "
	}
	// Initialize JSONByte
	JSONByte, err = json.Marshal(data)
	if err != nil {
		log.Println(err)
	}

	for {
		data = utils.RoomData{Status: "offline"}
		res, err := http.Get(fmt.Sprintf("https://wiimmfi.de/stats/mkwx/room/p%d?m=json", utils.LoadedConfig.Pid))
		if err != nil {
			log.Fatalf("Failed to get connection with wiimmfi. Please report it to the program owner: %v\n", err)
			return
		}

		body, _ := io.ReadAll(res.Body)
		//fmt.Printf("%v\n\n", string(body))
		res.Body.Close()

		var test interface{}
		if err := json.Unmarshal(body, &test); err != nil {
			log.Printf("Error occurred: %v", err)
			time.Sleep(time.Duration(utils.LoadedConfig.Interval) * time.Second)
			continue
		}

		if len(test.([]interface{})) < 4 {
			if !loggingAvoider {
				fmt.Println("Room not found. Seems the player is offline?")
			}
			JSONByte, err = json.Marshal(data)
			if err != nil {
				log.Println(err)
			}
			loggingAvoider = true
			time.Sleep(time.Duration(utils.LoadedConfig.Interval) * time.Second)
			continue
		}

		/// === Start to output room details === ///
		loggingAvoider = false
		/// Room name
		fmt.Printf("=== Room: %s ===\n", test.([]interface{})[2].(map[string]interface{})["room_name"])
		data.Id = test.([]interface{})[2].(map[string]interface{})["room_name"].(string)

		/// Mode-related
		gameMode := utils.CheckGameMode(int(test.([]interface{})[2].(map[string]interface{})["ol_status"].([]interface{})[0].(float64)))
		data.Setting.GameMode = gameMode
		switch gameMode {
		case utils.ModePrivateVS:
		case utils.ModeVS:
			fmt.Printf("Engine: %s\n", utils.ENGINE[int(test.([]interface{})[2].(map[string]interface{})["engine"].(float64))])
			data.Setting.Engine = int(test.([]interface{})[2].(map[string]interface{})["engine"].(float64))
		case utils.ModePrivateBalloonBattle:
		case utils.ModeBalloonBattle:
			fmt.Println("Balloon Battle")
		case utils.ModePrivateCoinBattle:
		case utils.ModeCoinBattle:
			fmt.Println("Coin Battle")
		}

		/// Current track/arena
		fmt.Printf("Track: %s\n", test.([]interface{})[2].(map[string]interface{})["track"].([]interface{})[1])
		data.Setting.Course = test.([]interface{})[2].(map[string]interface{})["track"].([]interface{})[1].(string)
		data.Setting.CourseId = cast.ToInt(test.([]interface{})[2].(map[string]interface{})["track"].([]interface{})[0])
		/// Players
		players := test.([]interface{})[2].(map[string]interface{})["members"].([]interface{})
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
			}
			// If guest exists then print guest name
			if player.(map[string]interface{})["name"].([]interface{})[1].([]interface{})[0] != nil {
				fmt.Printf("%s %-15s   %-20s\n", checkSelf(utils.LoadedConfig.Pid, cast.ToInt(player.(map[string]interface{})["pid"])), "", player.(map[string]interface{})["name"].([]interface{})[1].([]interface{})[0])
				member.GuestName = player.(map[string]interface{})["name"].([]interface{})[1].([]interface{})[0].(string)
			}
			data.Members = append(data.Members, member)
			data.Status = "success"

			JSONByte, err = json.Marshal(data)
			if err != nil {
				log.Println(err)
			}
		}
		fmt.Println("")

		time.Sleep(time.Duration(utils.LoadedConfig.Interval) * time.Second)
	}
}
