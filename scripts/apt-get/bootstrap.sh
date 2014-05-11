#!/bin/bash -e

#
# Configure apt-get for living in a docker
#

# Speed up installs and don't create cache files
#	See: https://github.com/dotcloud/docker/pull/1883#issuecomment-24434115
echo "force-unsafe-io"                 > /etc/dpkg/dpkg.cfg.d/02apt-speedup
echo "Acquire::http {No-Cache=True;};" > /etc/apt/apt.conf.d/no-cache
chmod 0644 /etc/dpkg/dpkg.cfg.d/02apt-speedup
chmod 0644 /etc/apt/apt.conf.d/no-cache

# Prevent daemons from auto-starting on install or upgrade
#	See: https://github.com/dotcloud/docker/issues/446#issuecomment-16953173
cat > /usr/sbin/policy-rc.d <<EOF
#!/bin/sh
exit 101
EOF
chmod 0755 /usr/sbin/policy-rc.d

# Get release codename ("saucy", "trusty", etc)
codename=`/usr/bin/lsb_release -cs`

# Apt-get sources:
#	Tell apt to use a nearby mirror. Can *dramatically* increase update speed.
#	Default Ubuntu sources generated via http://repogen.simplylinux.ch
cat > /etc/apt/sources.list <<EOF
###### Ubuntu Main Repos
deb http://mirrors.linode.com/ubuntu/ $codename main restricted universe

###### Ubuntu Update Repos
deb http://mirrors.linode.com/ubuntu/ $codename-security main restricted universe
deb http://mirrors.linode.com/ubuntu/ $codename-updates  main restricted universe
EOF
chmod 0644 /etc/apt/sources.list

# Pin initscripts, because docker conflicts with files that package thinks it owns.
apt-mark hold initscripts makedev
