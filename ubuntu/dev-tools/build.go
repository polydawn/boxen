package main

import (
	. "polydawn.net/pogo/gosh"
)

func main() {
	// Sh("echo").BakeOpts(DefaultIO)("hello")()

	//Prepare apt-get command template for non-interactive installs
	apt := Sh("apt-get").
		BakeOpts(DefaultIO).
		BakeEnv(Env{"DEBIAN_FRONTEND": "noninteractive"})

	//Package check
	apt("update", "-y")()

	//Build packages
	apt("install", "-y" , "build-essential", "libpcre3-dev", "zlib1g-dev", "libssl-dev")()

	//Developer packages
	apt("install", "-y" , "git")()

	//Ask apt to remove cache state that shouldn't be preserved in an image
	apt("autoremove", "-y")()
	apt("clean",      "-y")()

	//Forcibly remove more apt state
	Sh("rm")("-rf", "/var/lib/apt/lists")()

}
