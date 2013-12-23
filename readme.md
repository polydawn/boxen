# Boxen

Boxen is a set of easy-to-use configuration scripts for [Docket](https://github.com/polydawn/docket).
Build and run any service quickly & repeatably!

The [Docket tutorial](https://github.com/polydawn/docket#how-do-i-use-it) will show you how to get up & running.

## Quickstart

Feeling adventurous? With Docker & Docket on your path, try:


```bash
# Clone down some example config files
git clone https://github.com/polydawn/boxen.git && cd boxen/ubuntu

# Download ubuntu from public index, save into git
docket build -s index -d graph --noop

# Upgrade apt-get packages
docket build

# Load repeatable ubuntu from git and start an interactive shell
docket run bash
```

You're now running a shell in your very own repeatable container.

When you're done with that, let's run nginx:

```bash
# Build the dev-tools container (common build libraries)
cd dev-tools && docket build

# Build nginx
cd ../nginx && docket build

#Launch nginx
docket run
```

You are now running a home-built nginx with SSL. [Check it out](https://127.0.0.1). *Your nginx is now diamonds*

Along the way, we built the `build-essential` box, which our nginx image is based upon.
