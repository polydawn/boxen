#!/bin/bash -e

#
#Prepare apt-get for installs/updates
#

# Set to non-interactive installs
export DEBIAN_FRONTEND=noninteractive

# You'll need to run this redundantly sometimes because cleanup.sh clears state
apt-get update -y
