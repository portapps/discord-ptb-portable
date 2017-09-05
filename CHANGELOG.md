# Changelog

## 0.0.34-4 (2017/09/05)

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
