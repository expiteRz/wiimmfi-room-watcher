package test

import (
	"app.rz-public.xyz/wiimmfi-room-watcher/utils"
	"app.rz-public.xyz/wiimmfi-room-watcher/web"
	"testing"
)

func TestParseRoom(t *testing.T) {
	utils.LoadedConfig.Pid = 600613674
	room, _, err := web.InitParseRoom()
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", *room)
}
