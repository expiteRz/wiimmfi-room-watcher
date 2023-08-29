//go:build debug

package updater

import "log"

func DoUpdate() {
	log.SetPrefix("[Updater] ")
	log.Println("You launched the app on debug. Updater will not start")
}
