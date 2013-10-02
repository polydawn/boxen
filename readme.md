# Boxen

Boxen is a set of easy-to-use configuration scripts for dockctrl.
Build and run any service quickly & repeatably!

Boxen configuration is recursive; each folder is a refinement of its parent.
This allows you to express a complex, structured installation system in a natural way.

Combined with

## Quickstart

Boxen uses [Docker](http://www.docker.io/), an excellent container helper based on LXC.

On ubuntu, using the latest packaged installation (0.6.x) works fine:
```bash
sudo apt-get install docker
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

