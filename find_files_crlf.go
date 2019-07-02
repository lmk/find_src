package main

import (
        "fmt"
        "os"
        "path/filepath"
        "io/ioutil"
        "strings"
)

func visit(path string, f os.FileInfo, err error) error {
        if f.IsDir() == true {
                return nil
        }

        if filepath.Ext(path) == ".cpp" || filepath.Ext(path) == ".c" || filepath.Ext(path) == ".h" {
                //fmt.Printf("Visited: %s\n", path)
                file, _ := os.Open(path)
                defer file.Close()

                b, _ := ioutil.ReadAll(file)
                if strings.Contains(string(b), "\r\n") {
                        fmt.Printf("%s\n", path)
                }
        }

        return nil
}

func main() {
        root := os.Args[1]
        fmt.Printf("root: %s\n", root)
        err := filepath.Walk(root, visit)
        fmt.Printf("filepath.Walk() returned %v\n", err)
}
