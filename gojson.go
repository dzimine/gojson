// Pretty-prints JSON from stdin.
// Usage cat myfile.json | gojson

package main

import (
    "encoding/json"
    "os"
    "fmt"
    "io/ioutil"
)

func main() {
    // Read from stdin
    bytes, err := ioutil.ReadAll(os.Stdin)
    if err != nil {
        panic(err)
    }
    // fmt.Println(bytes)

    // Turn string to JSON
    var dat map[string]interface{}
    if err := json.Unmarshal(bytes, &dat); err != nil {
        panic(err)
    }

    // Pretty-print
    res, err := json.MarshalIndent(dat, "", "  ")
    if err != nil {
        panic(err)
    }
    os.Stdout.Write(res)
    fmt.Println()

}