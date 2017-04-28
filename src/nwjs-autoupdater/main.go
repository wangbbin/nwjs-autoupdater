package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"nwjs-autoupdater/updater"
)

func main() {
	var bundle, instDir, appName string

	flag.StringVar(&bundle, "bundle", "", "Path to the update package")
	flag.StringVar(&instDir, "inst-dir", "", "Path to the application install dir")
	flag.StringVar(&appName, "app-name", "", "Name of the app (with extension)")
	flag.Parse()


	cwd, _ := os.Getwd()
	logfile, err := os.Create(filepath.Join(cwd, "updater.log"))
	if err != nil {
		panic(err)
	}
	defer logfile.Close()

	logger := log.New(logfile, "", log.LstdFlags)


	err, _ = updater.Update(bundle, instDir, appName)
	if err != nil {
		logger.Fatal(err)
	}
}
