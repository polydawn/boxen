#!/bin/bash -e

# Prepare apt-get for usage
. /boxen/scripts/apt-get/update.sh

# Upgrades from ubuntu
apt-get install -y slapd ldap-utils

# Cleanup
. /boxen/scripts/apt-get/clean.sh
