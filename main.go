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
	Papp.ID = "discord-ptb-portable"
	Papp.Name = "DiscordPTB"
	Init()
}

func main() {
	Papp.AppPath = AppPathJoin("app")
	Papp.DataPath = AppPathJoin("data")

	electronBinPath := PathJoin(Papp.AppPath, FindElectronAppFolder("app-", Papp.AppPath))

	Papp.Process = PathJoin(Papp.AppPath, "Update.exe")
	Papp.Args = []string{"--processStart", "DiscordPTB.exe"}
	Papp.WorkingDir = electronBinPath

	// Workaround for tray.png not found issue (https://github.com/portapps/discord-ptb-portable/issues/2)
	appdata.RestoreAssets(PathJoin(Papp.DataPath, "AppData", "Roaming"), "discord")

	Launch()
}
