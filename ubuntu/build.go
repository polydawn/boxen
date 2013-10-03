package main

import (
	. "polydawn.net/pogo/gosh"
	. "io/ioutil"
	. "fmt"
	"os"
)

func main() {
	errs := make([]error, 4) // Handle write errors once

	//Speed up installs and don't create cache files
	//	See: https://github.com/dotcloud/docker/pull/1883#issuecomment-24434115
	errs[0] = WriteFile("/etc/dpkg/dpkg.cfg.d/02apt-speedup", []byte("force-unsafe-io"), 0644)
	errs[1] = WriteFile("/etc/apt/apt.conf.d/no-cache",       []byte("Acquire::http {No-Cache=True;};"), 0644)

	//Prevent daemons from auto-starting on install or upgrade
	//	See: https://github.com/dotcloud/docker/issues/446#issuecomment-16953173
	errs[2] = WriteFile("/usr/sbin/policy-rc.d", []byte("#!/bin/sh\n" + "exit 101\n\n"), 0755)

	//Tell apt to use a nearby mirror. Can *dramatically* increase update speed.
	var mirrors =
		"###### Mirrors for apt-get" +
		"deb mirror://mirrors.ubuntu.com/mirrors.txt quantal main restricted universe multiverse \n" +
		"deb mirror://mirrors.ubuntu.com/mirrors.txt quantal-updates main restricted universe multiverse \n" +
		"deb mirror://mirrors.ubuntu.com/mirrors.txt quantal-security main restricted universe multiverse \n\n"

	//Default Ubuntu sources generated via http://repogen.simplylinux.ch
	var sources =
		"###### Ubuntu Main Repos \n" +
		"deb http://us.archive.ubuntu.com/ubuntu/ quantal main restricted universe \n" +
		"###### Ubuntu Update Repos \n" +
		"deb http://us.archive.ubuntu.com/ubuntu/ quantal-security main restricted universe \n" +
		"deb http://us.archive.ubuntu.com/ubuntu/ quantal-updates main restricted universe \n\n"
		//TODO add universe

	//Update the apt-get sources
	errs[3] = WriteFile("/etc/apt/sources.list", []byte(mirrors + sources), 0644)

	//Just wrote a bunch of files, check for errors
	for i := range errs {
		if errs[i] != nil {
			Println("Error: Could not write file", i, ":", errs[i])
			os.Exit(1)
		}
	}

	// Pin initscripts, because docker conflicts with files that package thinks it owns.
	Sh("apt-mark")("hold", "initscripts")(DefaultIO)()

	//Prepare apt-get command template for non-interactive installs
	apt := Sh("apt-get").
		BakeOpts(DefaultIO).
		BakeEnv(Env{"DEBIAN_FRONTEND": "noninteractive"})

	//Update to latest Ubuntu packages
	apt("update",       "-y")()
	apt("dist-upgrade", "-y")()

	//Ask apt to remove cache state that shouldn't be preserved in an image
	apt("autoremove",   "-y")()
	apt("clean",        "-y")()

	//Forcibly remove more apt state
	Sh("rm")("-rf", "/var/lib/apt/lists")()
}
