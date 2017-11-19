//go:generate go get -v github.com/jteeuwen/go-bindata/go-bindata/...
//go:generate go-bindata -pkg appdata -o src/appdata/appdata.go -prefix "data/" data/...
//go:generate go get -v github.com/josephspurrier/goversioninfo/...
//go:generate goversioninfo -icon=res/papp.ico
package main

import (
	"github.com/portapps/discord-ptb-portable/src/appdata"
	. "github.com/portapps/portapps"
)

func init() {
	App.Id = "discord-ptb-portable"
	App.Name = "DiscordPTB"
	Init()
}

func main() {
	App.MainPath = FindElectronMainFolder("app-")
	App.DataPath = CreateFolder(PathJoin(App.RootDataPath, "AppData", "Roaming", "discord"))
	App.Process = RootPathJoin("Update.exe")
	App.Args = []string{"--processStart", "DiscordPTB.exe"}
	App.WorkingDir = App.MainPath

	// Workaround for tray.png not found issue (https://github.com/portapps/discord-ptb-portable/issues/2)
	appdata.RestoreAssets(PathJoin(App.RootDataPath, "AppData", "Roaming"), "discord")

	OverrideUserprofilePath(App.RootDataPath)
	Launch()
}
