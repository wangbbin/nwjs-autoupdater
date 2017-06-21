package updater

import (
  "os"
  "path/filepath"
  "github.com/mholt/archiver"
  "github.com/skratchdot/open-golang/open"
)

func Update(bundle, instDir, appName string) (error, string) {

  extractDir := "./files";
  if _, err := os.Stat(extractDir); err == nil {
    os.RemoveAll(extractDir)
  }

  appExecName := appName + ".app"
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

      // If the extracted file exist in the installation, rename it to end with .bak
      oldFileBackup := ""
      if _, err = os.Stat(newFileInstPath); err == nil {
        oldFileBackup = newFileInstPath + ".bak"
        err = os.Rename(newFileInstPath, oldFileBackup)
        if err != nil {
          return err
        }
      }

      // Move the extracted file to instdir
      err = os.Rename(extractedFile, newFileInstPath)
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

  err = os.RemoveAll(extractDir)
  if err != nil {
    return err, appExec
  }
  err = os.RemoveAll(bundle)
  if err != nil {
    return err, appExec
  }

  /*cmd := exec.Command(appExec)
  err = cmd.Start()
  if err != nil {
    return err, appExec
  }*/

  open.Start(appExec)
  
  return nil, appExec
}
