#!/bin/bash

# This script installs Vale on Linux

# Exit script immediately if any command fails
set -e

# Set the desired Vale version
VALE_VERSION="3.0.5"

# Define the URL for the Vale binary
VALE_URL="https://github.com/errata-ai/vale/releases/download/v$VALE_VERSION/vale_${VALE_VERSION}_Linux_64-bit.tar.gz"

# Directory to store temporary files during installation
TEMP_DIR=$(mktemp -d)

# Download and extract the Vale binary
wget -O "$TEMP_DIR/vale.tar.gz" "$VALE_URL"
tar -xzvf "$TEMP_DIR/vale.tar.gz" -C "$TEMP_DIR"

# Move the Vale binary to /usr/local/bin/
# The /usr/local/bin directory is a standard location for installing
# user-specific executables. Placing binaries here ensures they are accessible
# system-wide without conflicting with packages managed by the system's package
# manager. Using this so users also (usually) won't need to update their PATH
# variable which also depends on which shell they are using.
sudo mv "$TEMP_DIR/vale" /usr/local/bin/

# Clean up temporary files
rm -rf "$TEMP_DIR"

# Print a message indicating successful installation
echo "Vale $VALE_VERSION has been successfully installed to /usr/local/bin/"