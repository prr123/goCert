// creCa.go
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
    "goCert/certLib"
)

func main() {

    numarg := len(os.Args)

    switch numarg {
        case 1:
            fmt.Println("too many cmd line arguments!")
            fmt.Println("usage is:")
            fmt.Println("   creCa folder")
            os.Exit(-1)
        case 2:
            path := os.Args[1]

        case 3:
            path := os.Args[1]
            opt  := os.Args[2]

        default:
            fmt.Println("too many cmd line arguments!")
            fmt.Println("usage is:")
            fmt.Println("   creCa folder")
            os.Exit(-1)
    }

    fmt.Printf("folder path: %s option: %s\n", path, opt)

    err := certLib.CreCaFolders(path)

    fmt.Println("***** success *****")
}

