#!/bin/bash

# Check we have permissions
# if [[ $EUID -ne 0 ]]; then
#    echo "This script must be run as root" 1>&2
#    exit 1
# fi

if [ "$1" == "-u" ]||[ "$1" == "--uninstall" ]; then
    echo "Uninstalling application"
    rm -f "/usr/local/bin/git-user"
    rm -f "/usr/local/share/man/man1/git-user*"
    exit 0;
fi

# Remove the app if it already exists
APP="/usr/local/bin/git-user"
if [ -d "$APP" ]; then
    echo "Removing existing application"
    rm -rf "${APP}"
fi

# Copy our executable to /usr/local/bin
echo "Installing git-user"
cp -f git-user /usr/local/bin

# Copy our man files
mkdir -p /usr/local/share/man/man1
cp man/*.1 /usr/local/share/man/man1
