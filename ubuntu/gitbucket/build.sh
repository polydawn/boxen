#!/bin/bash -e

# Download tagged release
mkdir -p /service
wget https://github.com/takezoe/gitbucket/releases/download/1.9/gitbucket.war -O /service/gitbucket.war

# Install java
. /boxen/scripts/apt-get/update.sh
apt-get install -y openjdk-7-jre-headless
. /boxen/scripts/apt-get/clean.sh
