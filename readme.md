# Boxen

Boxen is a set of easy-to-use configuration scripts for dockctrl.
Build and run any service quickly & repeatably!

Boxen configuration is recursive; each folder is a refinement of its parent.
This allows you to express a complex, structured installation system in a natural way.

Combined with

## Quickstart

Coming soon!

## Using with the Default Docker Daemon

By default, dockctrl will use a `dock` folder it finds in the configuration directories.
If that folder has `docker.pid` and `docker.sock` files, dockctrl connects to an existing daemon, otherwise it will start one for you.
We have added an empty folder, so that running boxen images will be hosted in the same folder.

If you want, you can opt to use your own docker daemon instead.
For example, assuming you have started docker elsewhere, run these commands from the `boxen` folder:

```
rm -rf dock
ln -s /var/run/ && mv run dock
```

All boxen images will now connect to the system-default docker daemon.

## Install

Boxen uses [Docker](http://www.docker.io/), an excellent container helper based on LXC.

On ubuntu 13.04, using the latest packaged installation (0.6.x) works fine. From their [instructions](http://docs.docker.io/en/latest/installation/ubuntulinux/):
```bash
sudo apt-get update
sudo apt-get install linux-image-extra-`uname -r`
sudo sh -c "wget -qO- https://get.docker.io/gpg | apt-key add -"
sudo sh -c "echo deb http://get.docker.io/ubuntu docker main > /etc/apt/sources.list.d/docker.list"
sudo apt-get update
sudo apt-get install lxc-docker
```

You will need Go 1.1 or newer. Boxen assumed it's located in /usr/local/go, but this is easy to [./docker.json](configure).
Following the [golang instructions](http://golang.org/doc/install#bsd_linux) for 64-bit linux:

```bash
curl https://go.googlecode.com/files/go1.1.2.linux-amd64.tar.gz -o golang.tar.gz
sudo tar -C /usr/local -xzf golang.tar.gz
export PATH=$PATH:/usr/local/go/bin # Add this to /etc/profile or similar
```

Grab dockctrl and place it on your path:
```bash
git clone https://github.com/polydawn/dockctrl --recursive
dockctrl/build.sh
sudo cp dockctrl/dockctrl
```

Clone this repo and the libraries the boxen scripts will use:

```bash
git clone https://github.com/polydawn/boxen && cd boxen
git submodule update --init .gopath/src/polydawn.net/pogo
```

Now, you can run `dockctrl launch` or `dockctrl build` in any folder of boxen.
For example, to launch docker's default ubuntu shell:

```bash
cd boxen/ubuntu
dockctrl launch
```

By default, dockctrl will clean up the docker for you ; no hanging leftovers.
Now let's build nginx:

```bash
cd build-essential && dockctrl build && cd nginx && dockctrl build
dockctrl launch
```

You are now running a home-built nginx. Easy!
Along the way, we built the `build-essential` box, which our nginx image is based upon.

