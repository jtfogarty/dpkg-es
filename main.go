package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/jtfogarty/usb_dpkg/DpkgList"
)

func main() {
	var (
		cmdOut []byte
		err    error
	)
	cmdName := "dpkg"
	cmdArgs := []string{"--list"}
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git rev-parse command: ", err)
		os.Exit(1)
	}
	sha := string(cmdOut)
	firstSix := sha[:100]
	fmt.Println("The first six chars of the list of packages", firstSix)
	s := strings.Split(sha, "\n")
	fmt.Println("The number of lines is", len(s))
	for _, l := range s {
		if len(l) != 0 {
			k := string(l)
			if k[:2] != "De" && k[:1] != "||" && k[:3] != "+++" {
				DpkgList.ListPkgs(k)
			}
		}
	}
}
