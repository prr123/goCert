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
    // certs
    certPath := path + "/certs"
    err = os.Mkdir(certPath, os.ModePerm)
    if err != nil {
        return fmt.Errorf("could not create cert folder %s: %v", certPath, err)
    }

    // crl
    crlPath := path + "/crl"
    err = os.Mkdir(crlPath, os.ModePerm)
    if err != nil {
        return fmt.Errorf("could not create crl folder %s: %v", crlPath, err)
    }

    // private
    privPath := path + "/private"
    err = os.Mkdir(privPath, os.ModePerm)
    if err != nil {
        return fmt.Errorf("could not create private folder %s: %v", privPath, err)
    }

    // files
    indexFilnam := path + "/index.txt"
    _, err = os.Create(indexFilnam)
    if err != nil {
        return fmt.Errorf("could not create index file %s: %v", indexFilnam, err)
    }

    serialFilnam := path + "/serial"
    _, err = os.Create(serialFilnam)
    if err != nil {
        return fmt.Errorf("could not create serial file %s: %v", serialFilnam, err)
    }

    return nil
}

func DelCaFolders (path string) (err error) {

    _, err = os.Stat(path)
    if os.IsNotExist(err) {
        return fmt.Errorf("root folder: %s does not exist!", path)
    }

    err = os.RemoveAll(path)
    if err != nil {
        return fmt.Errorf("error deleting rootfolder %s: %v", path, err)
    }
    return nil
}
