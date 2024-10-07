package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"
)

const (
	appName    = "anime-list"
	moduleName = "github.com/dongwlin/anime-list"
	binDir     = "bin"
)

func main() {
	var buildAt, goVersion, version string
	flag.StringVar(&buildAt, "buildAt", time.Now().Format(time.RFC3339), "Build time")
	flag.StringVar(&goVersion, "goVersion", runtime.Version(), "Go version used for build")
	flag.StringVar(&version, "version", "dev", "Application version")
	flag.Parse()

	if err := os.MkdirAll(binDir, 0755); err != nil {
		log.Fatalf("Failed to create bin directory: %v\n", err)
	}

	if err := buildProject(buildAt, goVersion, version); err != nil {
		log.Fatalf("Failed to build %s: %v\n", appName, err)
	}
}

func buildProject(buildAt, goVersion, version string) error {
	xBuildAt := fmt.Sprintf("-X '%s/cmd.buildAt=%s'", moduleName, buildAt)
	xGoVersion := fmt.Sprintf("-X '%s/cmd.goVersion=%s'", moduleName, goVersion)
	xVersion := fmt.Sprintf("-X '%s/cmd.version=%s'", moduleName, version)

	var ldflags string
	if version != "dev" {
		ldflags = fmt.Sprintf("-w -s %s %s %s", xBuildAt, xGoVersion, xVersion)
	} else {
		ldflags = fmt.Sprintf("%s %s %s", xBuildAt, xGoVersion, xVersion)
	}

	log.Printf("Start building for %s.\n", appName)

	outputPath := filepath.Join(".", binDir, appName)
	if runtime.GOOS == "windows" {
		outputPath += ".exe"
	}

	buildCommand := []string{
		"build",
		"-o", outputPath,
		"-ldflags", ldflags,
		"main.go",
	}

	cmd := exec.Command("go", buildCommand...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	startTime := time.Now()
	if err := cmd.Run(); err != nil {
		return err
	}
	duration := time.Since(startTime)

	log.Printf("Successfully built %s in %s.\n", appName, duration)
	return nil
}
