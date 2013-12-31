# Boxen

Boxen is a set of easy-to-use configuration scripts for [Hroot](https://github.com/polydawn/hroot).
Build and run any service quickly & repeatably!

The [Hroot tutorial](https://github.com/polydawn/hroot#how-do-i-use-it) will show you how to get up & running.

## Quickstart

Feeling adventurous? With Docker & Hroot on your path, try:

```bash
# Clone down some example config files
git clone https://github.com/polydawn/boxen.git && cd boxen

# Download ubuntu from public index, save into git
cd ubuntu-index && hroot build

# Build your own ubuntu image (updates apt-get)
cd ../ubuntu    && hroot build

# Load repeatable ubuntu from git and start an interactive shell
hroot run bash
```

You're now running a shell in your very own repeatable container.

When you're done with that, let's run nginx:

```bash
# Build the dev-tools container (common build libraries)
cd dev-tools && hroot build

# Build nginx
cd ../nginx  && hroot build

#Launch nginx
hroot run
```

You are now running a home-built nginx with SSL. [Check it out](https://127.0.0.1). *Your nginx is now diamonds*

Along the way, we built the `build-essential` box, which our nginx image is based upon.
