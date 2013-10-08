#!/bin/bash -e

# Prepare apt-get for usage
/boxen/scripts/apt-get/update.sh

# Build packages
apt-get install -y build-essential

# Development libraries
apt-get install -y libpcre3-dev zlib1g-dev libssl-dev

# Developer packages
apt-get install -y git curl wget

# Cleanup
/boxen/scripts/apt-get/clean.sh
