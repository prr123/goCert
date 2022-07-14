// program creCert.go
// credits to:
//  Shane
//
// authors: prr, azul software
// Date: 14 July 2022
// copyright 2022 the authors
//
package main

import (
    "os"
    "fmt"
    "goCert/certLib"
)

func main() {
var path, opt string

    numarg := len(os.Args)

    switch numarg {
        case 1:
            fmt.Println("insufficient cmd line arguments!")
            fmt.Println("usage is:")
            fmt.Println("   creCa folder")
            os.Exit(-1)
        case 2:
            path = os.Args[1]

        case 3:
            path = os.Args[1]
            opt  = os.Args[2]

        default:
            fmt.Println("too many cmd line arguments!")
            fmt.Println("usage is:")
            fmt.Println("   creCa folder")
            os.Exit(-1)
    }

    if len(opt) > 0 {
        fmt.Printf("cmd line: folder path: %s option: %s\n", path, opt)
    } else {
        fmt.Printf("cmd line: folder path: %s no option\n", path)
    }

    switch opt {
    case "del":
        fmt.Println("*** deleting Ca folders ***")
        err := certLib.DelCaFolders(path)
        if err != nil {
            fmt.Printf("error DelCaFolders: %v\n", err)
            os.Exit(-1)
        }

    default:
        fmt.Println("*** creating Ca folders ***")
        err := certLib.CreCaFolders(path)
        if err != nil {
            fmt.Printf("error CreCaFolders: %v\n", err)
            os.Exit(-1)
        }
    }

    fmt.Println("***** success *****")
}

