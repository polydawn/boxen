# Boxen

Boxen is a set of easy-to-use configuration scripts for [Docket](https://github.com/polydawn/docket).
Build and run any service quickly & repeatably!

Docket configuration is recursive; each folder is a refinement of its parent.
This allows you to express a complex, structured ecosystem in a natural way.
Boxen shows this off a bit by arranging example Docket files for various popular services.

## Quickstart

With [Docket](https://github.com/polydawn/docket) on your path, try:

```bash
git clone https://github.com/polydawn/boxen
cd boxen/ubuntu
docket run
```

That was all it took to launch an ubuntu docker with some basic configuration.

To launch a bash shell, add the target:

```
docket run bash
```

By default, boxen will clean up the images for you; no garbage containers need apply.
Now let's build nginx:

```bash
cd dev-tools && docket export && cd nginx && docket export
docket run
```

You are now running a home-built nginx with SSL. [Check it out](https://127.0.0.1). *Your nginx is now diamonds*

Along the way, we built the `build-essential` box, which our nginx image is based upon.

## Using the default docker daemon

By default, docket will use a `dock` folder it finds in the configuration directories.
If that folder has `docker.pid` and `docker.sock` files, docket connects to an existing daemon, otherwise it will start one for you.
We have added an empty folder, so all running boxen images will be hosted in the same folder.

If you want, you can opt to use your own docker daemon instead.
To use the default daemon, run this from the `boxen` folder:

```
(cd dock && ln -s /var/run/docker.pid && ln -s /var/run/docker.sock)
```

All boxen images will now connect to the system-default docker daemon.
