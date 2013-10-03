#!/bin/bash -e

#Folder layout
mkdir /service && cd /service

# Get the source, checkout tagged release
git clone https://git.kofalt.com/git/public/nginx.git src && cd src
git checkout --quiet release-1.5.4

# Configure
# Nginx's build process creates some default configuration.
# Because we can't avoid this, we point it at /service/default and ignore it.
# When launched, nginx is pointed at /service/conf instead.
./auto/configure \
	--prefix=/service/default \
	--conf-path=/service/default/nginx.conf \
	--pid-path=/service/var/nginx.pid \
	--lock-path=/service/var/nginx.lock \
	--error-log-path=/service/log/error.log \
	--http-log-path=/service/log/access.log \
	--sbin-path=/service/bin/nginx \
	--with-pcre-jit \
	--with-http_ssl_module

# Compile and install
make
make install
