# CW4

**CW4 (CleanWise 4 or Newer Clean-Wise version 4)** is a multiplatform, multilingual CLI tool written in Go that acts as a **modular installer, system setup assistant, and lightweight cleaner** for Windows and Linux.

It's not flashy. It's not fancy. It just works (usually). Designed with specific use cases in mind—like setting up Flatpak or Choco environments, debloating a fresh Windows install, or quickly installing Firefox or Steam—it does the one job you need without asking "Are you sure?" five times.

## Devnote

Don’t ask why this exists. I needed it. It made sense in my head. And dont ask about the name. Its based on an old project. And its not CW-I. That was based on the same old project simply called CW 2. And that was a version of the original CW thats turning 8 years old on the day i release this (20/07/2025). If you want to find that than good luck since its lost media. I remember using it and later making CW 2 since i couldnt find the original. And the date was when i first used it. I know because i have written it down.

---

## Features

* ❯ Interactive CLI menus for installing browsers, office suites, gaming tools, media utilities, and more
* ❯ Works on both Linux and Windows (via Flatpak or Chocolatey)
* ❯ Multilingual support (English and Hungarian)
* ❯ Optional colorized output using ANSI sequences. Dont ask. It looks cool
* ❯ Built-in Flatpak bootstrap for distros with missing support
* ❯ Windows AppX bloatware removal with  automation
* ❯ Visual Studio Code, LibreWolf, Brave, Steam, Discord, VLC, and many more available
* ❯ Installer system identifies OS and configures accordingly
* ❯ Open Source and GPLv3 Licensed

---

## Installing CW4

CW4 is a portable executable. You can build it yourself or download binaries for your OS.

### Linux

To build and run:

```
go build -o cw4 main.go
./cw4
```

If you’re using Flatpak or want to set it up:

```
./cw4
```

### Windows

Run `cw4.exe` in a terminal. Chocolatey is required and will be auto-installed if not found.

Make sure  is enabled, and you're running as administrator for full functionality (especially for AppX cleanup).

---

## Compiling the Source

Make sure you have Go installed and configured. Then run:

```
go build -o cw4 main.go
```

On Windows:

```
go build -o cw4.exe main.go
```

This builds the latest version of CW4.

---

## Usage

Run the program:

```
./cw4
```

You’ll be greeted with a dynamic and colorful menu (if color mode is enabled). From here, you can:

* Launch the main installer menu
* Change the language
* Enable/disable color output
* Run cleanup (Windows)
* Run AppX debloat (Windows)

---

### Menu Options

| Option                   | Description                                  |
| ------------------------ | -------------------------------------------- |
| `1` Install Web Browser  | Install Firefox, Chrome, LibreWolf, or Brave |
| `2` Install Media Tools  | Install VLC, ffmpeg                          |
| `3` Install Dev Tools    | Install Git and VS Code                      |
| `4` Install Utilities    | PeaZip/WinRAR, VLC, and browser of choice    |
| `5` Install Gaming Stuff | Steam, Discord                               |
| `6` Install Office       | LibreOffice or OnlyOffice                    |
| `7` Install Custom       | Input your own package name and install it   |
| `0` Exit                 | Quit the program                             |

All options support both Chocolatey and Flatpak, and the tool auto-detects which to use.

---

## AppX Debloat (Windows)

CW4 includes a built-in script to remove Microsoft bloatware from Windows:

* Removes most preinstalled apps (News, Skype, 3D Viewer, etc.)
* Keeps essentials like Calculator, Paint, Store
* Can be run manually from the menu

You **must run as Administrator** for this feature.

---

## Flatpak Bootstrap (Linux)

If Flatpak is not found, CW4 will:

* Detect your package manager (apt, dnf, pacman)
* Attempt to install Flatpak
* Add the Flathub repository
* Return you to the installer

---

## Package Compatibility Table

| Package     | Chocolatey     | Flatpak Equivalent              |
| ----------- | -------------- | ------------------------------- |
| Firefox     | `firefox`      | `org.mozilla.firefox`           |
| Chrome      | `googlechrome` | `com.google.chrome`             |
| LibreOffice | `libreoffice`  | `org.libreoffice.LibreOffice`   |
| OnlyOffice  | `libreoffice`  | `org.onlyoffice.desktopeditors` |
| VLC         | `vlc`          | `org.videolan.VLC`              |
| Steam       | `steam`        | `com.valvesoftware.Steam`       |
| Discord     | `discord`      | `com.discordapp.Discord`        |
| VS Code     | `vscode`       | `com.visualstudio.code`         |
| LibreWolf   | `librewolf`    | `io.gitlab.librewolf-community` |
| Brave       | `brave`        | `com.brave.Browser`             |

---

## Localization

CW4 supports **English and Hungarian** via the `useHungarian` variable. This can be toggled anytime via the main menu.

---

## Color Output

Enable or disable color mode from the menu. CW4 uses a randomized ANSI color scheme to print titles and prompts.

---

## Dependencies

* **Go** (for compiling the source)
* **Flatpak** (for Linux installs, handled by CW4)
* **Chocolatey** (for Windows installs, auto-installed by CW4)

---

## License

CW4 is licensed under the **GNU GPLv3**.

| Right           | Description                                       |
| --------------- | ------------------------------------------------- |
| Use             | Use the program freely                            |
| Modify          | Study and alter the source code                   |
| Share           | Distribute original or modified versions          |
| Source Required | Source code must accompany distributed binaries   |
| No DRM          | You can’t lock it down with hardware restrictions |
| Patent-Safe     | You’re protected against software patents         |

See [LICENSE.md](LICENSE.md) for full license terms.

The documentation and README are licensed under the [GFDL](fdl.md)

---


## Final Words

CW4 isn’t made for mass adoption. It’s made for **fast prototyping**, **power users**, and **autistic perfectionists** (like me) who just want their setup ready without bloat.

Use at your own risk. Or don’t.
Either way, enjoy the color explosions.