#!/bin/bash -e

# Folder layout
mkdir /service && cd /service

# Prepare apt-get for usage
. /boxen/scripts/apt-get/update.sh

# Developer tools
apt-get install -y git

# Node dependancies
apt-get install -y g++ python2.7 make

# Closure compiler dependancies
apt-get install -y openjdk-7-jre-headless

# Sass dependancies
apt-get install -y ruby1.9.1

# Clone node
git clone https://github.com/joyent/node.git && cd node
git checkout -q v0.10.28

# Build node
$PYTHON ./configure
make
make install

# Install Sass
gem install sass --no-ri --no-rdoc

# Install less
npm install -g less

# Common node tools
npm install -g grunt-cli
npm install -g bower

# Cleanup
. /boxen/scripts/apt-get/clean.sh
