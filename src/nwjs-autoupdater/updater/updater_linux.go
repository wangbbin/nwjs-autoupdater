package updater

import (
	"os"
	"path/filepath"

	"github.com/mholt/archiver"
)

func Update(bundle, instDir, appName string) (error, string) {

  appExec := filepath.Join(instDir, appName)
  appDir := appExec
  appBak := appExec + ".bak"

  err := archiver.Zip.Open(bundle, ".")
	if err != nil {
		return err, appExec
	}

  err = os.Rename(appDir, appBak)
  if err != nil {
    return err, appExec
  }

  updateFiles := filepath.Join(".", appName)

  err = os.Rename(updateFiles, appExec)
  if err != nil {
    os.RemoveAll(appExec)
    os.Rename(appBak, appExec)

    return err, appExec
  }

  err = os.RemoveAll(appBak)
  if err != nil {
    return err, appExec
  }

  err = os.RemoveAll(bundle)
  if err != nil {
    return err, appExec
  }

  return nil, appExec
}
