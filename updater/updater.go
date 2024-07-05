//go:build !debug

package updater

import (
	"app.rz-public.xyz/wiimmfi-room-watcher/utils/log"
	"bufio"
	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
	"os"
	"path/filepath"
)

const version, repoPath = "1.0.0", "expiteRz/wiimmfi-room-watcher"

func DoUpdate() {
	log.Logger.Info().Msg("Now checking the new version. This will take some time if you have bad network...")
	exec, err := os.Executable()
	if err != nil {
		log.Logger.Error().Err(err).Msg("Failed to update application")
		return
	}

	ver := semver.MustParse(version)
	update, err := selfupdate.UpdateSelf(ver, repoPath)
	if err != nil {
		log.Logger.Error().Err(err).Msg("Failed to update application")
		return
	}

	if update.Version.EQ(ver) {
		log.Logger.Info().Msg("Application is already up-to-date!")
		full, _ := os.Executable()
		dir, file := filepath.Split(full)
		oldName := filepath.Join(dir, "."+file+".old")
		if err := os.Remove(oldName); err != nil {
			log.Logger.Error().Err(err).Msg("")
		}
		return
	}

	log.Logger.Info().Msgf("Application updated to the latest version: %s", update.Version.String())
	log.Logger.Info().Msgf("Release notes:\n%s", update.ReleaseNotes)
	log.Logger.Info().Msg("Press any key to restart the application automatically...")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	_, err = os.Open(exec)
	if err != nil {
		log.Logger.Error().Err(err).Msg("Failed to restart the application. Please execute it manually.")
		return
	}
	os.Exit(0)
}
