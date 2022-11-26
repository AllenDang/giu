package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Usage: gmdeploy [os] [icon] [path/to/project]")
		fmt.Println("Flags:")
		flag.PrintDefaults()

		os.Exit(0)
	}

	var targetOS string

	flag.StringVar(&targetOS, "os", runtime.GOOS, "target deploy os [windows, darwin, linux]")

	var iconPath string

	flag.StringVar(&iconPath, "icon", "", "application icon file path")

	var upx bool

	flag.BoolVar(&upx, "upx", false, "use upx to compress executable")

	var createUniversalBinaryForMacOS bool

	flag.BoolVar(&createUniversalBinaryForMacOS, "ub", false, "create universal binary for macOS")

	flag.Parse()

	if targetOS == "" {
		targetOS = runtime.GOOS
	}

	projectPath, _ := os.Getwd()
	appName := filepath.Base(projectPath)

	// Prepare build dir
	outputDir := filepath.Join(projectPath, "build", targetOS)
	if err := os.RemoveAll(outputDir); err != nil {
		log.Fatalf("error removing content of %s", outputDir)
	}

	mkdirAll(outputDir)

	switch targetOS {
	case "darwin":
		const iconExtension = ".icns"
		//nolint:gosec // Compile: cannot fix
		cmd := exec.Command("bash", "-c", fmt.Sprintf("go build -ldflags='-s -w' -o %s", appName))
		cmd.Dir = projectPath
		runCmd(cmd)

		// Create universal binary for macOS
		// Build for arm64
		//nolint:gosec // cannot fix
		cmd = exec.Command("bash", "-c", fmt.Sprintf("CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 go build -o %s_binary_arm64 -ldflags '-s -w' .", appName))
		cmd.Dir = projectPath
		runCmd(cmd)

		// Build for amd64
		//nolint:gosec // cannot fix
		cmd = exec.Command("bash", "-c", fmt.Sprintf("CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -o %s_binary_amd64 -ldflags '-s -w' .", appName))
		runCmd(cmd)

		// Merge them together with lipo
		//nolint:gosec // cannot fix
		cmd = exec.Command("bash", "-c", fmt.Sprintf("lipo -create -output %[1]s %[1]s_binary_amd64 %[1]s_binary_arm64", appName))
		runCmd(cmd)

		// Clean up
		_ = os.Remove(filepath.Join(projectPath, fmt.Sprintf("%s_binary_arm64", appName)))
		_ = os.Remove(filepath.Join(projectPath, fmt.Sprintf("%s_binary_amd64", appName)))

		// Upx
		if upx {
			//nolint:gosec // cannot fix
			cmd = exec.Command("upx", appName)
			runCmd(cmd)
		}

		// Bundle
		macOSPath := filepath.Join(outputDir, fmt.Sprintf("%s.app", appName), "Contents", "MacOS")
		mkdirAll(macOSPath)

		// Copy compiled executable to build folder
		//nolint:gosec // cannot fix
		cmd = exec.Command("mv", appName, macOSPath)
		runCmd(cmd)

		// Prepare Info.plist
		contentsPath := filepath.Join(outputDir, fmt.Sprintf("%s.app", appName), "Contents")
		save(filepath.Join(contentsPath, "Info.plist"), darwinPlist(appName))

		// Prepare PkgInfo
		save(filepath.Join(contentsPath, "PkgInfo"), darwinPkginfo())

		if len(iconPath) > 0 && filepath.Ext(iconPath) == iconExtension {
			// Prepare icon
			resourcesPath := filepath.Join(contentsPath, "Resources")
			mkdirAll(resourcesPath)

			// Rename icon file name to [appName].icns
			//nolint:gosec // cannot fix
			cmd = exec.Command("cp", iconPath, filepath.Join(resourcesPath, fmt.Sprintf("%s%s", appName, iconExtension)))
			runCmd(cmd)
		}

		fmt.Printf("%s.app is generated at %s/build/%s/\n", appName, projectPath, targetOS)
	case "linux":
		//nolint:gosec // Compile: cannot fix
		cmd := exec.Command("bash", "-c", fmt.Sprintf("go build -ldflags='-s -w' -o %s", appName))
		cmd.Dir = projectPath
		runCmd(cmd)

		// Bundle
		contentsPath := filepath.Join(outputDir, fmt.Sprintf("%s.app", appName))
		binPath := filepath.Join(contentsPath, "bin")
		mkdirAll(binPath)

		// Copy compiled executable to build folder
		//nolint:gosec // rename command - cannot be fixed
		cmd = exec.Command("mv", appName, binPath)
		runCmd(cmd)

		// create desktop entry
		hasIcon := iconPath != "" && filepath.Ext(iconPath) == ".icns"

		desktopPath := filepath.Join(contentsPath, "share", "applications")
		mkdirAll(desktopPath)

		save(filepath.Join(desktopPath, fmt.Sprintf("%s.desktop", appName)), linuxDesktop(appName, hasIcon))

		if hasIcon {
			// Prepare icon
			iconsPath := filepath.Join(contentsPath, "share", "icons")
			mkdirAll(iconsPath)

			// Rename icon file name to [appName].icns
			newIconName := filepath.Join(iconsPath, fmt.Sprintf("%s.icns", appName))
			//nolint:gosec // cp command - cannot fix
			cmd = exec.Command("cp", iconPath, newIconName)
			runCmd(cmd)
		}
	default:
		fmt.Printf("Sorry, %s is not supported yet.\n", targetOS)
	}
}
