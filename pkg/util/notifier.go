package util

import (
	"fmt"
	"os/exec"
)


func NotifyMacOS(title, message string) error{
	cmd := exec.Command("osascript", "-e", 
	fmt.Sprintf("display notification \"%s\" with title \"%s\"", message, title))

	return cmd.Run()
}