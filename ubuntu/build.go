package main

import (
	. "polydawn.net/gosh/psh"
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

	//Default Ubuntu sources via http://repogen.simplylinux.ch
	var sources =
		"###### Ubuntu Main Repos \n" +
		"deb http://us.archive.ubuntu.com/ubuntu/ quantal main restricted universe \n" +
		"###### Ubuntu Update Repos \n" +
		"deb http://us.archive.ubuntu.com/ubuntu/ quantal-security main restricted universe \n" +
		"deb http://us.archive.ubuntu.com/ubuntu/ quantal-updates main restricted universe \n\n"
		//TODO add universe

	//Update the apt-get sources
	errs[2] = WriteFile("/etc/apt/sources.list", []byte(sources), 0644)

	//Prevent daemons from auto-starting on install or upgrade
	//	See: https://github.com/dotcloud/docker/issues/446#issuecomment-16953173
	errs[3] = WriteFile("/usr/sbin/policy-rc.d", []byte("#!/bin/sh \n" + "exit 101 \n\n"), 0644)

	//Just wrote a bunch of files, check for errors
	for i := range errs {
		if errs[i] != nil {
			Println("Error: Could not write file", i, ":", errs[i])
			os.Exit(1)
		}
	}

	// Pin initscripts, because docker is messing with files that package thinks it owns.
	// Also procps, because it's trying to dial socket /com/ubuntu/upstart and um no.
	// Also udev, because what i don't even
	Sh("apt-mark")("hold", "initscripts", "procps", "udev", "makedev")(DefaultIO)()

	//Update to latest Ubuntu packages!
	apt := Sh("apt-get").
		BakeOpts(DefaultIO).
		BakeEnv(Env{"DEBIAN_FRONTEND": "noninteractive"})
	apt("update",       "-y")()
	apt("dist-upgrade", "-y")()

	// Ask apt to clean up after itself (we don't want to waste disk committing a bunch of caches).
	apt("autoremove",   "-y")()
	apt("clean",        "-y")()
	// Forcibly remove more apt state that is evidentally not actually required.
	Sh("rm")("-rf", "/var/lib/apt/lists")()
}
