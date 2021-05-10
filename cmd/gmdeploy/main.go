package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Usage: gmdeploy [os] [icon] [path/to/progject]")
		fmt.Println("Flags:")
		flag.PrintDefaults()

		os.Exit(0)
	}

	var targetOS string
	flag.StringVar(&targetOS, "os", "", "target deploy os [windows, darwin, linux]")

	var iconPath string
	flag.StringVar(&iconPath, "icon", "", "applicatio icon file path")

	flag.Parse()

	if len(targetOS) == 0 {
		targetOS = runtime.GOOS
	}

	projectPath, _ := os.Getwd()
	appName := filepath.Base(projectPath)

	// Prepare build dir
	outputDir := filepath.Join(projectPath, "build", targetOS)
	os.RemoveAll(outputDir)
	MkdirAll(outputDir)

	switch targetOS {
	case "darwin":
		// Compile
		cmd := exec.Command("bash", "-c", fmt.Sprintf("go build -ldflags='-s -w' -o %s", appName))
		cmd.Dir = projectPath
		RunCmd(cmd)

		// Upx
		// cmd = exec.Command("upx", appName)
		// RunCmd(cmd)

		// Bundle
		macOSPath := filepath.Join(outputDir, fmt.Sprintf("%s.app", appName), "Contents", "MacOS")
		MkdirAll(macOSPath)

		// Copy compiled executable to build folder
		cmd = exec.Command("mv", appName, macOSPath)
		RunCmd(cmd)

		// Prepare Info.plist
		contentsPath := filepath.Join(outputDir, fmt.Sprintf("%s.app", appName), "Contents")
		Save(filepath.Join(contentsPath, "Info.plist"), darwinPlist(appName))

		// Prepare PkgInfo
		Save(filepath.Join(contentsPath, "PkgInfo"), darwinPkginfo())

		if len(iconPath) > 0 && filepath.Ext(iconPath) == ".icns" {
			// Prepare icon
			resourcesPath := filepath.Join(contentsPath, "Resources")
			MkdirAll(resourcesPath)

			// Rename icon file name to [appName].icns
			cmd = exec.Command("cp", iconPath, filepath.Join(resourcesPath, fmt.Sprintf("%s.icns", appName)))
			RunCmd(cmd)
		}

		fmt.Printf("%s.app is generated at %s/build/%s/\n", appName, projectPath, targetOS)

	default:
		fmt.Printf("Sorry, %s is not supported yet.\n", targetOS)
	}
}
