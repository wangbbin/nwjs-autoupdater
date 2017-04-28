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
				extractedFile := path
				relExtractDir, err := filepath.Rel(extractDir, path)	// remove "./files/" from path
				instDirSubdir := filepath.Join(instDir, relExtractDir)	// installation sub-directory for the file
				if err != nil {
					return err
				}
				
				oldFileToReplace := instDir + "\\" + relExtractDir
				
				
				// Make sure the subdirectory/subdirectories (if any) for the new file exist 
				if _, err = os.Stat(instDirSubdir); os.IsNotExist(err) {
					os.MkdirAll(instDirSubdir, 0777)
				}

				// If the extracted file exist in the installation, rename it to end with .bak
				oldFileBackup := ""
				if _, err = os.Stat(oldFileToReplace); err == nil {
					oldFileBackup = oldFileToReplace + ".bak"
					err = os.Rename(oldFileToReplace, oldFileBackup)
					if err != nil {
						return err
					}
				}

				// Move the extracted file to instdir
				err = os.Rename(extractedFile, oldFileToReplace)
				if err != nil {
					return err
				}

				// Remove the .bak-file
				if oldFileBackup != "" {
					os.Remove(oldFileBackup)
				}
			}
			return nil
	})


	if err != nil {
		return err, appExec
	}

  return nil, appExec
}
