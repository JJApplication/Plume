/*
Name: plume
File: utils_kernel.go
Author: Landers
*/

package main

import (
	"os/exec"
	"strings"
)

func getKernelData() KernelInfo {
	var k KernelInfo

	sh := "uname -r -i"
	cmd := exec.Command("bash", "-c", sh)
	res, e := cmd.Output()
	if e != nil {
		return k
	}

	data := strings.Fields(strings.Trim(string(res), "\n"))
	if len(data) < 2 {
		return k
	}

	sh = "lsb_release -d|sed 's/'$(lsb_release -d|awk '{print $1}')'//g'"
	cmd = exec.Command("bash", "-c", sh)
	res, e = cmd.Output()
	if e != nil {
		k.KernelOs = "linux"
	} else {
		k.KernelOs = checkOS(string(res))
	}

	k.KernelVersion = data[0]
	k.KernelType = data[1]

	return k
}

func checkOS(s string) string {
	con := strings.ToLower(s)
	if strings.Contains(con, "ubuntu") {
		return "ubuntu"
	}

	if strings.Contains(con, "centos") {
		return "centos"
	}

	if strings.Contains(con, "darwin") {
		return "linux"
	}

	if strings.Contains(con, "suse") {
		return "suse"
	}

	if strings.Contains(con, "fedora") {
		return "fedora"
	}

	if strings.Contains(con, "redhat") {
		return "redhat"
	}

	return "linux"
}