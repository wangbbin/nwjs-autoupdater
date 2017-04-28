package updater

import (
  "io/ioutil"
  "os"
  "path/filepath"

  xattr "github.com/ivaxer/go-xattr"
  "github.com/mholt/archiver"

  "github.com/skratchdot/open-golang/open"
)

// Update ...
func Update(bundle, instDir, appName string) (error, string) {
  appExecName := appName + ".app"
  appExec := filepath.Join(instDir, appExecName)
  appDir := appExec
  appBak := appExec + ".bak"

  tempDir, err := ioutil.TempDir("", appName)
  if err != nil {
    return err, appExec
  }
  defer os.RemoveAll(tempDir)

  err = archiver.Zip.Open(bundle, tempDir)
  if err != nil {
    return err, appExec
  }

  err = os.Rename(appDir, appBak)
  if err != nil {
    return err, appExec
  }

  updateFiles := filepath.Join(tempDir, appExecName)

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

  attrData, _ := xattr.Get(appExec, "com.apple.quarantine")

  if len(attrData) == 0 {
    return nil, appExec
  }

  err = xattr.Remove(appExec, "com.apple.quarantine")
  if err != nil {
    return err, appExec
  }

  open.Start(appExec)

  return nil, appExec
}
