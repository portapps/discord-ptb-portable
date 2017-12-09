# Changelog

## 0.0.39-10 (2017/12/09)

* New release of Discord PTB : 0.0.39

## 0.0.38-9 (2017/11/22)

* New release of Discord PTB : 0.0.38
* Move app to a subfolder
* Reduce dependencies to avoid heuristic detection
* Add UPX compression

> :warning: **UPGRADE NOTES**
> * Delete all files and folders in root folder except `data` folder.

## 0.0.37-8 (2017/11/19)

* New release of Discord PTB : 0.0.37
* Move repository to [Portapps](https://github.com/portapps) org

## 0.0.36-7 (2017/11/11)

* New release of Discord PTB : 0.0.36
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

* New release of Discord PTB : 0.0.34

## 0.0.32-1 (2017/07/20)

* Initial version
