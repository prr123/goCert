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

    folderInfo, err := os.Stat(path)
    if !os.IsNotExist(err) {
        return fmt.Errorf("folder: %s already exists!", path)
    }

    return nil
}
