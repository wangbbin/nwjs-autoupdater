package updater

import (
	"os"
	"path/filepath"

	"github.com/mholt/archiver"
)

func Update(bundle, instDir, appName string) (error, string) {

	// bundle: C:\Users\Even\AppData\Local\Temp\vnpcUpdate\update.zip
	// instDir: C:\Users\Even\Documents\VNPCv2\nw_vnpc\build\VNPCv2\win32
	// appName: VNPCv2


///////////////////////////////////
	extractDir := "./files";

	appExecName := appName + ".exe"
  appExec := filepath.Join(instDir, appExecName)

  err := archiver.Zip.Open(bundle, extractDir)

	err = filepath.Walk(extractDir, func(path string, f os.FileInfo, err error) error {
			if(!f.IsDir()) {
				newFile := path
				fileToReplace := instDir + "\\" + f.Name()
				fileToReplaceBackup := fileToReplace + ".bak"

				err := os.Rename(fileToReplace, fileToReplaceBackup)
				if err != nil {
					return err
				}
				err = os.Rename(newFile, fileToReplace)
				if err != nil {
					return err
				}
				os.Remove(fileToReplaceBackup)
			}
			return nil
	})


	if err != nil {
		return err, appExec
	}

  return nil, appExec
}
