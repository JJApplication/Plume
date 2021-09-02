/*
Name: plume
File: cmds.go
Author: Landers
*/

package main

import (
	"fmt"
	"os/exec"
)

// 封装的cmd
func cmdRun(sh string, debug bool) ([]byte, error) {
	cmd := exec.Command("bash", "-c", sh)
	res, e := cmd.Output()

	if debug {
		fmt.Printf("[cmd debug] %s\n", cmd.String())
	}
	if e != nil {
		defer cmd.Process.Kill()
		return nil, e
	}

	defer cmd.Process.Kill()
	return res, e
}
