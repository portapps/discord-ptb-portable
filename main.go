//go:generate go get -v github.com/jteeuwen/go-bindata/go-bindata/...
//go:generate go-bindata -pkg appdata -o src/appdata/appdata.go -prefix "data/" data/...
//go:generate go get -v github.com/josephspurrier/goversioninfo/...
//go:generate goversioninfo -icon=res/app-portable.ico

package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/crazy-max/discord-ptb-portable/src/appdata"
	"github.com/google/logger"
)

const (
	NAME            = "discord-ptb-portable"
	APP_NAME        = "DiscordPTB"
	APP_DATA_FOLDER = "discordptb"
	APP_PROCESS     = "DiscordPTB.exe"
)

func main() {
	// Current path
	currentPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		logger.Error("Current path:", err)
	}

	// Log file
	logfile, err := os.OpenFile(pathJoin(currentPath, NAME+".log"), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		logger.Error("Log file:", err)
	}
	defer logfile.Close()

	// Init logger
	logger.Init(NAME, false, false, logfile)
	logger.Info("--------")
	logger.Infof("Starting %s...", NAME)
	logger.Infof("Current path: %s", currentPath)

	// Find app folder
	logger.Infof("Lookup app folder in: %s", currentPath)
	appPath := ""
	rootFiles, _ := ioutil.ReadDir(currentPath)
	for _, f := range rootFiles {
		if strings.HasPrefix(f.Name(), "app-") && f.IsDir() {
			logger.Infof("App folder found: %s", f.Name())
			appPath = pathJoin(currentPath, f.Name())
			break
		}
	}
	if _, err := os.Stat(appPath); err == nil {
		logger.Infof("App path: %s", appPath)
	} else {
		logger.Error("App path does not exist")
	}

	// Init vars
	appExe := pathJoin(currentPath, "Update.exe")
	dataPath := pathJoin(currentPath, "data")
	dataAppPath := pathJoin(dataPath, "AppData", "Roaming", APP_DATA_FOLDER)
	logger.Infof("App executable: %s", appExe)
	logger.Infof("Data path: %s", dataPath)

	// Create data folder
	logger.Infof("Create data folder %s...", dataAppPath)
	err = os.MkdirAll(dataAppPath, 777)
	if err != nil {
		logger.Error("Create data folder:", err)
	}

	// Override USERPROFILE env var
	if err := os.Setenv("USERPROFILE", dataPath); err != nil {
		logger.Error("Cannot set USERPROFILE env var:", err)
	}

	// Workaround for tray.png (see #2)
	appdata.RestoreAssets(pathJoin(dataPath, "AppData", "Roaming"), "discord")

	// Launch
	logger.Infof("Launch %s...", APP_NAME)
	execApp := exec.Command(appExe, "--processStart", APP_PROCESS)
	execApp.Dir = appPath

	defer logfile.Close()
	execApp.Stdout = logfile
	execApp.Stderr = logfile

	if err := execApp.Start(); err != nil {
		logger.Error("Cmd Start:", err)
	}

	execApp.Wait()
}

func pathJoin(elem ...string) string {
	for i, e := range elem {
		if e != "" {
			return strings.Join(elem[i:], `\`)
		}
	}
	return ""
}
