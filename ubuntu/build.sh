#!/bin/bash -e

# Prepare apt-get for usage
/boxen/scripts/apt-get/bootstrap.sh
/boxen/scripts/apt-get/update.sh

# Upgrades from ubuntu
apt-get dist-upgrade -y

# Cleanup
/boxen/scripts/apt-get/clean.sh
