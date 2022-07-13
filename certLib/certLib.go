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

    // root folder
//    folderInfo, err := os.Stat(path)
    _, err = os.Stat(path)
    if !os.IsNotExist(err) {
        return fmt.Errorf("root folder: %s already exists!", path)
    }

    err = os.MkdirAll(path, os.ModePerm)
    if err != nil {
        return fmt.Errorf("could not create root folder %s: %v", path, err)
    }

    // child folders
    //

    return nil
}

func DelCaFolders (path string) (err error) {

    _, err = os.Stat(path)
    if os.IsNotExist(err) {
        return fmt.Errorf("root folder: %s already exists!", path)
    }


    return nil
}
