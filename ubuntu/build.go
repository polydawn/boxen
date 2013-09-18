package main

import (
	. "polydawn.net/gosh/psh"
	. "io/ioutil"
	. "fmt"
	"os"
)

func main() {
	//Default Ubuntu sources via http://repogen.simplylinux.ch
	var sources =
		"###### Ubuntu Main Repos \n" +
		"deb http://us.archive.ubuntu.com/ubuntu/ quantal main restricted universe \n" +
		"###### Ubuntu Update Repos \n" +
		"deb http://us.archive.ubuntu.com/ubuntu/ quantal-security main restricted universe \n" +
		"deb http://us.archive.ubuntu.com/ubuntu/ quantal-updates main restricted universe \n\n"
		//TODO add universe

	//Update the apt-get sources
	err := WriteFile("/etc/apt/sources.list", []byte(sources), 0644)
	if err != nil {
		Println("Could not write sources file:", err)
		os.Exit(1)
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
