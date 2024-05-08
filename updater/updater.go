//go:build !debug

package updater

import (
	"bufio"
	"fmt"
	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
	"log"
	"os"
	"path/filepath"
	"time"
)

const version, repoPath = "1.0.0", "expiteRz/wiimmfi-room-watcher"

func DoUpdate() {
	fmt.Println("Now checking the new version. This will take some time if you have bad network...")
	exec, err := os.Executable()
	if err != nil {
		log.SetPrefix("[Updater] ")
		log.Println("Failed to update application", err)
		time.Sleep(5 * time.Second)
		return
	}

	ver := semver.MustParse(version)
	update, err := selfupdate.UpdateSelf(ver, repoPath)
	if err != nil {
		fmt.Println("Failed to update application")
		return
	}

	if update.Version.EQ(ver) {
		fmt.Println("Application is already up-to-date!")
		full, _ := os.Executable()
		dir, file := filepath.Split(full)
		oldName := filepath.Join(dir, "."+file+".old")
		os.Remove(oldName)
		return
	}

	fmt.Println("Application updated to the latest version:", update.Version.String())
	fmt.Println("Release notes:\n", update.ReleaseNotes)
	fmt.Println("Press any key to restart the application automatically...")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	_, err = os.Open(exec)
	if err != nil {
		log.Fatalln("Failed to restart the application. Please execute it manually.")
		return
	}
	os.Exit(0)
}
