package updater

import (
	"os"
	"io"
	"path/filepath"
	"github.com/mholt/archiver"
	"github.com/skratchdot/open-golang/open"
)

func Update(bundle, instDir, appName string) (error, string) {
	extractDir := "./files";
	if _, err := os.Stat(extractDir); err == nil {
		os.RemoveAll(extractDir)
	}

	appExecName := appName + ".exe"
	appExec := filepath.Join(instDir, appExecName)

	err := archiver.Zip.Open(bundle, extractDir)
	err = filepath.Walk(extractDir, func(path string, f os.FileInfo, err error) error {
		if(!f.IsDir()) {
			extractedFile := path
			// Remove "./files/" from path
			relExtractPath, err := filepath.Rel(extractDir, extractedFile)	
			// Remove filename from path
			relExtractDir := filepath.Dir(relExtractPath)

			// Installation sub-directory for the file
			instDirSubdir := filepath.Join(instDir, relExtractDir)
			if err != nil {
				return err
			}
			
			// Full installation path of the new file
			newFileInstPath := filepath.Join(instDir, relExtractPath)

			// Make sure the subdirectory/subdirectories (if any) for the new file exist 
			if _, err = os.Stat(instDirSubdir); os.IsNotExist(err) {
				os.MkdirAll(instDirSubdir, 0777)
			}

			src, err := os.Open(extractedFile)
			if err != nil {
				return err
			}
			defer src.Close()

			dst, err := os.OpenFile(newFileInstPath, os.O_WRONLY|os.O_CREATE, 0644)
			if err != nil {
				return err
			}
			defer dst.Close()

			// Move the extracted file to instdir
			_, err = io.Copy(dst, src)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		start(appExec)
		return err, appExec
	}

	err = os.RemoveAll(extractDir)
	if err != nil {
		start(appExec)
		return err, appExec
	}
	err = os.RemoveAll(bundle)
	if err != nil {
		start(appExec)
		return err, appExec
	}

	start(appExec)
	
	return nil, appExec
}

func start(appExec string) {
	open.Start(appExec)
}
