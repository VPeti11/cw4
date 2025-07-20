package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

// Localization
var useHungarian = false

// This one speaks for itself
var enableColor = false

// Package manager varriable
var manager = detectOS()

// Global reset code
const reset = "\033[0m"

// All supported ANSI colors
var colors = []string{
	"\033[31m", // Red
	"\033[33m", // Yellow
	"\033[32m", // Green
	"\033[36m", // Cyan
	"\033[34m", // Blue
	"\033[35m", // Magenta
	"\033[91m", // Bright Red
	"\033[92m", // Bright Green
	"\033[93m", // Bright Yellow
	"\033[94m", // Bright Blue
	"\033[95m", // Bright Magenta
	"\033[96m", // Bright Cyan
}

func tr(en, hu string) string {
	if useHungarian {
		return hu
	}
	return en
}

func clearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		fmt.Print("\033[H\033[2J")
	}
}

func pause() {
	printFancy(tr("Press ENTER to continue...", "Nyomj ENTER-t a folytatáshoz..."))
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func detectOS() string {
	if runtime.GOOS == "windows" {
		return "choco"
	} else {
		return "flatpak"
	}

}

func installPackage(manager, pkg string) {
	var cmd string
	switch manager {
	case "choco":
		cmd = fmt.Sprintf("choco install %s -y", pkg)
	case "flatpak":
		cmd = fmt.Sprintf("flatpak install %s -y", pkg)
	default:
		printFancy(tr("Unsupported OS or missing package manager.", "Nem támogatott rendszer vagy hiányzó csomagkezelő."))
		return
	}
	printFancy(tr("Installing:", "Telepítés:") + " " + pkg)
	printFancy(tr("Running command:", "Futtatott parancs:") + " " + cmd)
	exec.Command("sh", "-c", cmd).Run()
}

func packageDict(pkg string, osvar string) string {
	switch pkg {
	case "firefox":
		switch osvar {
		case "choco":
			return "firefox"
		default:
			return "org.mozilla.firefox"
		}
	case "chrome":
		switch osvar {
		case "choco":
			return "googlechrome"
		default:
			return "com.google.chrome"
		}
	case "libreoffice":
		switch osvar {
		case "choco":
			return "libreoffice"
		default:
			return "org.libreoffice.LibreOffice"
		}
	case "onlyoffice":
		switch osvar {
		case "choco":
			return "libreoffice"
		default:
			return "org.onlyoffice.desktopeditors"
		}
	case "steam":
		switch osvar {
		case "choco":
			return "steam"
		default:
			return "com.valvesoftware.Steam"
		}
	case "vlc":
		switch osvar {
		case "choco":
			return "vlc"
		default:
			return "org.videolan.VLC"
		}
	case "unzip":
		switch osvar {
		case "choco":
			return "winrar"
		default:
			return "io.github.peazip.peazip"
		}
	case "discord":
		switch osvar {
		case "choco":
			return "discord"
		default:
			return "com.discordapp.Discord"
		}
	case "vs":
		switch osvar {
		case "choco":
			return "vscode"
		default:
			return "com.visualstudio.code"
		}
	case "librewolf":
		switch osvar {
		case "choco":
			return "librewolf"
		default:
			return "io.gitlab.librewolf-community"
		}
	case "brave":
		switch osvar {
		case "choco":
			return "brave"
		default:
			return "com.brave.Browser"
		}
	default:
		return ""
	}
}

func browserSelect() {
	clearScreen()
	printName()
	printFancy(tr("Install Web Browsers", "Böngészők telepítése"))
	printFancy(tr("Do you want to install Chrome, Firefox, Brave or librewolf?", "Chrome, Firefox, Brave vagy librewolf-ot szeretne telepíteni?"))
	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)
	choice = strings.ToLower(choice)
	switch choice {
	case "chrome":
		installPackage(manager, packageDict("chrome", detectOS()))
	case "firefox":
		installPackage(manager, packageDict("firefox", detectOS()))
	case "brave":
		installPackage(manager, packageDict("brave", detectOS()))
	case "librewolf":
		installPackage(manager, packageDict("librewolf", detectOS()))
	default:
		printFancy(tr("Invalid input.", "Érvénytelen bemenet."))

	}

}

func installMenu() {
	reader := bufio.NewReader(os.Stdin)

	if manager == "" {
		printFancy(tr("No supported package manager found.", "Nem található támogatott csomagkezelő."))
		return
	}
	for {
		clearScreen()
		printName()
		printFancy("1. " + tr("Install Web Browser", "Böngésző telepítése"))
		printFancy("2. " + tr("Install Media Tools", "Média eszközök telepítése"))
		printFancy("3. " + tr("Install Dev Tools", "Fejlesztői eszközök telepítése"))
		printFancy("4. " + tr("Install basic utilities", "Alap programok telepítése"))
		printFancy("5. " + tr("Install gaming software", "Játék programok telepítése"))
		printFancy("6. " + tr("Install office", "Office telepítése"))
		printFancy("7. " + tr("Install other package", "Más csomag telepítése telepítése"))
		printFancy("0. " + tr("Exit", "Kilépés"))
		printFancyInline(tr("Choose an option: ", "Válassz egy lehetőséget: "))

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			browserSelect()
			pause()
		case "2":
			clearScreen()
			printName()
			printFancy("2. " + tr("Install Media Tools", "Média eszközök telepítése"))
			installPackage(manager, packageDict("vlc", detectOS()))
			installPackage(manager, "ffmpeg")
			pause()
		case "3":
			clearScreen()
			printName()
			printFancy("3. " + tr("Install Dev Tools", "Fejlesztői eszközök telepítése"))
			installPackage(manager, "git")
			installPackage(manager, packageDict("vs", detectOS()))
			pause()
		case "4":
			clearScreen()
			printName()
			printFancy("4. " + tr("Install basic utilities", "Alap programok telepítése"))
			installPackage(manager, packageDict("unzip", detectOS()))
			installPackage(manager, packageDict("vlc", detectOS()))
			browserSelect()
			pause()
		case "5":
			clearScreen()
			printName()
			printFancy("5. " + tr("Install gaming software", "Játék programok telepítése"))
			installPackage(manager, packageDict("steam", detectOS()))
			installPackage(manager, packageDict("discord", detectOS()))
			pause()
		case "6":
			clearScreen()
			printName()
			printFancy(tr("Install office", "Office telepítése"))
			printFancy(tr("Do you want to install LibreOffice or Onlyoffice?", "Libreoffice-t vagy Onlyoffice-t szeretne telepíteni?"))
			reader := bufio.NewReader(os.Stdin)
			choice, _ := reader.ReadString('\n')
			choice = strings.TrimSpace(choice)
			choice = strings.ToLower(choice)
			switch choice {
			case "libreoffice":
				installPackage(manager, packageDict("libreoffice", detectOS()))
			case "onlyoffice":
				installPackage(manager, packageDict("onlyoffice", detectOS()))
			default:
				printFancy(tr("Invalid input.", "Érvénytelen bemenet."))
				pause()
			}
		case "7":
			clearScreen()
			printName()
			printFancy(tr("What package do you want to install?", "Milyen csomagot szeretne telepíteni?"))
			reader := bufio.NewReader(os.Stdin)
			choice, _ := reader.ReadString('\n')
			choice = strings.TrimSpace(choice)
			choice = strings.ToLower(choice)
			installPackage(manager, choice)
		case "0":
			os.Exit(0)
		default:
			printFancy(tr("Invalid input.", "Érvénytelen bemenet."))
			pause()
		}
	}
}

func printName() {
	for _, ch := range tr("==== CW4 ====", "==== CW4 Menü ====") {
		fmt.Print(randColor() + string(ch))
	}
	fmt.Println(reset)
}

func devMenu() {
	reader := bufio.NewReader(os.Stdin)
	for {
		clearScreen()
		for _, ch := range tr("CW4 Menu", "CW4 Menü") {
			fmt.Print(randColor() + string(ch))
		}
		fmt.Println(reset)
		printFancy("1. " + tr("Run Installer Menu", "Telepítő menü futtatása"))
		printFancy("2. " + tr("Change Language", "Nyelv módosítása"))
		printFancy("3. " + tr("Color mode", "Színes mód"))
		printFancy("4. " + tr("Run cleanup script", "Tisztító indítása"))
		printFancy("5. " + tr("Run debloat script", "Debloat szkript indítása"))
		printFancy("0. " + tr("Exit", "Kilépés"))
		printFancyInline(tr("Choose an option: ", "Válassz egy lehetőséget: "))

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			installMenu()
		case "2":
			useHungarian = !useHungarian
		case "3":
			enableColor = !enableColor
		case "4":
			cleanupWindows()
			pause()
		case "5":
			removeAppxPackages()
			pause()
		case "0":
			os.Exit(0)
		default:
			printFancy(tr("Invalid input.", "Érvénytelen bemenet."))
			pause()
		}
	}
}

func showCredits() {
	pause := func() {
		printFancy("\n" + tr("Press Enter to continue...", "Nyomj Entert a folytatáshoz..."))
		fmt.Scanln()
	}

	messages := []string{
		tr("==== Clean-Wise version 4 (I2) ====", "==== Clean-Wise 4-es verzió (I2) ===="),
		tr("👤 Created by VPeti", "👤 Készítette: VPeti"),
		tr("🔨 Build number: v1.0", "🔨 Build szám: v1.0"),
		tr("📜 Licensed under GPLv3", "📜 GPLv3 licenc alatt"),
		tr("🌐 Open-source system cleaner and setup program for Windows & Linux", "🌐 Nyílt forráskódú rendszer takarító és telepítő Windowsra és Linuxra"),
		tr("📦 Features temp cleanup, system tune-up, Flatpak and Choco support, and multilingual UI", "📦 Tartalmaz ideiglenes fájlok törlését, rendszer gyorsítást, Flatpak és Choco támogatást és többnyelvű felületet"),
	}

	for _, line := range messages {
		clearScreen()
		printFancy(line)
		time.Sleep(1200 * time.Millisecond)
	}

	pause()
	main()
}

func removeAppxPackages() {
	if detectOS() != "choco" {
		printFancy(tr("This feature is Windows-only.", "Ez a funkció csak Windows rendszeren működik."))
		return
	}

	script := `
$AppXApps = @(
"*Microsoft.BingNews*", "*Microsoft.GetHelp*", "*Microsoft.Getstarted*",
"*Microsoft.Messaging*", "*Microsoft.Microsoft3DViewer*",
"*Microsoft.MicrosoftOfficeHub*", "*Microsoft.MicrosoftSolitaireCollection*",
"*Microsoft.NetworkSpeedTest*", "*Microsoft.Office.Sway*", "*Microsoft.OneConnect*",
"*Microsoft.People*", "*Microsoft.Print3D*", "*Microsoft.SkypeApp*", "*Microsoft.WindowsAlarms*",
"*Microsoft.WindowsCamera*", "*microsoft.windowscommunicationsapps*", "*Microsoft.WindowsFeedbackHub*",
"*Microsoft.WindowsMaps*", "*Microsoft.WindowsSoundRecorder*", "*Microsoft.Xbox.TCUI*",
"*Microsoft.XboxApp*", "*Microsoft.XboxGameOverlay*", "*Microsoft.XboxIdentityProvider*",
"*Microsoft.XboxSpeechToTextOverlay*", "*Microsoft.ZuneMusic*", "*Microsoft.ZuneVideo*",
"*EclipseManager*", "*ActiproSoftwareLLC*", "*AdobeSystemsIncorporated.AdobePhotoshopExpress*",
"*Duolingo-LearnLanguagesforFree*", "*PandoraMediaInc*", "*CandyCrush*",
"*Wunderlist*", "*Flipboard*", "*Twitter*", "*Facebook*", "*Spotify*"
)

foreach ($App in $AppXApps) {
    Get-AppxPackage -Name $App | Remove-AppxPackage -ErrorAction SilentlyContinue
    Get-AppxPackage -Name $App -AllUsers | Remove-AppxPackage -AllUsers -ErrorAction SilentlyContinue
    Get-AppxProvisionedPackage -Online | Where-Object DisplayName -like $App | Remove-AppxProvisionedPackage -Online -ErrorAction SilentlyContinue
}

$WhitelistedApps = 'Microsoft.Paint3D|Microsoft.WindowsCalculator|Microsoft.WindowsStore|Microsoft.Windows.Photos|CanonicalGroupLimited.UbuntuonWindows|Microsoft.XboxGameCallableUI|Microsoft.XboxGamingOverlay|Microsoft.Xbox.TCUI|Microsoft.XboxGamingOverlay|Microsoft.XboxIdentityProvider|Microsoft.MicrosoftStickyNotes|Microsoft.MSPaint*'

Get-AppxPackage -AllUsers | Where-Object {$_.Name -NotMatch $WhitelistedApps} | Remove-AppxPackage
Get-AppxPackage | Where-Object {$_.Name -NotMatch $WhitelistedApps} | Remove-AppxPackage
Get-AppxProvisionedPackage -Online | Where-Object {$_.PackageName -NotMatch $WhitelistedApps} | Remove-AppxProvisionedPackage -Online
`

	cmd := exec.Command("powershell", "-NoProfile", "-ExecutionPolicy", "Bypass", "-Command", script)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	printFancy(tr("Running AppX cleanup...", "AppX takarítás folyamatban..."))
	err := cmd.Run()
	if err != nil {
		printFancy(tr("Error during cleanup: ", "Hiba a takarítás során: ") + err.Error())
		return
	}
	printFancy(tr("AppX cleanup complete.", "AppX takarítás befejezve."))
}

func printFancyInline(args ...interface{}) {
	text := fmt.Sprint(args...)

	if !enableColor {
		fmt.Print(text)
		return
	}

	for _, ch := range text {
		fmt.Print(randColor() + string(ch))
	}
	fmt.Print("\033[0m") // Reset color but no newline
}

func randColor() string {
	return colors[rand.Intn(len(colors))]
}

func min(a, b, c int) int {
	if a < b && a < c {
		return a
	} else if b < c {
		return b
	}
	return c
}

func flatpakInstall() {
	if _, err := exec.LookPath("flatpak"); err == nil {
		printFancy(tr("Flatpak is already installed.", "Flatpak már telepítve van."))
		return
	}

	printFancy(tr("Flatpak not found. Attempting to install...", "Flatpak nem található. Telepítés folyamatban..."))

	pmCommands := map[string][]string{
		"apt":    {"sudo", "apt", "install", "-y", "flatpak"},
		"dnf":    {"sudo", "dnf", "install", "-y", "flatpak"},
		"pacman": {"sudo", "pacman", "-S", "--noconfirm", "flatpak"},
	}

	installed := false

	for pm, cmd := range pmCommands {
		if _, err := exec.LookPath(pm); err == nil {
			printFancy(tr("Using "+pm+" to install Flatpak...", "A Flatpak telepítése "+pm+" segítségével..."))
			install := exec.Command(cmd[0], cmd[1:]...)
			install.Stdout = nil
			install.Stderr = nil
			err := install.Run()
			if err != nil {
				printFancy(tr("Installation with "+pm+" failed: "+err.Error(), "A Flatpak telepítése "+pm+" segítségével sikertelen: "+err.Error()))
			} else {
				installed = true
				break
			}
		}
	}

	if !installed {
		printFancy(tr("No supported package manager found or installation failed.", "Nem található támogatott csomagkezelő, vagy a telepítés sikertelen."))
		return
	}

	printFancy(tr("Adding Flathub repository...", "Flathub tároló hozzáadása..."))
	addRepo := exec.Command("flatpak", "remote-add", "--if-not-exists", "flathub", "https://dl.flathub.org/repo/flathub.flatpakrepo")
	err := addRepo.Run()
	if err != nil {
		printFancy(tr("Failed to add Flathub: "+err.Error(), "Flathub hozzáadása sikertelen: "+err.Error()))
	} else {
		printFancy(tr("Flathub repository added successfully.", "Flathub tároló sikeresen hozzáadva."))
	}
}

func depinstall() {
	chocoPath := "C:\\ProgramData\\chocolatey\\bin\\choco.exe"

	if _, err := os.Stat(chocoPath); os.IsNotExist(err) {
		printFancy(tr("No Chocolatey detected.", "A Chocolatey nem található."))
		time.Sleep(1 * time.Second)

		clearScreen()
		pause()

		printFancy(tr("Installing Chocolatey.", "Chocolatey telepítése."))
		printFancy(tr("DO NOT DISCONNECT FROM THE INTERNET OR CLOSE THIS WINDOW!", "NE SZAKÍTSD MEG AZ INTERNETKAPCSOLATOT ÉS NE ZÁRD BE AZ ABLAKOT!"))

		installCmd := `Set-ExecutionPolicy Bypass -Scope Process -Force; ` +
			`[System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; ` +
			`iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))`

		cmd := exec.Command("powershell", "-NoProfile", "-Command", installCmd)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err = cmd.Run()
		if err != nil {
			printFancy(tr("Chocolatey install failed: "+err.Error(), "Chocolatey telepítése sikertelen: "+err.Error()))
			return
		}

		printFancy(tr("Dependency install complete.", "A függőségek telepítése befejeződött."))
		printFancy(tr("Press Enter to continue...", "Nyomj Entert a folytatáshoz..."))
		fmt.Scanln()
	}
}

func printFancy(args ...interface{}) {
	text := fmt.Sprint(args...)

	if !enableColor {
		fmt.Print(text + "\n") // Replace printFancy with raw print
		return
	}

	for _, ch := range text {
		fmt.Print(randColor() + string(ch))
	}
	fmt.Print("\033[0m\n")
}

func cleanupWindows() {
	if detectOS() != "choco" {
		fmt.Println("This cleanup tool is only for Windows.")
		return
	}

	err := exec.Command("openfiles").Run()
	if err != nil {
		fmt.Println("You must run this program as Administrator.")
		fmt.Println("Right-click and select 'Run as Administrator'.")
		fmt.Println("Press Enter to exit...")
		fmt.Scanln()
		os.Exit(1)
	}
	fmt.Println("Running with admin privileges...")

	clearScreen()

	commands := [][]string{
		{"cmd", "/C", "del", "/s", "/f", "/q", "%WinDir%\\Temp\\*.*"},
		{"cmd", "/C", "del", "/s", "/f", "/q", "%WinDir%\\Prefetch\\*.*"},
		{"cmd", "/C", "del", "/s", "/f", "/q", "%Temp%\\*.*"},
		{"cmd", "/C", "del", "/s", "/f", "/q", "%AppData%\\Temp\\*.*"},
		{"cmd", "/C", "del", "/s", "/f", "/q", "%HomePath%\\AppData\\LocalLow\\Temp\\*.*"},
		{"cmd", "/C", "rd", "/s", "/q", "%WinDir%\\Temp"},
		{"cmd", "/C", "rd", "/s", "/q", "%WinDir%\\Prefetch"},
		{"cmd", "/C", "rd", "/s", "/q", "%Temp%"},
		{"cmd", "/C", "rd", "/s", "/q", "%AppData%\\Temp"},
		{"cmd", "/C", "rd", "/s", "/q", "%HomePath%\\AppData\\LocalLow\\Temp"},
		{"cmd", "/C", "md", "%WinDir%\\Temp"},
		{"cmd", "/C", "md", "%WinDir%\\Prefetch"},
		{"cmd", "/C", "md", "%Temp%"},
		{"cmd", "/C", "md", "%AppData%\\Temp"},
		{"cmd", "/C", "md", "%HomePath%\\AppData\\LocalLow\\Temp"},
		{"cmd", "/C", "TASKKILL", "/f", "/im", "explorer.exe"},
		{"cmd", "/C", "del", "/s", "/f", "/q", "%userprofile%\\AppData\\Local\\IconCache.db"},
		{"cmd", "/C", "del", "/s", "/f", "/q", "%userprofile%\\AppData\\Local\\Microsoft\\Windows\\Explorer\\iconcache*.*"},
		{"cmd", "/C", "del", "/s", "/f", "/q", "%userprofile%\\AppData\\Local\\Microsoft\\Windows\\Explorer\\thumbcache*.*"},
		{"cmd", "/C", "START", "explorer.exe"},
	}

	for _, cmd := range commands {
		_ = exec.Command(cmd[0], cmd[1:]...).Run()
	}

	fmt.Println("\nCleaning done.")
	fmt.Println("Press Enter to exit...")
	fmt.Scanln()
}

func main() {
	clearScreen()
	rand.Seed(time.Now().UnixNano())
	cLines := []string{
		" ____      ",
		"/\\  _`\\    ",
		"\\ \\ \\/\\_\\  ",
		" \\ \\ \\/_/_ ",
		"  \\ \\ \\L\\ \\",
		"   \\ \\____/",
		"    \\/___/ ",
		"           ",
	}

	wLines := []string{
		" __      __    ",
		"/\\ \\  __/\\ \\   ",
		"\\ \\ \\/\\ \\ \\ \\  ",
		" \\ \\ \\ \\ \\ \\ \\ ",
		"  \\ \\ \\_/ \\_\\ \\",
		"   \\ `\\___x___/",
		"    '\\/__//__/ ",
	}

	fourLines := []string{
		" __ __      ",
		"/\\ \\\\ \\     ",
		"\\ \\ \\\\ \\    ",
		" \\ \\ \\\\ \\_  ",
		"  \\ \\__ ,__\\",
		"   \\/_/\\_\\_/",
		"      \\/_/  ",
	}

	maxLines := min(len(cLines), len(wLines), len(fourLines))
	for i := 0; i < maxLines; i++ {
		fmt.Println(
			randColor() + cLines[i] + " " +
				randColor() + wLines[i] + " " +
				randColor() + fourLines[i] +
				reset)
	}
	fmt.Print("\n")

	for _, ch := range tr("By VPeti", "Készítette VPeti") {
		fmt.Print(randColor() + string(ch))
	}
	fmt.Println(reset)

	time.Sleep(2 * time.Second)

	switch manager {
	case "choco":
		depinstall()
	default:
		flatpakInstall()
	}

	devMenu()
}
