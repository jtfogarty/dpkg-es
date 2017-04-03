package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/jtfogarty/dpkg-es/config"
	"github.com/jtfogarty/dpkg-es/dpkgList"
)

var configObj config.Object

func init() {
	configJSON := config.GetConfig()
	//fmt.Println(configJSON)
	configObj = *configJSON //i really do not understand why I have to do this
	//i was thinking that I could set configObj := config.GetConfig() something to do with * / &
}

func main() {
	var (
		cmdOut []byte
		err    error
	)
	cmdName := "dpkg"
	cmdArgs := []string{"--list"}
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running 'dpkg --list' command: ", err)
		os.Exit(1)
	}
	sha := string(cmdOut)
	s := strings.Split(sha, "\n")
	for _, l := range s {
		if len(l) != 0 {
			k := string(l)
			if k[:2] != "De" && k[:1] != "|" && k[:3] != "+++" {
				DpkgList.ListPkgs(k, configObj)
			}
		}
	}
}
