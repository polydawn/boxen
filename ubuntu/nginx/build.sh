#!/bin/bash -e

#Folder layout
mkdir -p /service/nginx
mkdir -p /jank
cd /service/nginx

# Possibly roll some packages into upstream
apt-get update
apt-get install -y git

# Get the source
git clone https://git.kofalt.com/git/public/nginx.git src
cd src

# Build deps
apt-get install -y g++ make libpcre3-dev zlib1g-dev libssl-dev

./auto/configure \
	--prefix=/jank \
	--conf-path=/jank/nginx.conf \
	--pid-path=/service/var/nginx.pid \
	--lock-path=/service/var/nginx.lock \
	--error-log-path=/service/log/error.log \
	--http-log-path=/service/log/access.log \
	--sbin-path=/service/bin/nginx \
	--with-pcre-jit \
	--with-http_ssl_module


make
make install
