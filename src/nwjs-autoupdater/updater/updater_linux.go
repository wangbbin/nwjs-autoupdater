package updater

import (
	"os"
	"path/filepath"

	"github.com/mholt/archiver"
)

func Update(bundle, instDir, appName string) (error, string) {

	extractDir := "./files";

	appExecName := appName
  	appExec := filepath.Join(instDir, appExecName)

  	err := archiver.Zip.Open(bundle, extractDir)

	err = filepath.Walk(extractDir, func(path string, f os.FileInfo, err error) error {
		if(!f.IsDir()) {
			newFile := path
			fileToReplace := instDir + "/" + f.Name()
			fileToReplaceBackup := fileToReplace + ".bak"
			bakFileCreated := false

			// Append ".bak" on file to be replaced (if it exist)
			if _, err := os.Stat(fileToReplace); err == nil {
				err := os.Rename(fileToReplace, fileToReplaceBackup)
				bakFileCreated = true
				if err != nil {
					return err
				}
			}

			// Move the new file to replace the old file
			err = os.Rename(newFile, fileToReplace)
			if err != nil {
				return err
			}

			if(bakFileCreated) {
				os.Remove(fileToReplaceBackup)
			}
		}
		return nil
	})


	if err != nil {
		return err, appExec
	}

  return nil, appExec
}
