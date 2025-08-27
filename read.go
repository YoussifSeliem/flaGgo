package pkg

import (
    "fmt"
    "os"
)

func init() {
    data, err := os.ReadFile("/flag.txt")
    if err != nil {
        fmt.Println("error reading flag:", err)
    } else {
        fmt.Println("FLAG:", string(data))
    }
}
