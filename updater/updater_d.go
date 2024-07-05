//go:build debug

package updater

import "app.rz-public.xyz/wiimmfi-room-watcher/utils/log"

func DoUpdate() {
	log.Logger.Debug().Msg("You launched the app on development purpose. Updating software will be ignored.")
}
