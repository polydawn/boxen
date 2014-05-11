#!/bin/bash -e

# Folder layout
mkdir /service && cd /service

# Prepare apt-get for usage
. /boxen/scripts/apt-get/update.sh

# Developer tools
apt-get install -y git

# Node dependancies
apt-get install -y g++ python2.7 make

# Seriously? Your build process, over point releases, has requested:
#  * python in PATH
#  * python2.7 in PATH
#  * a $PYTHON environment variable
# ... all pointing to python 2.7.
#
# Mix in one part Canonical endlessly playing Where's Waldo, Python Edition and blend on high for a delicious smoothie!
#
# I have 8 bosses Bob. Pardon me? 8 Bosses. Eight? Eight Bob! That means that when I screw something up, I have to hear 8 different people tell me about it.
(cd /usr/bin; ln -s python2.7 python)

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
