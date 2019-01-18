# Changelog

## 0.0.46-21 (2019/01/17)

* Upgrade to Discord PTB 0.0.46

## 0.0.45-20 (2019/01/10)

* Upgrade to Discord PTB 0.0.45

## 0.0.44-19 (2018/12/21)

* Upgrade to Discord PTB 0.0.44

## 0.0.43-18 (2018/10/24)

* Fix Windows desktop notifications
* Upgrade to Go 1.11
* Use [go mod](https://golang.org/cmd/go/#hdr-Module_maintenance) instead of `dep`

## 0.0.43-17 (2018/05/23)

* Upgrade to Discord PTB 0.0.43
* Update bootstrap script for multiInstance

## 0.0.42-16 (2018/04/20)

* Upgrade to Discord PTB 0.0.42

## 0.0.41-15 (2018/02/11)

* Move ia32/x64 to win32/win64 for arch def
* Remove nupkg file

## 0.0.41-14 (2018/02/11)

* Disable host updates (Issue portapps/discord-portable#10)

## 0.0.41-13 (2018/02/09)

* Ability to pass custom args to the portable process

## 0.0.41-12 (2018/01/13)

* Rebuild electron.asar to push data directly in `data` subfolder
* No need to override USERPROFILE environment variable anymore
* Allow multiple instances

> :warning: **UPGRADE NOTES**
> * Move everything in `data\AppData\Roaming\discordptb\*` to `data` folder and remove folder `data\AppData`.

## 0.0.41-11 (2018/01/12)

* Upgrade to Discord PTB 0.0.41

## 0.0.39-10 (2017/12/09)

* Upgrade to Discord PTB 0.0.39

## 0.0.38-9 (2017/11/22)

* Upgrade to Discord PTB 0.0.38
* Move app to a subfolder
* Reduce dependencies to avoid heuristic detection
* Add UPX compression

> :warning: **UPGRADE NOTES**
> * Delete all files and folders in root folder except `data` folder.

## 0.0.37-8 (2017/11/19)

* Upgrade to Discord PTB 0.0.37
* Move repository to [Portapps](https://github.com/portapps) org

## 0.0.36-7 (2017/11/11)

* Upgrade to Discord PTB 0.0.36
* Workaround for Discord PTB bug (Issue #2)

## 0.0.34-6 (2017/10/29)

* Add info in the README about BetterDiscord
* Switch to [Golang Dep](https://github.com/golang/dep) as dependency manager
* Upgrade to Go 1.9.1

## 0.0.34-5 (2017/09/05)

* New logger
* Override USERPROFILE env var instead of using symlink to APPDATA to store data
* Do not migrate old data folder from APPDATA
* Reduce dependencies and system calls to avoid heuristic detection

> :warning: **UPGRADE NOTES**
> * Move the content of `data\*` in `data\AppData\Roaming\discordptb\`
> * Remove symlink `%APPDATA%\discordptb`

## 0.0.34-3 (2017/08/26)

* Upgrade to Go 1.9
* Add support guidelines

## 0.0.34-2 (2017/08/09)

* Upgrade to Discord PTB 0.0.34

## 0.0.32-1 (2017/07/20)

* Initial version
