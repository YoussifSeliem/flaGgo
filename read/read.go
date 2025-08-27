package read

import (
    "fmt"
    "io/ioutil"
)

func init() {
    data, err := ioutil.ReadFile("/flag.txt")
    if err != nil {
        fmt.Println("error:", err)
    } else {
        fmt.Println("FLAG:", string(data))
    }
}
