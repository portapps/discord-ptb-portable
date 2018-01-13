<p align="center"><a href="https://github.com/portapps/discord-ptb-portable" target="_blank"><img width="100" src="https://github.com/portapps/discord-ptb-portable/blob/master/res/papp.png"></a></p>

<p align="center">
  <a href="https://github.com/portapps/discord-ptb-portable/releases/latest"><img src="https://img.shields.io/github/release/portapps/discord-ptb-portable.svg?style=flat-square" alt="GitHub release"></a>
  <a href="https://github.com/portapps/discord-ptb-portable/releases/latest"><img src="https://img.shields.io/github/downloads/portapps/discord-ptb-portable/total.svg?style=flat-square" alt="Total downloads"></a>
  <a href="https://ci.appveyor.com/project/crazy-max/discord-ptb-portable"><img src="https://img.shields.io/appveyor/ci/crazy-max/discord-ptb-portable.svg?style=flat-square" alt="AppVeyor"></a>
  <a href="https://goreportcard.com/report/github.com/portapps/discord-ptb-portable"><img src="https://goreportcard.com/badge/github.com/portapps/discord-ptb-portable?style=flat-square" alt="Go Report"></a>
  <a href="https://www.codacy.com/app/crazy-max/discord-ptb-portable"><img src="https://img.shields.io/codacy/grade/8556c9e756164889b0307dbc7282ef0a.svg?style=flat-square" alt="Code Quality"></a>
  <a href="https://saythanks.io/to/crazymax"><img src="https://img.shields.io/badge/thank-crazymax-426aa5.svg?style=flat-square" alt="Say Thanks"></a>
  <a href="https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=WQD7AQGPDEPSG"><img src="https://img.shields.io/badge/donate-paypal-7057ff.svg?style=flat-square" alt="Donate Paypal"></a>
</p>

## About

[DiscordPTB (Public Test Build)](https://discordapp.com) portable app made with 🚀 [Portapps](https://github.com/portapps).<br />
Tested on Windows 7, Windows 8.1 and Windows 10.

> 💡 Stable portable version of Discord is available [here](https://github.com/portapps/discord-portable)

## Installation

There are different kinds of artifacts :

* `discord-ptb-portable-{ia32,x64}-x.x.x-x-setup.exe` : Full portable release of Discord as a setup. **Recommended way**!
* `discord-ptb-portable-{ia32,x64}-x.x.x-x.7z` : Full portable release of Discord as a 7z archive.
* `discord-ptb-portable-{ia32,x64}.exe` : Only the portable binary (must be renamed `discord-ptb-portable.exe`)
* `DiscordPTBSetup-{ia32,x64}-x.x.x.exe` : The original setup from the [official website](https://discordapp.com/download).
* `DiscordPTB-{ia32,x64}-x.x.x-full.nupkg` : The original NUPKG file extracted from the original setup.

### Fresh installation

Install `discord-ptb-portable-{ia32,x64}-x.x.x-x-setup.exe` where you want then run `discord-ptb-portable.exe`.

### App already installed

If you have already installed DiscordPTB from the original setup, do the same thing as a fresh installation and :

* Move data located in `%APPDATA%\discordptb\*` to `data` folder.

Run `discord-ptb-portable.exe` and then you can [remove](https://support.microsoft.com/en-us/instantanswers/ce7ba88b-4e95-4354-b807-35732db36c4d/repair-or-remove-programs) DiscordPTB from your computer.

### Upgrade

For an upgrade, simply download and install the [latest setup](https://github.com/portapps/discord-ptb-portable/releases/latest).

### BetterDiscord

DiscordPTB Portable is also compatible with [BetterDiscord](https://betterdiscord.net).<br />
Download the latest release of BetterDiscord, launch the setup wizard, choose the install location `discord-ptb-portable\app\app-x.x.x` and press `Install` :

```
Deleting old cached files
Deleting temp path
Downloading Resource: BetterDiscord.zip
Extracting BetterDiscord
node_modules doesn't exist, creating
Extracting app.asar
Moving BetterDiscord to resources\node_modules\
Spicing index
Writing index
Finished installation, verifying installation...
Finished installing BetterDiscord with 0 errors
```

## How can i help ?

We welcome all kinds of contributions :raised_hands:!<br />
The most basic way to show your support is to star :star2: the project, or to raise issues :speech_balloon:<br />
Any funds donated will be used to help further development on this project! :gift_heart:

[![Donate Paypal](https://raw.githubusercontent.com/portapps/portapps/master/res/paypal.png)](https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=WQD7AQGPDEPSG)

## License

MIT. See `LICENSE` for more details.<br />
Rocket icon credit to [Squid Ink](http://thesquid.ink).
