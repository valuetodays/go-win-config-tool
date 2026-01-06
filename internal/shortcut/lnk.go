package shortcut

import (
	"os/exec"
	"strings"
)

func ReadTarget(lnk string) (string, error) {
	cmd := exec.Command(
		"powershell",
		"-NoProfile",
		"-Command",
		`(New-Object -COM WScript.Shell).CreateShortcut("`+lnk+`").TargetPath`,
	)

	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}
