package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"os/exec"
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

	var appExec string;
	err, appExec = updater.Update(bundle, instDir, appName)
	if err != nil {
		logger.Fatal(err)
	}
	
	cmd := exec.Command(appExec)
	err = cmd.Start()
	if err != nil {
	    log.Fatal(err)
	}
}
