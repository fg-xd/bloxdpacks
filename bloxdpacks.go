package main

import (
	"fmt"
	"os"
	"os/exec"
)

var AppName string = "BloxdPacks"
var Version string = "BETA-0.0.1"

var packs = map[string]string{
	"test": "https://bloxdpacks.fgxd.int.yt/p/test.txz",
}

func main() {
	arg := os.Args
	if len(arg) < 2 {
		fmt.Println("BloxdPacks - CLI Written in Go")
		fmt.Println("type bloxdpacks --help to see commands!")
		os.Exit(0)
	} else if arg[1] == "install" || arg[1] == "i" {
		userinputPack := arg[2]
		if _, exists := packs[userinputPack]; exists {
			installCmd := exec.Command("curl", "-fsSL", packs[userinputPack])
			installErr := installCmd.Run()
			if installErr != nil {
				fmt.Println("Error in downloading texture pack:", installErr)
				os.Exit(1)
			} else {
				fmt.Println("Intalling Texture Pack...")
				installCmd.Run()
				fmt.Println("Done! Now run 'bloxdpacks extract [texture-pack-name] to extract it into a folder!")
				os.Exit(0)
			}
		}
	} else if arg[1] == "extract" || arg[1] == "e" {
		userinputExtract := arg[2]
		extractCmd := exec.Command("tar", "-xJf", userinputExtract+".txz")
		extractErr := extractCmd.Run()
		if extractErr != nil {
			fmt.Println("Error occured while extracting:", extractErr)
			os.Exit(1)
		} else {
			extractCmd.Run()
			os.Exit(0)
		}
	} else if arg[1] == "--help" || arg[1] == "-h" {
		fmt.Println("===BloxdPacks Help Menu===")
		fmt.Println("bloxdpacks [command / alternative] [options] -> Description")
		fmt.Println("bloxdpacks --help / -h                       -> Help Menu")
		fmt.Println("bloxdpacks install / i [texture pack name]   -> Install Texture Pack")
		fmt.Println("bloxdpacks extract / e [texture pack name]   -> Extract Texture Pack")
		fmt.Println("bloxdpacks --version / -v                    -> Current Version")
		os.Exit(0)
	} else if arg[1] == "--version" || arg[1] == "-v" {
		fmt.Println("BloxdPacks Version: " + Version)
		os.Exit(0)
	}
}
