#!/bin/bash -e

#
# Clean up apt-get state that should not go in an image
#

apt-get autoremove -y
apt-get clean -y
rm -rf /var/lib/apt/lists
