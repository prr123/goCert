// certLib
// credits to:
//  Shane
//
// authors: prr, azul software
// Date: 14 July 2022
// copyright 2022 the authors
//
package certLib

import (
    "os"
    "fmt"
)

func CreCaFolders (path string) (err error) {

//    folderInfo, err := os.Stat(path)
    _, err = os.Stat(path)
    if !os.IsNotExist(err) {
        return fmt.Errorf("folder: %s already exists!", path)
    }

    err = os.MkdirAll(path, os.ModePerm)
    if err != nil {
        return fmt.Errorf("could not create path %s: %v", path, err)
    }

    return nil
}
