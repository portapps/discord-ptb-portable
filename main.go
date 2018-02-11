//go:generate go get -v github.com/jteeuwen/go-bindata/go-bindata/...
//go:generate go-bindata -pkg appdata -o src/appdata/appdata.go data/...
//go:generate go get -v github.com/josephspurrier/goversioninfo/...
//go:generate goversioninfo -icon=res/papp.ico
package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/portapps/discord-ptb-portable/src/appdata"
	. "github.com/portapps/portapps"
)

func init() {
	Papp.ID = "discord-ptb-portable"
	Papp.Name = "DiscordPTB"
	Init()
}

func main() {
	Papp.AppPath = AppPathJoin("app")
	Papp.DataPath = AppPathJoin("data")

	electronBinPath := PathJoin(Papp.AppPath, FindElectronAppFolder("app-", Papp.AppPath))

	Papp.Process = PathJoin(electronBinPath, "DiscordPTB.exe")
	Papp.WorkingDir = electronBinPath

	// Update settings
	settingsPath := PathJoin(Papp.DataPath, "settings.json")
	if _, err := os.Stat(settingsPath); err == nil {
		Log.Info("Update settings...")
		rawSettings, err := ioutil.ReadFile(settingsPath)
		if err == nil {
			jsonMapSettings := make(map[string]interface{})
			json.Unmarshal(rawSettings, &jsonMapSettings)
			Log.Info("Current settings:", jsonMapSettings)

			jsonMapSettings["SKIP_HOST_UPDATE"] = true
			Log.Info("New settings:", jsonMapSettings)

			jsonSettings, err := json.Marshal(jsonMapSettings)
			if err != nil {
				Log.Error("Settings marshal:", err)
			}
			err = ioutil.WriteFile(settingsPath, jsonSettings, 0644)
			if err != nil {
				Log.Error("Write settings:", err)
			}
		}
	}

	// Workaround for tray.png not found issue (https://github.com/portapps/discord-ptb-portable/issues/2)
	appdata.RestoreAssets(Papp.Path, "data")

	Launch(os.Args[1:])
}
